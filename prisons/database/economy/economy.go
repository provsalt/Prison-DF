package economy

import (
	"Prison/prisons/database"
	"database/sql"
	"errors"
	"github.com/df-mc/dragonfly/server/player"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Economy struct {
	Database *sql.DB
}

func New(connection database.Credentials, minConn int, maxconn int) Economy {
	db, err := sql.Open("mysql", connection.Username+":"+connection.Password+"@("+connection.IP+")/"+connection.Schema)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(minConn)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(XUID BIGINT, username TEXT, money FLOAT, FOREIGN KEY (XUID) REFERENCES userinfo(XUID));")
	if err != nil {
		panic(err)
	}
	return Economy{
		Database: db,
	}
}

func (e Economy) InitPlayer(player *player.Player, defaultmoney int) bool {
	r := e.Database.QueryRow("SELECT XUID FROM economy WHERE username=?", player.Name())
	var XUID int
	err := r.Scan(&XUID)
	if err == nil {
		return true
	}
	if errors.Is(err, sql.ErrNoRows) {
		_, err := e.Database.Exec("REPLACE INTO economy (XUID, username, money) VALUES (?, ?, ?)", player.XUID(), player.Name(), defaultmoney)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return true
}
func (e Economy) Close() {
	err := e.Database.Close()
	if err != nil {
		panic(err)
	}
}

func (e Economy) Balance(player *player.Player) (uint, error) {
	r := e.Database.QueryRow("SELECT money FROM economy WHERE XUID=?", player.XUID())
	var money uint
	err := r.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

func (e Economy) BalanceFromName(player string) (uint, error) {
	r := e.Database.QueryRow("SELECT money FROM economy WHERE username=?", player)
	var money uint
	err := r.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

func (e Economy) AddMoney(player *player.Player, amount uint) error {
	bal, err := e.Balance(player)
	if err != nil {
		return err
	}
	_, err = e.Database.Exec("REPLACE INTO economy (money) VALUES (?)", bal+amount)
	if err != nil {
		return err
	}
	return nil
}

func (e Economy) ReduceMoney(player *player.Player, amount uint) error {
	bal, err := e.Balance(player)
	if err != nil {
		return err
	}
	_, err = e.Database.Exec("REPLACE INTO economy (money) VALUES (?)", bal-amount)
	if err != nil {
		return err
	}
	return nil
}

func (e Economy) SetMoney(player *player.Player, amount uint) error {
	_, err := e.Database.Exec("REPLACE INTO economy (money) VALUES (?) ", amount)
	if err != nil {
		return err
	}
	return nil
}
