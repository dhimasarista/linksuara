package models

import (
	"database/sql"
	"linksuara/app/config"
	"time"

	"gorm.io/gorm"
)

type Caleg struct {
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

func (cl *Caleg) GetByID(id int) error {
	db := config.ConnectGormDB()
	const query string = `
	SELECT
		cl.id,
		cl.nama AS caleg,
		cl.dapil_id,
		dp.nama AS nama_dapil
	FROM
		caleg cl
	JOIN
		dapil dp ON cl.dapil_id = dp.id
	WHERE 
		cl.id = ? AND cl.deleted_at IS NULL;
	`

	results := db.Raw(query, id).Scan(&cl)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (cl *Caleg) FindAllByDapil(id int) ([]map[string]any, error) {
	db := config.ConnectGormDB()
	const query string = `
	SELECT
		cl.id,
		cl.nama AS caleg,
		cl.dapil_id,
		dp.nama AS nama_dapil
	FROM
		caleg cl
	JOIN
		dapil dp ON cl.dapil_id = dp.id
	WHERE 
		cl.dapil_id = ? AND cl.deleted_at IS NULL;
	`

	rows, err := db.Raw(query, id).Rows()
	if err != nil {
		return nil, err
	}
	var daftarCaleg []map[string]any
	for rows.Next() {
		err = rows.Scan(
			&cl.ID,
			&cl.Nama,
			&cl.DapilID,
			&cl.Dapil.Nama,
		)
		if err != nil {
			return nil, err
		}
		var caleg = map[string]any{
			"id":         cl.ID.Int64,
			"nama":       cl.Nama.String,
			"dapil_id":   cl.DapilID.Int64,
			"nama_dapil": cl.Dapil.Nama.String,
		}

		daftarCaleg = append(daftarCaleg, caleg)
	}
	return daftarCaleg, nil
}
