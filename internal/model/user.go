package model

type User struct {
	ID     int
	Name   string
	Age    int
	Verify bool   `gorm:"column:trusted"`
	Cards  []Card `gorm:"foreignKey:UserID"`
}
