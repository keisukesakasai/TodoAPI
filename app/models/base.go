package models

import (
	"database/sql"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error
var tracer = otel.Tracer("TodoAPI-models")

/*
func init() {
	fmt.Println("initializing...")
	Db, err = sql.Open("sqlite3", config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, "todos")

	Db.Exec(cmdT)

	log.Println("initializing...DONE!!!!")
}
*/

func init() {
	fmt.Println("Now migration...")
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

	fmt.Println("Now migration...DONE!!")

	log.Println("initializing...DONE!!!!")
}

/*
func createUUID(c *gin.Context) (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}
*/
