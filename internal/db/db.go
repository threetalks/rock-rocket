/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

*/

package db

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Initial(dsn string, driverName string, cfg *gorm.Config) error {
	var err error
	var dial gorm.Dialector
	switch driverName {
	case "mysql":
		dial = mysql.Open(dsn)
	case "sqlite":
		dial = sqlite.Open(dsn)
	case "postgresql", "postgres":
		dial = postgres.Open(dsn)
	case "sqlserver":
		dial = sqlserver.Open(dsn)
	case "clickhouse":
		dial = clickhouse.Open(dsn)
	}
	db, err = gorm.Open(dial, cfg)
	if err != nil {
		return err
	}
	return err
}

func DB() *gorm.DB {
	return db
}
