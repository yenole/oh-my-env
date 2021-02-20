package cmd

type cmdView []string

func (cmd cmdView) Check() error {
	if len(cmd) == 0 {
		return nil
	}
	return errArgsWrong
}

func (cmd cmdView) Execute() error {
	return nil
}

func obtainCmdView(args []string) Cmd {
	return cmdView(args)
}
