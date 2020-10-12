package version

import (
	"bufio"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"os"
	"strings"
)

type Version struct{}

func (v Version) Run(source cmd.Source, output *cmd.Output) {
	f, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	dfver := []string{}
	for _, eachline := range txtlines {
		if strings.Contains(eachline, "github.com/df-mc/dragonfly") {
			dfver = strings.Split(eachline, " ")
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	output.Printf("This server is running Dragonfly for Minecraft: Bedrock Edition v " + protocol.CurrentVersion + " implementing dragonfly version " + dfver[1])
}
