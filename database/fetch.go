package database 

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func filter(db *sql.DB, c *gin.Context, id string) {
	command, err := db.Prepare("SELECT id, title, status FROM todos WHERE id=$1")

	if err != nil {
		log.Fatal("cannot connect from Db", err.Error())
	}

	num, _ := strconv.Atoi(id)
		row := command.QueryRow(num)
		log.Println(row)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"MessageError": "Cannot Query"})
			return
		}

		t := Todo{}
		err = row.Scan(&t.ID, &t.Title, &t.Status)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"MessageError": "Cannot Map todo"})
			return
		}

		c.JSON(http.StatusOK, t)
		return
}

func getAll(db *sql.DB, c *gin.Context) { 
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("cannot connect DB")
	}

	rows, _ := stmt.Query()

	var todos []Todo
	for rows.Next() {
		t := Todo{}

		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"MessageError": "Cannot Map Data to Struct"})
			return
		}
		todos = append(todos, t)
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodos(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("host"))
	if err != nil {
		log.Println(os.Getenv("host"))
		log.Fatal("cannot open DB")
	}

	queryParams := c.Request.URL.Query()
	id, ok := queryParams["id"]
	

	defer db.Close()

	if ok {
		filter(db, c, id[0])
		return
	} 

	getAll(db, c)
}