package cmd

import (
	"github.com/yenole/oh-my-env/storage"
)

type cmdEnv []string

func (cmd cmdEnv) Check() error {
	if len(cmd) == 1 || len(cmd) == 2 {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdEnv) Execute() error {
	store, err := storage.Load()
	if err != nil {
		return err
	}

	switch len(cmd) {
	case 1:
		store.Path(cmd[0])

	case 2:
		store.Env(cmd[0], cmd[1])

	default:
		return nil
	}
	return store.Flush()
}

func obtainCmdEnv(args []string) Cmd {
	return cmdEnv(args)
}
