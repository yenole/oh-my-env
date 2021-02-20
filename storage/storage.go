package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	storage = homeDir(`/data.data`)
)

type Storage struct {
	cfgs map[string]*cfg
}

func (store *Storage) cfg(key string) *cfg {
	if _, ok := store.cfgs[key]; !ok {
		store.cfgs[key] = newCfg()
	}
	return store.cfgs[key]
}

func (store *Storage) Env(key, value string) *Storage {
	if !filepath.IsAbs(value) {
		value, _ = filepath.Abs(value)
	}

	cfg := store.cfg(`default`)
	if len(value) == 0 {
		delete(cfg.Envs, key)
	} else {
		cfg.Envs[key] = value
	}
	return store
}

func (store *Storage) Alias(key, value string) *Storage {
	cfg := store.cfg(`default`)
	if len(value) == 0 {
		delete(cfg.Alias, key)
	} else {
		cfg.Alias[key] = value
	}
	return store
}

func (store *Storage) Path(value string) *Storage {
	if !filepath.IsAbs(value) {
		value, _ = filepath.Abs(value)
	}

	cfg := store.cfg(`default`)
	for i, v := range cfg.Path {
		if v == value {
			cfg.Path = append(cfg.Path[:i], cfg.Path[i+1:]...)
			return store
		}
	}
	cfg.Path = append(cfg.Path, value)
	return store
}

func (store *Storage) View() string {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	cfg := store.cfg(`default`)

	if len(cfg.Path)+len(cfg.Envs) > 0 {
		buffer.WriteString("envi:\n")
		for _, v := range cfg.Path {
			buffer.WriteString(fmt.Sprintf("\tPATH -> %v\n", v))
		}

		for k, v := range cfg.Envs {
			buffer.WriteString(fmt.Sprintf("\t%v=%v\n", k, v))
		}
	}

	if len(cfg.Alias) > 0 {
		buffer.WriteString("alias:\n")
		for k, v := range cfg.Alias {
			buffer.WriteString(fmt.Sprintf("\t%v=\"%v\"\n", k, v))
		}
	}
	return buffer.String()
}

func (store *Storage) Flush() error {
	if err := store.profile(); err != nil {
		return err
	}

	byts, err := json.Marshal(store.cfgs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(storage, byts, 0777)
}

func Load() (*Storage, error) {
	if _, err := os.Lstat(storage); os.IsNotExist(err) {
		return &Storage{cfgs: make(map[string]*cfg)}, nil
	}
	byts, err := ioutil.ReadFile(storage)
	if err != nil {
		return nil, err
	}
	var cfgs map[string]*cfg
	err = json.Unmarshal(byts, &cfgs)
	if err != nil {
		return nil, err
	}
	return &Storage{cfgs: cfgs}, nil
}

func homeDir(path string) string {
	dir := filepath.Dir(os.Args[0])
	return fmt.Sprint(dir, path)
}
