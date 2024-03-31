package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"iamricky.com/truck-rental/rootpath"
)

func Init(envFile string) {
	err := godotenv.Load(rootpath.Root + "/" + envFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Load(name string) string {
	return os.Getenv(name)
}
