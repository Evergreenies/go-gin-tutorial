package internal

type User struct {
	ID       int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
