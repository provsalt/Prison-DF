package userinfo

import (
	"Prison/prisons/database"
	"database/sql"
	"time"
)

type Database *sql.DB

func New(credentials database.DatabaseCredentials) Database {
	db, err := sql.Open("mysql", credentials.Username+":"+credentials.Password+"@("+credentials.IP+")/"+credentials.Schema)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS userinfo(XUID BIGINT PRIMARY KEY , logins INT, join_time BIGINT, IPV4 TEXT);")
	if err != nil {
		panic(err)
	}
	return db
}
