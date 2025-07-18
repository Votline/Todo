package repo

import (
	"github.com/jmoiron/sqlx"
	sq "github.com/Masterminds/squirrel"

	"todo-service/internal/models"
)

type TodoRepoSql struct {
	db *sqlx.DB
	bd sq.StatementBuilderType
}
func NewTRS(sourceDB *sqlx.DB) *TodoRepoSql {
	return &TodoRepoSql{
		db: sourceDB,
		bd: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (trs *TodoRepoSql) AddUser(user models.User) error {
	query, args, err := trs.bd.
		Insert("users").
		Columns(
			"id", "first_name", "last_name", "password_hash").
		Values(
			user.Id, user.FName, user.LName, user.PdHash).
		ToSql()
	if err != nil {return err}

	_, err = trs.db.Exec(query, args...)
	return err
}

func (trs *TodoRepoSql) AddOrUpdTask(task *models.Task, userID string) error {
	var taskIDPart interface{}
	if task.ID != nil && *task.ID != 0 {
		taskIDPart = task.ID
	} else {
		taskIDPart = sq.Expr(`COALESCE((SELECT MAX(task_id) +1 
			FROM tasks WHERE user_id = ?), 1)`, userID)
	}
	
	query, args, err := trs.bd.
		Insert("tasks").
		Columns(
			"user_id", "task_id",
			"title", "content", "category_id", "done").
		Values(
			userID, taskIDPart,
			task.Title, task.Content, task.Category, task.Done).
		Suffix(`
			ON CONFLICT (user_id, task_id) DO UPDATE SET
				title = EXCLUDED.title,
				content = EXCLUDED.content,
				category_id = EXCLUDED.category_id,
				done = EXCLUDED.done
			RETURNING task_id

		`).
		ToSql()
	if err != nil {return err}

	return trs.db.QueryRow(query, args...).Scan(&task.ID)
}

func (trs *TodoRepoSql) GetTask(userID string, task models.Task, conditions map[string]interface{}) ([]models.Task, error) {
	qb := trs.bd.Select("*").From("tasks").Where(sq.Eq{"user_id": userID})

	for field, value := range conditions {
		qb = qb.Where(sq.Eq{field: value})
	}
	query, args, err := qb.ToSql()
	if err != nil {return nil, err}
	
	var rows *sqlx.Rows
	rows, err = trs.db.Queryx(query, args...)
	if err != nil {return nil, err}

	var tasks []models.Task
	var newTask models.Task
	for rows.Next() {
		if err := rows.StructScan(&newTask); err != nil {
			return nil, err
		}
		tasks = append(tasks, newTask)
	}
	rows.Close()
	return tasks, nil
}
