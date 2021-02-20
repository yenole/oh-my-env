package cmd

type cmdAlias []string

func (cmd cmdAlias) Check() error {
	return nil
}

func (cmd cmdAlias) Execute() error {
	return nil
}

func obtainCmdAlias(args []string) Cmd {
	return cmdAlias(args)
}
