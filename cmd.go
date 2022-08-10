package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

type (
	Commands struct {
		App *cli.App
	}
)

var (
	PrefixName = "APP"
	Commit     string
	Name       = "godzilla"
)

func NewCmd() *Commands {
	cmd := &Commands{
		App: cli.NewApp(),
	}

	cmd.App.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "c,config",
			Value:   "config/config.yaml",
			Usage:   "server config path",
			EnvVars: cmd.name("CONFIG"),
		},
		&cli.StringFlag{
			Name:    "apollo",
			Value:   "config/apollo.json",
			Usage:   "server apollo config path",
			EnvVars: cmd.name("APOLLO_CONFIG"),
		},
	}

	cmd.App.Name = Name
	cmd.App.Version = Commit
	cmd.App.Compiled = time.Now()

	return cmd
}

func (cmd *Commands) name(name string) []string {
	return []string{strings.ToUpper(fmt.Sprintf("%s_%s", PrefixName, name))}
}
