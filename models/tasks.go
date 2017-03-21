package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//Struct with data of the
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//List of goals
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	//Exit if SQL not worked somehow
	if err != nil {
		panic(err)
	}
	//Make sure that everything closes when you exit the program
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		//Exit on error
		err2 := rows.Scan(&task.ID, task.Name)
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"
	//Execute the SQL query
	stmt, err := db.Prepare(sql)
	//Exit on error
	if err != nil {
		panic(err)
	}
	//Make sure that everything closes when you exit the program
	defer stmt.Close()

	// replace the '?' In the query for 'name'
	result, err2 := stmt.Exec(name)
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	//Execute the SQL query
	stmt, err := db.Prepare(sql)
	//Exit on error
	if err != nil {
		panic(err)
	}
	// replace the '?' In the query for 'id'
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
