package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hanxin/school/database"
)

func main() {

	r := gin.Default()
	r.GET("/api/todos", database.GetTodos)
	r.Run(":1234")

}
