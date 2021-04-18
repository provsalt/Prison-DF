package economy

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type SubTop string

func (t SubTop) SubName() string {
	return "top"
}

type Top struct {
	Sub  SubTop
	Page string
}

func (t Top) Run(source cmd.Source, output *cmd.Output) {
	// stmt, err := utils.EconomyDB.Database.Prepare("SELECT (username, money) FROM economy ORDER BY money DESC LIMIT 10 OFFSET 20")
	// if err != nil {
	// 	utils.GetLogger().Errorf(err.Error())
	// }
}
