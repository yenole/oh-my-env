package cmd

type cmdEnv []string

func (cmd cmdEnv) Check() error {
	if len(cmd) == 1 || len(cmd) == 2 {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdEnv) Execute() error {
	return nil
}

func obtainCmdEnv(args []string) Cmd {
	return cmdEnv(args)
}
