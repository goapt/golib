package db

import (
	upperdb "github.com/verystar/db"
)

type Config struct {
	Enable       bool   `json:"enable"`
	Driver       string `json:"driver"`
	Dsn          string `json:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns" json:"max_idle_conns"`
	ShowSql      bool   `toml:"show_sql" json:"show_sql"`
}

// Error messages.
var (
	ErrNoMoreRows = upperdb.ErrNoMoreRows
)
