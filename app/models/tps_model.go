package models

import (
	"database/sql"
	"linksuara/app/config"
	"time"

	"gorm.io/gorm"
)

type TPS struct {
	ID   sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Nama sql.NullString `json:"nama" gorm:"column:nama"`

	// Foreign Key
	Kelurahan   Kelurahan     `gorm:"foreignKey:KelurahanID" json:"kelurahan"`
	KelurahanID sql.NullInt64 `gorm:"column:kelurahan_id" json:"kelurahan_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (tps *TPS) TotalTPS() (int, error) {
	db := config.ConnectGormDB()
	var query string = "SELECT COUNT(*) AS total FROM tps WHERE deleted_at IS NULL;"

	var total int // Variabel menampung jumlah dapil
	results := db.Raw(query).Scan(&total)
	if results.Error != nil {
		return -1, results.Error
	}

	return total, nil
}
