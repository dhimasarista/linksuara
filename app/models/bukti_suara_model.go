package models

import (
	"database/sql"
	"linksuara/app/config"
	"linksuara/app/utility"
	"time"

	"gorm.io/gorm"
)

type BuktiSuara struct {
	ID         sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Foto       []byte         `json:"foto" gorm:"column:foto"`
	Keterangan sql.NullString `json:"keterangan" gorm:"column:keterangan"`

	// Foreign Key
	JumlahSuaraID sql.NullInt64 `gorm:"column:jumlah_suara_id" json:"jumlah_suara_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (bs *BuktiSuara) NewData(filename, keterangan string) error {
	db := config.ConnectGormDB()

	// Membaca file yang ada di app/uploads/images/
	fileContent, err := utility.CompressFirst("./app/uploads/images/" + filename)
	if err != nil {
		return err
	}

	const query string = "INSERT INTO bukti_suara(id, foto, keterangan, jumlah_suara_id, created_at, updated_at, deleted_at) VALUES (0, ?, ?, NULL, NOW(), NOW(), NULL);"
	result := db.Exec(query, fileContent, keterangan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
