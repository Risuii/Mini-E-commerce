package models

type Barang struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	IdToko      int    `json:"id_toko"`
	Nama        string `gorm:"varchar(300)" json:"name"`
	Description string `gorm:"varchar(300)" json:"Description"`
	Stock       int
	Harga       int
}
