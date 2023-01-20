package command

import (
	"github.com/mitchellh/cli"
)

var (
	_ cli.Command = ()
)

type AgentCommand struct {
	*Base
}