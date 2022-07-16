package models

import (
	"database/sql"
	"fmt"
	"log"
	"todoapi/config"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error
var dbName = config.Config.DbName
var deployEnv = config.Config.Deploy

func init() {

	if deployEnv == "local" {
		Db, err = sql.Open("sqlite3", dbName)
		if err != nil {
			log.Fatalln(err)
		}

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, "todos")

		Db.Exec(cmdT)
	}

	if deployEnv == "prod" {
		Db, err = sql.Open("postgres", "host=postgresql.prod.svc.cluster.local port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
		if err != nil {
			log.Println(err)
		}

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id serial PRIMARY KEY,
		content text,
		user_id integer,
		created_at timestamp)`, "todos")

		Db.Exec(cmdT)
	}
}
