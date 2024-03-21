package main

import (
	"database/sql"
	"errors"
	"time"
)

func readDB() error {
	sqlinfo, ok := config["sql"].(string)
	if !ok {
		return errors.New("Assertion failed")
	}
	if dbtmp, err := sql.Open("mysql", sqlinfo); err != nil {
		return err
	} else {
		db = dbtmp
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
