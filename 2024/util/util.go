package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func DataPath(relativePath string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(pwd, relativePath)
}
