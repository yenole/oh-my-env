package main

import (
	"fmt"
	"github.com/yenole/oh-my-env/cmd"
	"os"
	"strings"
)

func Use() string {
	return ""
}

func main() {
	if len(os.Args) > 1 {
		c := strings.TrimSpace(os.Args[1])
		obtain, ok := cmd.Cmds[c]
		if !ok {
			goto help
		}
		cmd := obtain(os.Args[2:])
		if err := cmd.Check(); err != nil {
			fmt.Println(err)
		}

		if err := cmd.Execute(); err != nil {

		}
	}

help:
	fmt.Println(Use())

}
