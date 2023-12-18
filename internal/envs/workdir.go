package envs

import (
	"errors"
	"os"
	"path/filepath"
)

func ProjectDir() string {
	dir, _ := os.Getwd()
	find, _ := popFind(dir, "go.mod")
	return find
}

func popFind(path, target string) (string, error) {
	_, err := os.Stat(filepath.Join(path, target))
	if err != nil {
		parentDir := filepath.Dir(path)
		if path == "/" && parentDir == "/" {
			return "", errors.New("not found")
		}
		return popFind(parentDir, target)
	} else {
		return path, nil
	}
}
