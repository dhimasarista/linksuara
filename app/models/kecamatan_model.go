package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Kecamatan struct {
	ID   sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Nama sql.NullString `json:"nama" gorm:"column:nama"`

	// Foreign Key
	Dapil   Dapil         `gorm:"foreignKey:DapilID" json:"dapil"`
	DapilID sql.NullInt64 `gorm:"column:dapil_id" json:"dapil_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
