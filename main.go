package main

import (
	"fmt"
	"github.com/yenole/oh-my-env/cmd"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		c := strings.TrimSpace(os.Args[1])
		obtain, ok := cmd.Cmds[c]
		if !ok {
			goto help
		}
		_cmd := obtain(os.Args[2:])
		if err := _cmd.Check(); err != nil {
			fmt.Printf(`cmd:%v check:%v`, c, err)
			return
		}

		if err := _cmd.Execute(); err != nil {
			fmt.Printf(`cmd:%v execute:%v`, c, err)
		}
		return
	}

help:
	fmt.Println(cmd.UseHelp())

}
