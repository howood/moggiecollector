package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/howood/moggiecollector/library/utils"
)

// sql and database info.
const (
	Source        = "file://./database/migrations/"
	dbPort string = "5433"
)

//nolint:gochecknoglobals
var (
	command = flag.String("cmd", "", "up or down")
	force   = flag.Bool("f", false, "force execute sql")
)

//nolint:forbidigo
func main() {
	flag.Parse()

	if len(*command) < 1 {
		fmt.Println("")
		fmt.Println("Error : no argument")
		fmt.Println("")
		os.Exit(1)

		return
	}

	addr := net.JoinHostPort(os.Getenv("DB_HOSTNAME"), utils.GetOsEnv("DB_PORT", dbPort))
	database := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		addr,
		os.Getenv("DB_DBNAME"),
	)

	fmt.Println("DSN:", database)

	m, err := migrate.New(Source, database)
	if err != nil {
		fmt.Println("err", err)
	}

	version, dirty, err := m.Version()
	showVersion(version, dirty, err, "current")

	fmt.Println("command: cmd", *command)
	apply(m, version, dirty)
}

//nolint:forbidigo
func apply(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		//nolint:gosec
		if err := m.Force(int(version)); err != nil {
			fmt.Println("err", err)
			os.Exit(1)
		}
	}

	var err error

	switch *command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		// do nothing
		return
	default:
		fmt.Println("")
		fmt.Println("Error : invalid command > ", *command)
		fmt.Println("")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)

		return
	}

	fmt.Println("success:", *command)

	version, dirty, err = m.Version()
	showVersion(version, dirty, err, "updated")
}

//nolint:forbidigo
func showVersion(version uint, dirty bool, err error, msg string) {
	fmt.Println("____/:::", msg, ":::\\_______")
	fmt.Println("version :", version)
	fmt.Println("dirty   :", dirty)
	fmt.Println("error   :", err)
	fmt.Println("_____________________________")
}
