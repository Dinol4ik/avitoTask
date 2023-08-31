package models

type Csv struct {
	UserName    string `db:"user_name" csv:"user"`
	SegmentName string `db:"segment_name" csv:"segment"`
	CreatedAt   string `db:"created_at" csv:"createdAt"`
	UpdatedAt   string `db:"updated_at" csv:"updatedAt"`
	IsRemoved   bool   `db:"is_removed" csv:"isRemoved"`
}

type DateFilter struct {
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	UserId    int    `json:"userId"`
}
