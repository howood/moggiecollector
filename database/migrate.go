package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/howood/moggiecollector/library/utils"
)

//sql and database info
const (
	Source = "file://./database/migrations/"
)

var (
	Command  = flag.String("cmd", "", "up or down")
	Force    = flag.Bool("f", false, "force execute sql")
	Database = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("YUGABYTEDB_USER"),
		os.Getenv("YUGABYTEDB_PASSWORD"),
		os.Getenv("YUGABYTEDB_HOSTNAME"),
		utils.GetOsEnvInt("YUGABYTEDB_PORT", 5433),
		os.Getenv("YUGABYTEDB_DBNAME"),
	)
)

func main() {

	flag.Parse()
	if len(*Command) < 1 {
		fmt.Println("")
		fmt.Println("Error : no argument")
		fmt.Println("")
		os.Exit(1)
		return
	}

	m, err := migrate.New(Source, Database)
	if err != nil {
		fmt.Println("err", err)
	}
	version, dirty, err := m.Version()
	showVersion(version, dirty, err, "current")

	fmt.Println("command: cmd", *Command)
	apply(m, version, dirty)
}

func apply(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *Force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	var err error
	switch *Command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		//do nothing
		return
	default:
		fmt.Println("")
		fmt.Println("Error : invalid command > ", *Command)
		fmt.Println("")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *Command)
		version, dirty, err := m.Version()
		showVersion(version, dirty, err, "updated")
	}
}

func showVersion(version uint, dirty bool, err error, msg string) {
	fmt.Println("____/:::", msg, ":::\\_______")
	fmt.Println("version :", version)
	fmt.Println("dirty   :", dirty)
	fmt.Println("error   :", err)
	fmt.Println("_____________________________")
}
