package cmd

import "errors"

type Cmd interface {
	Check() error
	Execute() error
}

var (
	Cmds = map[string]func([]string) Cmd{
		`init`:  obtainCmdInit,
		`env`:   obtainCmdEnv,
		`view`:  obtainCmdView,
		`alias`: obtainCmdAlias,
	}
	errArgsWrong = errors.New(`args wrong`)
)
