package repo

import (
	"fmt"

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
	if task.TaskID != 0 {
		taskIDPart = task.TaskID
	} else {
		taskIDPart = sq.Expr(`COALESCE((SELECT MAX(task_id) +1 
			FROM tasks WHERE user_id = ?), 1)`, userID)
	}
	fmt.Printf("\n\n\n%v\n%v\n\n\n", taskIDPart, task.TaskID)
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

	return trs.db.QueryRow(query, args...).Scan(&task.TaskID)
}
