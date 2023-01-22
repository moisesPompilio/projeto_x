package migrates

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func StarMigrate() {
	dbConnection, err := db.LoadDB_PostgreSQL()
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New(
		"file://src/pkg/migrates/users",
		dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
