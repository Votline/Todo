package models

type Task struct {
	ID       *int    `db:"task_id"     json:"task_id,omitempty"`
	UserID   *string `db:"user_id"     json:"-"`
	Title    *string `db:"title"       json:"title",omitempty`
	Content  *string `db:"content"     json:"content",omitempty`
	Category *string `db:"category_id" json:"category_id",omitempty`
	Done     *bool   `db:"done"        json:"done",omitempty`
}
