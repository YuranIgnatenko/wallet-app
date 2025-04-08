package db

import (
	"testing"
	"wallet-app/config"
)

func TestConnectDatabase(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	_, err := ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestMigrationReWriteTableWallets(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	conn, err := ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	err = MigrationReWriteTableWallets(conn)
	if err != nil {
		t.Error(err)
	}
}
