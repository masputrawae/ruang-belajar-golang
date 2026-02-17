package utils

import (
	"embed"
	"io/fs"
)

func LoadStaticAssets(embedFS embed.FS, path string) ([]byte, error) {
	data, err := fs.ReadFile(embedFS, path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
