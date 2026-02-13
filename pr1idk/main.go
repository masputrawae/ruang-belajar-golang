package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"pr1idk/cmd/server"
	"pr1idk/internal/configs"
	"pr1idk/internal/databases"
	"pr1idk/internal/repositories"
	"pr1idk/internal/services"
)

func main() {
	serve := flag.Bool("serve", true, "todo server")
	flag.Parse()

	cfg := configs.New("config.toml")
	db := databases.New(cfg.DB)

	defer db.Close()

	repoUser := repositories.NewUser(db)
	// repoTodo := repositories.NewTodo(db)

	serviceUser := services.NewUserService(repoUser)
	if err := serviceUser.UserCreate(context.Background(), "@john", "john_password", "john@example.com"); err != nil {
		log.Fatal(err)
	}
	ok, err := serviceUser.UserAuth(context.Background(), "@john", "john_password")
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("Sukses Login")
	}

	if *serve {
		log.Printf("👉️ Server running: http:%s%s\n", cfg.Server.Host, cfg.Server.Addr)
		if err := server.New(cfg, db); err != nil {
			log.Fatal(err)
		}
	}
}
