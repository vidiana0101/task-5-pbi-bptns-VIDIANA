package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	username string `gorm:"type:varchar(100)" json:"username"`
	password string `gorm:"type:varchar(100)" json:"password"`
	email    string `gorm:"type:varchar(100)" json:"email"`
	photo    []byte `gorm:"type:binary(255)" json:"photo"`
}
