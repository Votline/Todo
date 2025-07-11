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
