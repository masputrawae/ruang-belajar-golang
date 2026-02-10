package main

import (
	"fmt"
	"log"
	"pr1idk/pkg/model/config"
	"pr1idk/pkg/model/db"
	"pr1idk/pkg/model/todo"
)

func main() {
	cfg := config.New(&config.API{JsonFile: "config.json"})
	db := db.New(cfg.Load().DB)
	defer db.Close()

	todo := todo.New(&todo.API{DB: db})
	// if err := todo.Create("Tugas Baru Ditambahkan"); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := todo.Create("Lebih Baru Lagi"); err != nil {
	// 	log.Fatal(err)
	// }

	todos, _ := todo.FindALL()
	todoID1, _ := todo.FindByID(2)
	jsons, _ := todo.ToJson()

	if err := todo.Delete(9); err != nil {
		log.Fatal(err)
	}

	if err := todo.Update("Ini sudah di ubah", true, 1); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo Find ALL", todos)
	fmt.Println("Todo Find BY ID", todoID1)
	fmt.Println("Todo To Json", string(jsons))
}
