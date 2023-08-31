package models

type Segment struct {
	Name         string  `json:"Name" db:"name"`
	PercentUsers float64 `json:"PercentUsers"`
}

type SegmentsId struct {
	Id []int `db:"id"`
}
