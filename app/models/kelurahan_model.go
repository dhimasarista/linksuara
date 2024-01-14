package models

import (
	"database/sql"
	"linksuara/app/config"
	"linksuara/app/utility"
	"time"

	"gorm.io/gorm"
)

type Kelurahan struct {
	ID   sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Nama sql.NullString `json:"nama" gorm:"column:nama"`

	// Foreign Key
	Kecamatan   Kecamatan     `gorm:"foreignKey:KecamatanID" json:"kecamatan"`
	KecamatanID sql.NullInt64 `gorm:"column:kecamatan_id" json:"kecamatan_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (kl *Kelurahan) FindAll() ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var query string = `
		SELECT
			kl.id,
			kl.nama AS nama_kelurahan,
			kl.kecamatan_id,
			kc.nama AS nama_kecamatan
		FROM
			kelurahan kl
		JOIN
			kecamatan kc ON kl.kecamatan_id = kc.id
		WHERE kl.id > 1 AND kl.deleted_at IS NULL;
	`
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	var daftarKelurahan []map[string]any

	for rows.Next() {
		err = rows.Scan(
			&kl.ID,
			&kl.Nama,
			&kl.KecamatanID,
			&kl.Kecamatan.Nama,
		)

		if err != nil {
			return nil, err
		}

		var kelurahan = map[string]any{
			"id":             kl.ID.Int64,
			"nama":           utility.Capitalize(kl.Nama.String),
			"kecamatan_id":   kl.KecamatanID.Int64,
			"nama_kecamatan": kl.Kecamatan.Nama.Valid,
		}

		daftarKelurahan = append(daftarKelurahan, kelurahan)
	}
	return daftarKelurahan, nil
}

func (kl *Kelurahan) FindKelurahanByKecamatan(kc int) ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var query string = `
		SELECT
			kl.id,
			kl.nama AS nama_kelurahan,
			kl.kecamatan_id,
			kc.nama AS nama_kecamatan
		FROM
			kelurahan kl
		JOIN
			kecamatan kc ON kl.kecamatan_id = kc.id
		WHERE kc.id = ? AND kl.deleted_at IS NULL;
		`
	rows, err := db.Raw(query, kc).Rows()
	if err != nil {
		return nil, err
	}
	var daftarKelurahan []map[string]any
	for rows.Next() {
		err = rows.Scan(
			&kl.ID,
			&kl.Nama,
			&kl.KecamatanID,
			&kl.Kecamatan.Nama,
		)

		if err != nil {
			return nil, err
		}

		var kelurahan = map[string]any{
			"id":           kl.ID.Int64,
			"nama":         utility.Capitalize(kl.Nama.String),
			"kecamatan_id": kl.KecamatanID.Int64,
		}

		daftarKelurahan = append(daftarKelurahan, kelurahan)
	}

	return daftarKelurahan, nil
}
