package models

import (
	"database/sql"
	"linksuara/app/config"
	"linksuara/app/utility"
	"time"

	"gorm.io/gorm"
)

type Dapil struct {
	ID   sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Nama sql.NullString `json:"nama" gorm:"column:nama"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (d *Dapil) FindAll() ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var daftarDapil []map[string]any // Menampung Data
	rows, err := db.Raw("SELECT * FROM dapil").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&d.ID,
			&d.Nama,
			&d.CreatedAt,
			&d.UpdatedAt,
			&d.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		// idString := string(rune(d.ID.Int64))
		// hashedId, err := utility.GenerateHash(idString)
		// if err != nil {
		// 	return nil, err
		// }
		dapil := map[string]any{
			"id":          d.ID.Int64,
			"nama":        utility.Capitalize(d.Nama.String),
			"nomor_dapil": d.ID.Int64,
		}

		daftarDapil = append(daftarDapil, dapil)
	}

	return daftarDapil, nil
}

func (dp *Dapil) TotalDapil() (int, error) {
	db := config.ConnectGormDB()
	var query string = "SELECT COUNT(*) AS total FROM dapil WHERE deleted_at IS NULL;"

	var total int // Variabel menampung jumlah dapil
	results := db.Raw(query).Scan(&total)
	if results.Error != nil {
		return -1, results.Error
	}

	return total, nil
}
