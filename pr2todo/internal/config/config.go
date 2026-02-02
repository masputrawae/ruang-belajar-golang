package config

import (
	"os"
	"pr2todo/internal/model"

	"gopkg.in/yaml.v3"
)

var (
	FailedLoad  = "❌ Gagal Mengambil/Membaca File Konfigurasi\n✅ Menggunakan Konfigurasi Default"
	SuccessLoad = "✅ Konfigurasi, Berhasil Digunakan"
)

// Membaca File Konfigurasi
func Load(filepath string) (model.Config, string) {
	data := model.Config{
		App: model.App{
			Name:         "Todo App",
			Description:  "Aplikasi web sederhana untuk mengelola tugas",
			LanguageCode: "id",
			Version:      "v.0.0.1",
			Copyright:    "Copyright (c) 2026 Putra Jaya. All Rights Reserved.",
		},
		Server: model.Server{
			Host: "localhost",
			Port: ":8080",
		},
		DataFiles: model.DataFiles{
			Session: "data/sessions.json",
			User:    "data/users.json",
			Todo:    "data/todos.json",
		},
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		return data, FailedLoad
	}

	if err := yaml.Unmarshal(file, &data); err != nil {
		return data, FailedLoad
	}

	return data, SuccessLoad
}
