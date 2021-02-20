// +build darwin linux

package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var profile = `./.profile`

func (store *Storage) Profile() error {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	cfg := store.cfg(`default`)
	if len(cfg.Path) > 0 {
		paths := strings.Join(cfg.Path, ":")
		buffer.WriteString(fmt.Sprintf("export PATH=%s:$PATH\n", paths))
	}

	for k, v := range cfg.Envs {
		buffer.WriteString(fmt.Sprintf("export %s=%s\n", k, v))
	}

	for k, v := range cfg.Alias {
		buffer.WriteString(fmt.Sprintf("alias %s=\"%s\"\n", k, v))
	}

	return ioutil.WriteFile(profile, buffer.Bytes(), 0777)
}
