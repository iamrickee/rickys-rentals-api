package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"iamricky.com/truck-rental/rootpath"
)

var hasInit bool = false

func Init(envFile string) {
	err := godotenv.Load(filepath.Clean(rootpath.Root + "/" + envFile))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hasInit = true
}

func Load(name string) string {
	if !hasInit {
		Init(".env")
	}
	return os.Getenv(name)
}
