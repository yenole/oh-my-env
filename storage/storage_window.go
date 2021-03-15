// +build windows

package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var profile = homeDir(`/.profile`)

func (store *Storage) profile() error {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	cfg := store.cfg(`default`)
	if len(cfg.Path) > 0 {
		paths := strings.Join(cfg.Path, ";")
		buffer.WriteString(fmt.Sprintf("$ENV:PATH=$ENV:PATH + '%s;'\n", paths))
	}

	for k, v := range cfg.Envs {
		buffer.WriteString(fmt.Sprintf("$ENV:%s=%s\n", k, v))
	}

	for k, v := range cfg.Alias {
		buffer.WriteString(fmt.Sprintf("Invoke-Expression 'function %s { %s }'\n", k, v))
	}

	fmt.Println(">>", buffer.String())

	return ioutil.WriteFile(profile, buffer.Bytes(), 0777)
}
