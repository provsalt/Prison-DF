package punishment

import (
	"Prison/prisons/database"
	"database/sql"
	"time"
)

type Database *sql.DB

func New(credentials database.Credentials) Database {
	db, err := sql.Open("mysql", credentials.Username+":"+credentials.Password+"@("+credentials.IP+")/"+credentials.Schema)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS punishment(XUID BIGINT NOT NULL, punishment TEXT, endtime BIGINT, FOREIGN KEY (XUID) REFERENCES userinfo(XUID));")
	if err != nil {
		panic(err)
	}
	return db
}
