package database

import (
	"database/sql"
	"fmt"

	"fake-btc-markets/pkg/config"
	"fake-btc-markets/pkg/log"

	_ "github.com/lib/pq"
)

const (
	connectionFormat = "user=%s password=%s host=%s port=%d dbname=%s sslmode=disable"
)

var (
	connection *sql.DB
)

func GetConnection() *sql.DB {
	if connection == nil || connection.Ping() != nil {
		cfg := config.Get()
		c, err := sql.Open("postgres", fmt.Sprintf(
			connectionFormat,
			cfg.DatabaseUser,
			cfg.DatabasePass,
			cfg.DatabaseHost,
			cfg.DatabasePort,
			cfg.DatabaseName,
		))
		if err != nil {
			log.Error(err)
			return nil
		}

		connection = c
	}
	return connection
}
