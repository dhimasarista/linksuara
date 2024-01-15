package models

import (
	"database/sql"
	"linksuara/app/config"
	"linksuara/app/utility"
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

func (tps *TPS) FindTpsByKelurahan(kelurahanId int) ([]map[string]any, error) {
	db := config.ConnectGormDB()
	var query string = `
		SELECT
			tps.id,
			tps.nama,
			tps.kelurahan_id,
			tps.created_at,
			tps.updated_at,
			tps.deleted_at,
			kl.nama AS nama_kelurahan
		FROM
			tps
		JOIN
			kelurahan kl ON tps.kelurahan_id = kl.id
		WHERE tps.kelurahan_id = ? AND tps.deleted_at is NULL;
	`
	rows, err := db.Raw(query, kelurahanId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var daftarTps []map[string]any
	for rows.Next() {
		err = rows.Scan(
			&tps.ID,
			&tps.Nama,
			&tps.KelurahanID,
			&tps.CreatedAt,
			&tps.UpdatedAt,
			&tps.DeletedAt,
			&tps.Kelurahan.Nama,
		)
		if err != nil {
			return nil, err
		}

		var tps = map[string]any{
			"id":           tps.ID.Int64,
			"nama":         utility.CapitalizeAll(tps.Nama.String),
			"kelurahan_id": tps.KelurahanID.Int64,
			"created_at":   tps.CreatedAt,
		}

		daftarTps = append(daftarTps, tps)
	}

	return daftarTps, nil
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
