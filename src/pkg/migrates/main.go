package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func main() {
	dbConnection, err := db.Migrates_LoadDB_PostgreSQL()
	if err != nil {
		fmt.Println(err)
	}
	m, err := migrate.New(
		"file://src/pkg/migrates/users",
		dbConnection)
	if err != nil {
		fmt.Println(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
}
