package model

type Config struct {
	App       App       `yaml:"app"`
	Server    Server    `yaml:"server"`
	DataFiles DataFiles `yaml:"dataFiles"`
}

type App struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	LanguageCode string `yaml:"languageCode"`
	Version      string `yaml:"version"`
	Copyright    string `yaml:"copyright"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DataFiles struct {
	User    string `yaml:"user"`
	Session string `yaml:"session"`
	Todo    string `yaml:"todo"`
}
