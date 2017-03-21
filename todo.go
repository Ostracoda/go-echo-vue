package main

import (
	"database/sql"
	"go-echo-vue/handlers"

	"github.com/labstack/echo"
	//  "github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create Database
	db := initDB("dataset.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start as a web server
	//	e.Run(standard.New(":8000"))
	e.Logger.Fatal(e.Start(":8000"))
}

// Create method to work with database
func initDB(filepath string) *sql.DB {
	//Open connection or create it
	db, err := sql.Open("sqlite3", filepath)

	//Check for errors
	if err != nil {
		panic(err)
	}

	// If errors not found, but can't connect to db
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
  CREATE TABLE IF NOT EXISTS tasks(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL
  );
  `
	// Exit if errors witn SQL occurs above
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
