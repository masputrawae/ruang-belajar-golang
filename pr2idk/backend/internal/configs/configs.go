package configs

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	Database DBConfig     `toml:"database"`
	Server   ServerConfig `toml:"server"`
}

type DBConfig struct {
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Option   string `toml:"option"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Addr string `toml:"addr"`
}

func Load(fn string) Configs {
	var configs Configs

	f, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	if err = toml.Unmarshal(f, &configs); err != nil {
		log.Fatal(err)
	}

	return configs
}
