package config

import (
	"fmt"
	"time"

	_ "time/tzdata"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Mysql *gorm.DB
}

func Connection(user, password, host, port, dbname string) *DB {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		user,
		password,
		host,
		port,
		dbname,
		"Asia%2FJakarta",
	)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	pl, _ := db.DB()
	if err != nil {
		panic(err)
	}
	pl.SetMaxIdleConns(10)
	pl.SetMaxOpenConns(80)
	pl.SetConnMaxLifetime(time.Hour)
	return &DB{Mysql: db}
}
