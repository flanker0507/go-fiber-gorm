package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Deskripsi string         `json:"deskripsi"`
	Harga     string         `json:"harga"`
	Stok      string         `json:"stok"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
