package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

//Endpoint of GetTasks
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get tasks from model
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

//Endpoint of PutTask
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Create a new task
		var task models.Task
		//
		c.Bind(&task)
		//Add task with help of models
		id, err := models.PutTask(db, task.Name)
		//Return response in JSON if success
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			//Errors handling
		} else {
			return err
		}
	}
}

//Endpoint of DeleteTask
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		//Use model to delete task
		_, err := models.DeleteTask(db, id)
		//Return response in JSON if success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
