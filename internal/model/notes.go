package internal

type Notes struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
