package main

import (
	"fmt"
	"os"
)

var (
	cmds = []Cmd{
		{Name: "init", Desc: `env init [bash|zsh]`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
		{Name: `add`, Desc: `env add [key] path`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
		{Name: `del`, Desc: `env del key`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
		{Name: `cfg`, Desc: `env cfg [name] [save|del]`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
		{Name: `view`, Desc: `env view`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
		{Name: `alias`, Desc: `env alias key [value]`, Option: func(strings []string) error {
			return nil
		}, Invoke: func(strings []string) error {
			return nil
		}},
	}
)

func init() {

}

func main() {
	fmt.Println(os.Args)
}
