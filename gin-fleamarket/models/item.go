package models

import "gorm.io/gorm"

type Item struct {
	// 物理メモリをDBにする場合のコード
	//ID          uint

	// RBD + GORMの場合のコード
	gorm.Model

	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"`
}
