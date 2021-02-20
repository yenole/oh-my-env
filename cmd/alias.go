package cmd

import "github.com/yenole/oh-my-env/storage"

type cmdAlias []string

func (cmd cmdAlias) Check() error {
	if len(cmd) == 1 || len(cmd) == 2 {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdAlias) Execute() error {
	store, err := storage.Load()
	if err != nil {
		return err
	}
	if len(cmd) == 1 {
		store.Alias(cmd[0], ``)
	} else {
		store.Alias(cmd[0], cmd[1])
	}
	return store.Flush()
}

func obtainCmdAlias(args []string) Cmd {
	return cmdAlias(args)
}
