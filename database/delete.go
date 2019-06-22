package database

import (
	"database/sql"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/lib/pq"
)

func DeleteTodo(c *gin.Context) {
	param := c.Param("id")

	url := os.Getenv("host")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can't open", err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("DELETE from todos WHERE id=$1;")
	if err != nil {
		log.Fatal("prepare error", err.Error())
	}

	if _, err := stmt.Exec(param); err != nil {
		log.Fatal("prepare error", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{ "message": "Delete Successfully"})
}
