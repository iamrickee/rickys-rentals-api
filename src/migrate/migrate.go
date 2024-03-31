package migrate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"iamricky.com/truck-rental/db"
	"iamricky.com/truck-rental/errhandler"
	"iamricky.com/truck-rental/rootpath"
)

func Migrate() {
	var invalid = false
	args := os.Args[2:]
	if len(args) >= 1 {
		cmd := args[0]
		if cmd == "create" {
			if len(args) == 2 {
				name := args[1]
				create(name)
			} else {
				invalid = true
			}
		} else if cmd == "down" {
			if len(args) == 1 {
				down(0)
			} else if len(args) == 2 {
				count, err := strconv.Atoi(args[1])
				if err == nil {
					down(count)
				} else {
					invalid = true
				}
			} else {
				invalid = true
			}
		} else if cmd == "up" {
			if len(args) == 1 {
				up(0)
			} else if len(args) == 2 {
				count, err := strconv.Atoi(args[1])
				if err == nil {
					up(count)
				} else {
					invalid = true
				}
			} else {
				invalid = true
			}
		} else {
			invalid = true
		}
	} else {
		invalid = true
	}
	if invalid {
		fmt.Println("Invalid arguments.")
		fmt.Println("\nUsage:")
		fmt.Println("  create <name>")
		fmt.Println("  down [<amount>]")
		fmt.Println("  up [<amount>]")
	}
}

func create(name string) {
	migrationsDir := getMigrationsDir()
	var index int = 0
	filepath.Walk(migrationsDir, func(path string, info fs.FileInfo, err error) error {
		parts := strings.Split(filepath.Base(path), "_")
		if len(parts) >= 1 {
			prefix := parts[0]
			i, err := strconv.Atoi(prefix)
			if err == nil {
				if i > index {
					index = i
				}
			}
		}
		fmt.Println(filepath.Base(path))
		return nil
	})
	index++
	up := filepath.Join(migrationsDir, strconv.Itoa(index)+"_"+name+".up.sql")
	down := filepath.Join(migrationsDir, strconv.Itoa(index)+"_"+name+".down.sql")
	var content []byte
	errUp := os.WriteFile(up, content, 0755)
	errhandler.Fatal(errUp, true)
	errDown := os.WriteFile(down, content, 0755)
	errhandler.Fatal(errDown, true)
	fmt.Println("Created migrations: " + strconv.Itoa(index) + "_" + name)
}

func down(amount int) {
	db, _ := db.GetConn()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	errhandler.Fatal(err, true)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+getMigrationsDir(),
		"mysql",
		driver,
	)
	errhandler.Fatal(err, true)
	if amount > 0 {
		err = m.Steps(-amount)
	} else {
		err = m.Down()
	}
	errhandler.Fatal(err, true)
	fmt.Println("Successfully Migrated Down")
}

func up(amount int) {
	db, _ := db.GetConn()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	errhandler.Fatal(err, true)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+getMigrationsDir(),
		"mysql",
		driver,
	)
	errhandler.Fatal(err, true)
	if amount > 0 {
		err = m.Steps(amount)
	} else {
		err = m.Up()
	}
	errhandler.Fatal(err, true)
	fmt.Println("Successfully Migrated Up")
}

func getMigrationsDir() string {
	migrationsDir := filepath.Join(rootpath.Root, "migrations")
	os.Mkdir(migrationsDir, 0644)
	return migrationsDir
}
