package main

import (
	"fmt"
	"os"

	"iamricky.com/truck-rental/config"
	"iamricky.com/truck-rental/migrate"
	"iamricky.com/truck-rental/router"
)

func main() {
	config.Init()

	args := os.Args[1:]
	if len(args) >= 1 {
		if args[0] == "migrate" {
			migrate.Migrate()
		} else {
			fmt.Println("Invalid arguments.")
			fmt.Println("\nOptions:")
			fmt.Println("  migrate")
		}
	} else {
		router.Route()
	}
}
