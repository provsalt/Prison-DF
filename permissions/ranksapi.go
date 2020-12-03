package permissions

import (
	"Prison/prisons/utils"
	"database/sql"
	"errors"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"time"
)

type Connection struct {
	IP       string
	Username string
	Password string
	Schema   string
}
type RankApi struct {
	Database *sql.DB
	Logger   *logrus.Logger
}

func New(connection Connection, minConn int, maxconn int, logger *logrus.Logger) RankApi {
	db, err := sql.Open("mysql", connection.Username+":"+connection.Password+"@("+connection.IP+")/"+connection.Schema)
	if err != nil {
		utils.Logger.Errorln(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(minConn)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS permissions(XUID BIGINT, username TEXT, PrisonRanks TINYINT DEFAULT 0, PaidRanks SMALLINT DEFAULT NULL, StaffZRanks SMALLINT DEFAULT NULL)")
	if err != nil {
		utils.Logger.Errorln(err)
	}
	return RankApi{Database: db, Logger: logger}
}

func (r RankApi) InitPlayer(player *player.Player) bool {
	row := r.Database.QueryRow("SELECT XUID FROM permissions WHERE username=?", player.Name())
	var XUID int
	err := row.Scan(&XUID)
	if err == nil {
		return true
	}
	if errors.Is(err, sql.ErrNoRows) {
		_, err := r.Database.Exec("REPLACE INTO permissions (XUID, username, PrisonRanks) VALUES (?, ?, ?)", player.XUID(), player.Name(), 0)
		if err != nil {
			r.Logger.Error(err)
			player.Disconnect(text.Colourf("Oopsie! \n <red>We have met an exception on our part. \nError code: RANKS MYSQL DATABASE EXCEPTION</red>"))
		}
	} else {
		panic(err)
	}
	return true
}

func GetPermissionLevel(player *player.Player) {
	// row := r.Database.QueryRow("SELECT XUID FROM permissions WHERE username=?", player.Name())

}
