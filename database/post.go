package database 

import (
	"database/sql"
	"log"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func InsertDB(c *gin.Context) {
	var todo Todo 
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Wrong Request"})
	}

	url := os.Getenv("host")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("fatal", err.Error())
	}

	insertTB := `
		INSERT INTO todos(title, status) VALUES($1, $2) RETURNING ID;
	`
	var id int
	row := db.QueryRow(insertTB, todo.Title, todo.Status)
	err = row.Scan(&id)

	if err != nil {
		log.Fatal("Cannot Scan id", err.Error())
	}

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{ "message": "Insert Successfully"})
	
}