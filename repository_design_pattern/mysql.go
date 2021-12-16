package repository_design_pattern

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type MySQLRepository struct {
	DatabaseDialect string
	DatabaseURL     string
}

func (c *MySQLRepository) ConnectDB() *gorm.DB {
	connection, err := gorm.Open(c.DatabaseDialect, c.DatabaseURL)

	if err != nil {
		panic("failed to connect database")
	}

	if err = connection.DB().Ping(); err != nil {
		log.Fatalln("failed to ping database")
	}

	log.Println("connected to database")
	return connection
}

func (c *MySQLRepository) CloseDB(connection *gorm.DB) error {
	if err := connection.Close(); err != nil {
		return errors.New("cannot close current database")
	}
	log.Println("database closed successfully")
	return nil
}
