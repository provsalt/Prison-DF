package ranks

import (
	"Prison/prisons/database"
	"database/sql"
	"errors"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"time"
)

type RankApi struct {
	Database *sql.DB
	Logger   *logrus.Logger
}

type Ranks struct {
	PrisonRanks int
	PaidRanks   int
	StaffRanks  int
}

func New(connection database.Credentials, minConn int, maxconn int, logger *logrus.Logger) RankApi {
	db, err := sql.Open("mysql", connection.Username+":"+connection.Password+"@("+connection.IP+")/"+connection.Schema)
	if err != nil {
		logger.Errorln(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(minConn)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ranks(XUID BIGINT, username TEXT, PrisonRanks INT DEFAULT 0, PaidRanks INT DEFAULT 0, StaffRanks INT DEFAULT 0, FOREIGN KEY ranks(XUID) REFERENCES userinfo(XUID))")
	if err != nil {
		logger.Errorln(err)
	}
	return RankApi{Database: db, Logger: logger}
}

func (r RankApi) InitPlayer(player *player.Player) bool {
	row := r.Database.QueryRow("SELECT XUID FROM ranks WHERE username=?", player.Name())
	var XUID int
	err := row.Scan(&XUID)
	if err == nil {
		return true
	}
	if errors.Is(err, sql.ErrNoRows) {
		_, err := r.Database.Exec("REPLACE INTO ranks (XUID, username, PrisonRanks) VALUES (?, ?, ?)", player.XUID(), player.Name(), 0)
		if err != nil {
			r.Logger.Error(err)
			player.Disconnect(text.Colourf("Oopsie! \n <red>We have met an exception on our part. \nError code: RANKS MYSQL DATABASE EXCEPTION</red>"))
		}
	} else {
		panic(err)
	}
	return true
}

func (r RankApi) GetPermissionLevel(player *player.Player) Ranks {
	row := r.Database.QueryRow("SELECT PrisonRanks, PaidRanks, StaffRanks FROM ranks WHERE XUID=?", player.XUID())
	var ranks Ranks
	err := row.Scan(&ranks.PrisonRanks, &ranks.PaidRanks, &ranks.StaffRanks)
	if err != nil {
		player.Disconnect("An error has occured. Contact staff for more detailss.")
		r.Logger.Errorf(err.Error())
	}
	return ranks
}
