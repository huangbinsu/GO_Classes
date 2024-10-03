package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func api() {
	var toDoMap []ToDo
	toDoMap = append(toDoMap, ToDo{Task: "Sample", Done: false})
	r := gin.Default()
	api := r.Group("/todo")

	// Home page
	api.GET(".", func(c *gin.Context) {
		c.JSON(http.StatusOK, toDoMap)
	})

	// Read / filte list with done or task name
	api.GET("/show", func(c *gin.Context) {
		status_input := c.Query("status")
		name_input := c.Query("name")
		if name_input == "" && status_input == "" {
			c.JSON(http.StatusOK, toDoMap)
		} else if status_input != "" && name_input == "" {
			status, err := strconv.ParseBool(status_input)
			if err != nil {
				c.String(200, "status error: %s\n", status_input)
				return
			}
			filted_map := filte_map(toDoMap, status)
			c.JSON(http.StatusOK, filted_map)
		} else if name_input != "" && status_input == "" {
			key, repeat := find_repeat(toDoMap, name_input)
			if repeat {
				c.JSON(http.StatusOK, toDoMap[key])
			} else {
				c.String(200, "task name: %s not found", name_input)
				return
			}
		}
	})

	// Add new task, can specify done status
	api.POST("/add", func(c *gin.Context) {
		newToDo := new(ToDo)
		name_input := c.DefaultPostForm("name", "Sample")
		_, repeat := find_repeat(toDoMap, name_input)
		if !repeat {
			newToDo.Task = name_input
			done := c.DefaultPostForm("done", "false")
			finish, err := strconv.ParseBool(done)
			if err != nil {
				c.String(200, "status error: %s\n", done)
				return
			}
			newToDo.Done = finish
			toDoMap = append(toDoMap, *newToDo)
		} else {
			c.String(200, "Duplicate task name: %s\n", name_input)
			return
		}
		c.JSON(http.StatusOK, toDoMap)
	})

	// Update done status, can specify or invert state
	api.PUT("/update/:name", func(c *gin.Context) {
		name_input := c.Param("name")
		status_input := c.Query("status")
		key, repeat := find_repeat(toDoMap, name_input)
		if repeat {
			temp := toDoMap[key]
			if status_input != "" {
				if status_input == "true" {
					temp.Done = true
				} else if status_input == "false" {
					temp.Done = false
				} else {
					c.String(200, "status error: %s\n", status_input)
					return
				}
			} else {
				temp.Done = !temp.Done
			}
			toDoMap[key] = temp
			c.JSON(http.StatusOK, toDoMap)
		} else {
			c.String(200, "task name: %s not found", name_input)
			return
		}
	})

	// Delete task
	api.DELETE("/del/:name", func(c *gin.Context) {
		name_input := c.Param("name")
		key, repeat := find_repeat(toDoMap, name_input)
		if repeat {
			toDoMap = append(toDoMap[:key], toDoMap[key+1:]...)
		} else {
			c.String(200, "task name: %s not found", name_input)
			return
		}
		c.JSON(http.StatusOK, toDoMap)
	})

	r.Run(":800")
}

func find_repeat(toDoMap []ToDo, name string) (int, bool) {
	key := 0
	for key, value := range toDoMap {
		if value.Task == name {
			return key, true // 假設每個 Task 值唯一，找到後可以直接退出循環
		}
	}
	return key, false
}

func filte_map(toDoMap []ToDo, done bool) []ToDo {
	var filted_list []ToDo
	for _, value := range toDoMap {
		if value.Done == done {
			filted_list = append(filted_list, value) // 假設每個 Task 值唯一，找到後可以直接退出循環
		}
	}
	return filted_list
}
