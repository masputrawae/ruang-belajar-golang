package data

import (
	"encoding/json"
	"os"
	"pr4lgn/internal/model/user"
)

var DataUserFile = "data/users.json"

func SaveUserData(d user.Users) error {
	jsonFile, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(DataUserFile, jsonFile, 0644); err != nil {
		return err
	}

	return nil
}

func LoadUserData() (user.Users, error) {
	var data user.Users

	f, err := os.ReadFile(DataUserFile)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(f, &data); err != nil {
		return data, err
	}

	return data, nil
}
