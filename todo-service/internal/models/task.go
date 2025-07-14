package models

type Task struct {
	TaskID   int    `db:"task_id"`
	UserID   string `db:"user_id"  json:"-"`
	Title    string `db:"title"    json:"title"`
	Content  string `db:"content"  json:"content"`
	Category string `db:"category" json:"category"`
	Done     bool   `db:"done"     json:"done"`
}
