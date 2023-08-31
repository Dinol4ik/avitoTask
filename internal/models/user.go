package models

type User struct {
	Name string `json:"name"`
}
type UserId struct {
	Id int `db:"id"`
}
