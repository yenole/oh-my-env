package cmd

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	help = `Use:
	env init %v
	env env	[name] path
	env view
	env alias name [value]

Developers:
	auther:yenole
	email:xusir92@gmail.com
	https://yenole.com
	https://github.com/yenole
	https://gitee.com/yenole`
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
	if runtime.GOOS == `windows` {
		return fmt.Sprintf(help, ``)
	}
	return fmt.Sprintf(help, `[bash|zsh]`)
}
