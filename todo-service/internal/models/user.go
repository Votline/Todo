package models

type User struct {
	Id     string `db:"id"            json:"id"`
	FName  string `db:"first_name"    json:"first_name"`
	LName  string `db:"last_name"     json:"last_name"`
	PdHash string `db:"password_hash" json:"password_hash"`
}
