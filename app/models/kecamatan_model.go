package models

import (
	"database/sql"
	"linksuara/app/config"
	"linksuara/app/utility"
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

func (kc *Kecamatan) GetByID(id int) error {
	db := config.ConnectGormDB()

	query := "SELECT * FROM kecamatan WHERE id = ?;"
	results := db.Raw(query, id).Scan(&kc)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (kc *Kecamatan) FindAll() ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var query string = `
		SELECT
			kc.id,
			kc.nama AS nama_kecamatan,
			kc.dapil_id,
			dp.nama AS nama_dapil
		FROM
			kecamatan kc
		JOIN
			dapil dp ON kc.dapil_id = dp.id
		WHERE 
			kc.id > 1 AND kc.deleted_at IS NULL;
	`
	var daftarKecamatan []map[string]any

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&kc.ID,
			&kc.Nama,
			&kc.DapilID,
			&kc.Dapil.Nama,
		)

		if err != nil {
			return nil, err
		}

		var kecamatan = map[string]any{
			"id":         kc.ID.Int64,
			"nama":       utility.Capitalize(kc.Nama.String),
			"dapil_id":   kc.DapilID.Int64,
			"nama_dapil": kc.Dapil.Nama.String,
		}

		daftarKecamatan = append(daftarKecamatan, kecamatan)
	}

	return daftarKecamatan, nil
}

func (kc *Kecamatan) FindKecamatanByDapil(dp int) ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var query string = `
		SELECT
			kc.id,
			kc.nama AS nama_kecamatan,
			kc.dapil_id,
			dp.nama AS nama_dapil
		FROM
			kecamatan kc
		JOIN
			dapil dp ON kc.dapil_id = dp.id
		WHERE 
			kc.dapil_id = ? AND kc.deleted_at IS NULL;
	`
	var daftarKecamatan []map[string]any

	rows, err := db.Raw(query, dp).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&kc.ID,
			&kc.Nama,
			&kc.DapilID,
			&kc.Dapil.Nama,
		)

		if err != nil {
			return nil, err
		}

		var kecamatan = map[string]any{
			"id":         kc.ID.Int64,
			"nama":       utility.Capitalize(kc.Nama.String),
			"dapil_id":   kc.DapilID.Int64,
			"nama_dapil": kc.Dapil.Nama.String,
		}

		daftarKecamatan = append(daftarKecamatan, kecamatan)
	}

	return daftarKecamatan, nil
}

func (kc *Kecamatan) TotalKecamatan() (int, error) {
	db := config.ConnectGormDB()
	var query string = "SELECT COUNT(*) AS total FROM kecamatan WHERE deleted_at IS NULL;"

	var total int // Variabel menampung jumlah dapil
	results := db.Raw(query).Scan(&total)
	if results.Error != nil {
		return -1, results.Error
	}

	return total, nil
}
