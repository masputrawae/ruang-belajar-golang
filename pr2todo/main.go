package main

import (
	"fmt"
	"pr2todo/cmd/app"
	"pr2todo/internal/config"
)

func main() {
	cfg, log := config.Load("config.yaml")

	fmt.Println("PR2TODO Server")
	fmt.Printf("\n%s\n", log)

	app.Server(cfg.Server.Port, cfg.Server.Host)
}
