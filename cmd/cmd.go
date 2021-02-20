package cmd

import (
	"errors"
)

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

func UseHelp() string {
	return `Use:
	env init [bash|zsh]
	env env	[name] path
	env view
	env alias name [value]

Developers:
	auther:Yenole
	email:xusir92@gmail.com
	https://yenole.com
	https://github.com/yenole
	https://gitee.com/yenole
`
}
