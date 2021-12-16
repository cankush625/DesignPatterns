package repository_design_pattern

import "github.com/jinzhu/gorm"

type IDBRepository interface {
	ConnectDB() *gorm.DB
	CloseDB(connection *gorm.DB) error
}
