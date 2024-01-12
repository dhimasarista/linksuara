package test

import (
	"linksuara/app/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit-Test koneksi ke Database
func TestConnectDB(t *testing.T) {
	db := config.ConnectGormDB()

	query := "SELECT 1;"
	results := db.Raw(query)

	assert.Nil(t, results.Error)
}
