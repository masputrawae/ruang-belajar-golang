package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"pr3idk/internal/repository"
	"pr3idk/internal/service"
	"pr3idk/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error: ", err)
	}
	DB_DSN := os.Getenv("DB_DSN")
	// SERVER_ADDR := os.Getenv("SERVER_ADDR")

	db := utils.ConnectDB(DB_DSN)
	defer db.Close()

	repo := repository.New(db)
	service := service.New(repo)

	username := "Putra Jaya"
	password := "putra_password"

	// SUCCESS
	// if err := service.UserAdd(context.Background(), &username, &password); err != nil {
	// 	log.Fatal(err)
	// }

	// SUCCESS
	ok, err := service.UserAuth(context.Background(), username, password)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("Success")
	}

	// SUCCESS
	// task := "Coba insert tugas kedua mungkin"
	// err = service.TodoAdd(context.Background(), 5, &task, nil, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// SUCCESS
	todos, err := service.TodoFindByUserID(context.Background(), 5)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.MarshalIndent(todos, "", "    ")
	fmt.Println(string(data))

	// SUCCESS
	// if err := service.TodoDelete(context.Background(), 1); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("🚀 Server running: http://localhost%s\n", SERVER_ADDR)
	// if err := http.ListenAndServe(SERVER_ADDR, nil); err != nil {
	// 	log.Fatal(err)
	// }
}
