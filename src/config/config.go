package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parent := filepath.Dir(wd)
	err = godotenv.Load(filepath.Join(parent, ".env"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Load(name string) string {
	return os.Getenv(name)
}
