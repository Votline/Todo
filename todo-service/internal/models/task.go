package models

type Task struct {
	Id       string `db:"id"       json:"id"`
	UserId   string `db:"user_id"  json:"user_id"`
	Title    string `db:"title"    json:"title"`
	Content  string `db:"content"  json:"content"`
	Category string `db:"category" json:"category"`
	Done     bool   `db:"done"     json:"done"`
}
