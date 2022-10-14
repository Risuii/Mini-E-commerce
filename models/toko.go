package models

type Toko struct {
	Id          int64    `gorm:"primaryKey" json:"id"`
	UserId      int64    `json:"UserId"`
	Nama        string   `gorm:"varchar(300)" json:"name"`
	Description string   `gorm:"varchar(300)" json:"description"`
	Barang      []Barang `gorm:"foreignKey:IdToko"`
}
