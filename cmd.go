package main

type Cmd struct {
	Name   string
	Desc   string
	Option func([]string) error
	Invoke func([]string) error
}
