package configs

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	DB     DB     `toml:"db"`
	Server Server `toml:"server"`
}

type DB struct {
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Option   string `toml:"option"`
}

type Server struct {
	Host string `toml:"host"`
	Addr string `toml:"addr"`
}

func New(fileName string) Configs {
	var configs Configs
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err := toml.Unmarshal(f, &configs); err != nil {
		log.Fatal(err)
	}
	return configs
}
