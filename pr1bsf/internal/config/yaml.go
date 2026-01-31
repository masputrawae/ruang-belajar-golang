package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server     Server `yaml:"server"`
	DataDir    string `yaml:"dataDir"`
	ContentDir string `yaml:"contentDir"`
	UserData   string `yaml:"userData"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var (
	FailedLoadConfig = "Failed to load configuration. Using Default"
)

func LoadConfig(file string) Config {
	data := Config{
		Server: Server{
			Host: "localhost",
			Port: "8080",
		},
		DataDir:    "data",
		ContentDir: "content",
		UserData:   "users.json",
	}

	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(FailedLoadConfig, "\n", err)
		return data
	}

	if err := yaml.Unmarshal(f, &data); err != nil {
		fmt.Println(FailedLoadConfig, "\n", err)
		return data
	}

	fmt.Println("Success load config:", file)
	return data
}
