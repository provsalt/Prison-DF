package restart

import (
	"Prison/prisons/utils"
	"Prison/tasks"
	"time"
)

func Restartcheck() {
	tasks.Schedulerepeatingtask(func() {
		min := utils.GetServer().Uptime().Minutes()
		switch {
		case float64(time.Hour) <= min:
			err := utils.GetServer().Close()
			if err != nil {
				panic(err)
			}
		case float64(time.Minute*55) >= min, float64(time.Minute*56) >= min, float64(time.Minute*57) >= min, float64(time.Minute*58) >= min, float64(time.Minute*59) >= min:
			players := utils.GetServer().Players()
			left := 60 - utils.GetServer().Uptime().Minutes()
			for _, p := range players {
				p.Messagef("Server will restart in %.0f rejoin to play after restart", left)
			}
		}
		utils.GetLogger().Debugf("Server uptime %.0f", min)
	}, time.Minute)
}
