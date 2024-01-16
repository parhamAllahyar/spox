package dev

import (
	"asset/config"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migration(config config.PostgresDBConfig) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DatabaseName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Read SQL migration file
	sqlFile := "asset/tables.sql"
	migrationSQL, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Execute migration SQL
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Migration from %s executed successfully\n", sqlFile)

}
