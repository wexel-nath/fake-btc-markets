package database

import (
	"database/sql"
	"errors"
	"fmt"

	"fake-btc-markets/pkg/config"

	_ "github.com/lib/pq"
)

var (
	client Client
)

func Connect() (*dbClient, error) {
	cfg := config.Get()
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.DatabaseUser,
		cfg.DatabasePass,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	))
	if err != nil {
		return nil, err
	}

	c := &dbClient{
		db: db,
	}

	return c, err
}

type dbClient struct {
	db *sql.DB
}

func (c *dbClient) IsInitialised() error {
	if c == nil || c.db == nil {
		return errors.New("db client is not initialised")
	}

	return nil
}

func (c *dbClient) Ping() error {
	if err := c.IsInitialised(); err != nil {
		return err
	}

	return c.db.Ping()
}

func (c *dbClient) Exec(query string, params []interface{}) error {
	if err := c.IsInitialised(); err != nil {
		return err
	}

	_, err := c.db.Exec(query, params...)
	return err
}

func (c *dbClient) QueryRow(query string, params []interface{}, columns []string) (map[string]interface{}, error) {
	if err := c.IsInitialised(); err != nil {
		return nil, err
	}

	row := c.db.QueryRow(query, params...)
	return scanRowToMap(row, columns)
}

func (c *dbClient) QueryRows(query string, params []interface{}, columns []string) ([]map[string]interface{}, error) {
	if err := c.IsInitialised(); err != nil {
		return nil, err
	}

	rows, err := c.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	return scanRowsToMap(rows, columns)
}

func getClient() (Client, error) {
	if client != nil && client.Ping() == nil {
		return client, nil
	}

	c, err := Connect()
	if err != nil {
		return nil, err
	}

	client = c
	return c, nil
}
