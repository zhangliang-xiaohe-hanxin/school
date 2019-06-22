package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hanxin/school/database"
	"os"
)

func main() {

	r := gin.Default()
	r.GET("/api/todos", database.GetTodosHandler)
	r.POST("/api/todos", database.InsertDB)
	r.DELETE("/api/todos/:id", database.DeleteTodo)
	r.Run(os.Getenv("port"))

}
