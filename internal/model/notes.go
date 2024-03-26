package internal

type Notes struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
