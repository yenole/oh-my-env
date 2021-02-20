package storage

type cfg struct {
	Path  []string
	Envs  map[string]string
	Alias map[string]string
}

func newCfg() *cfg {
	return &cfg{Path: make([]string, 0, 10), Envs: make(map[string]string), Alias: make(map[string]string)}
}
