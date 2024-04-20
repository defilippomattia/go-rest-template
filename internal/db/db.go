package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Credentials struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetConnectionPool(credentials Credentials) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		credentials.Host, credentials.Port, credentials.User, credentials.Password, credentials.Database)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
