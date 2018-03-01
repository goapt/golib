package db

import (
	upperdb "github.com/verystar/db"
)
type Config struct {
	Enable       bool
	Driver       string
	Dsn          string
	MaxOpenConns int  `toml:"max_open_conns"`
	MaxIdleConns int  `toml:"max_idle_conns"`
	ShowSql      bool `toml:"show_sql"`
}

// Error messages.
var (
	ErrNoMoreRows               = upperdb.ErrNoMoreRows
)