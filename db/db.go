package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const (
	// drivers
	// https://github.com/golang/go/wiki/SQLDrivers
	DRIVER_MYSQL = "mysql"

	// default
	MYSQL_DEFAULT_PORT = 3306
)

// TODO : https://pkg.go.dev/github.com/go-playground/validator/v10
type Config struct {
	Driver   string
	Username string
	Pwd      string
	// Net      string // [tcp | unix]
	IP   string
	Port int
	DB   string
}

// username:password@protocol(address)/dbname?param=value
// https://github.com/go-sql-driver/mysql/
func (cfg *Config) ToMysqlDSN() string {
	if cfg.Port == 0 {
		cfg.Port = MYSQL_DEFAULT_PORT
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Pwd, cfg.IP, cfg.Port, cfg.DB)
}

type DB struct {
	cfg            *Config
	driverName     string
	dataSourceName string
	db             *sql.DB
}

func New(cfg *Config) (*DB, error) {
	var db *DB
	switch cfg.Driver {
	case DRIVER_MYSQL:
		db = &DB{
			driverName:     DRIVER_MYSQL,
			dataSourceName: cfg.ToMysqlDSN(),
		}
	default:
		return nil, errors.New("not supported dirver")
	}

	if sqlDB, err := sql.Open(db.driverName, db.dataSourceName); err != nil {
		return nil, err
	} else {
		db.db = sqlDB
	}

	return db, db.db.Ping()
}

func (db *DB) Connect(ctx context.Context) (*sql.Conn, error) {
	return db.db.Conn(ctx)
}

func (db *DB) Close() error {
	return db.db.Close()
}
