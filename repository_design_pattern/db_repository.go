package repository_design_pattern

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	IDBRepository IDBRepository
}

func (db Database) ConnectDB() *gorm.DB {
	return db.IDBRepository.ConnectDB()
}

func (db Database) CloseDB(conn *gorm.DB) error {
	return db.IDBRepository.CloseDB(conn)
}
