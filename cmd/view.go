package cmd

import (
	"fmt"
	"github.com/yenole/oh-my-env/storage"
)

type cmdView []string

func (cmd cmdView) Check() error {
	if len(cmd) == 0 {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdView) Execute() error {
	store, err := storage.Load()
	if err != nil {
		return err
	}
	fmt.Println(store.View())
	return nil
}

func obtainCmdView(args []string) Cmd {
	return cmdView(args)
}
