package cmd

type cmdInit []string

func (cmd cmdInit) Check() error {
	if ok := (len(cmd) == 1 && (cmd[0] == `bash` || cmd[0] == `zsh`)) || (len(cmd) == 0); ok {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdInit) Execute() error {
	return nil
}

func obtainCmdInit(args []string) Cmd {
	return cmdInit(args)
}
