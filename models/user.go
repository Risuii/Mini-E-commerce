package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Nama     string `gorm:"varchar(300)" json:"name"`
	Username string `gorm:"varchar(300)" json:"user_name"`
	Password string `gorm:"varchar(300)" json:"password"`
	Saldo    int    `json:"saldo"`
	Toko     Toko   `gorm:"foreignKey:UserId"`
}
