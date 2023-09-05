package model

type User struct {
	ID       int
	Name     string
	FullName string
	Email    string
	Age      int
	Verify   bool   `gorm:"column:trusted"`
	Cards    []Card `gorm:"foreignKey:UserID"`
	IsAdmin  bool
}
