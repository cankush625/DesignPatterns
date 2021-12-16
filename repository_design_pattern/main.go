package repository_design_pattern

import "os"

func PostgresDB() Database {
	postgres := Database{IDBRepository: &PostgreSQLRepository{DatabaseDialect: "postgres",
		DatabaseURL: os.Getenv("POSTGRES_URL")}}
	return postgres
}
