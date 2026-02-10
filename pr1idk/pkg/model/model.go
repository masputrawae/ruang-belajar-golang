package model

// Todo model
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

type DBConfig struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
