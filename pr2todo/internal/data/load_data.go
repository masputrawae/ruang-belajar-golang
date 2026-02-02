package data

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrFailedCreateDir  = errors.New("gagal membuat direktori")
	ErrFailedCreateFile = errors.New("gagal membuat file")
	ErrInvalidJSON      = errors.New("format JSON tidak valid")
	ErrUnmarshalFailed  = errors.New("gagal mengubah JSON ke struct")
	ErrMarshalFailed    = errors.New("gagal mengubah struct ke JSON")
)

func Load[T any](filePath string) (*T, error) {

	// Cek, Buat direktori parent jika belum ada
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, errors.Join(ErrFailedCreateDir, err)
	}

	// Baca File
	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {

			// File baru → buat kosong & kembalikan struct default
			if err := os.WriteFile(filePath, nil, 0644); err != nil {
				return nil, errors.Join(ErrFailedCreateFile, err)
			}

			var zero T
			return &zero, nil
		}
		return nil, err
	}

	// File kosong → kembalikan struct default
	if len(dataBytes) == 0 {
		var zero T
		return &zero, nil
	}

	// Parse JSON ke struct
	var target T
	if err := json.Unmarshal(dataBytes, &target); err != nil {
		return nil, errors.Join(ErrUnmarshalFailed, ErrInvalidJSON, err)
	}

	return &target, nil
}

// Save menulis struct ke file JSON (dengan indent).
func Save[T any](filePath string, data T) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Join(ErrFailedCreateDir, err)
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Join(ErrMarshalFailed, err)
	}

	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return errors.Join(ErrFailedCreateFile, err)
	}

	return nil
}
