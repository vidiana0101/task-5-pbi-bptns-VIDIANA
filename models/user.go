package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(100)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Photo    []byte `gorm:"type:binary(255)" json:"photo"`
}
