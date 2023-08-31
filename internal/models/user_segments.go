package models

type UserSegments struct {
	UserId   int `db:"user_id"`
	Segments []Segment
}
