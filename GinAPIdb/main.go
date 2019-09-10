package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//DB pointer
var db *sql.DB

//Todo struct for keeping information
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Error!!! Unauthorization"})
		c.Abort()
		return
	}

	c.Next()

}

func postTodo(c *gin.Context) {
	t := Todo{}

	err := c.ShouldBindJSON(&t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "JSON parsing on insertion error!!! " + err.Error()})
		return
	}

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1,$2) RETURNING id", t.Title, t.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Insertion error!!! " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK"})
}

func getContainTodo(c *gin.Context) {

	titleQ := c.Query("title")

	todosLocal := Todo{}
	todos := []Todo{}
	found := false

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where title like $1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL for selection error!!! " + err.Error()})
		return
	}

	rows, err := stmt.Query("%" + titleQ + "%")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Query to select row error!!! " + err.Error()})
		return
	}

	for rows.Next() {
		err := rows.Scan(&todosLocal.ID, &todosLocal.Title, &todosLocal.Status)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "Fetch next row error!!! " + err.Error()})
			return
		}
		todos = append(todos, todosLocal)
		found = true
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error!!! Not found element to print"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func getOneTodo(c *gin.Context) {
	id := c.Param("id")

	var todoLocal Todo

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL select error!!! " + err.Error()})
		return
	}

	idnum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Convert id to num error!!!" + err.Error()})
		return
	}

	row := stmt.QueryRow(idnum)

	err = row.Scan(&todoLocal.ID, &todoLocal.Title, &todoLocal.Status)
	if err != nil {
		log.Println("Select id = " + id)
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": fmt.Sprintf("Select id %d error!!!: %s", idnum, err.Error())})
		return
	}

	c.JSON(http.StatusOK, todoLocal)
}

func deleteTodo(c *gin.Context) {

	id := c.Param("id")
	//Pre-select.
	var todoLocal Todo

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL select error!!! " + err.Error()})
		return
	}

	idnum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Convert id to num error!!!" + err.Error()})
		return
	}

	row := stmt.QueryRow(idnum)

	err = row.Scan(&todoLocal.ID, &todoLocal.Title, &todoLocal.Status)
	if err != nil {
		log.Println("Select id = " + id)
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": fmt.Sprintf("Select id %d error!!!: %s", idnum, err.Error())})
		return
	}

	//Actual delete.
	stmt, err = db.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL for delete row error!!! " + err.Error()})
		return
	}

	_, err = stmt.Exec(idnum)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Execute deletion error!!! " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	//Pre-select.
	var todoLocal Todo

	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id = $1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL select error!!! " + err.Error()})
		return
	}

	idnum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Convert id to num error!!! " + err.Error()})
		return
	}

	row := stmt.QueryRow(idnum)

	err = row.Scan(&todoLocal.ID, &todoLocal.Title, &todoLocal.Status)
	if err != nil {
		log.Println("Select id = " + id)
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": fmt.Sprintf("Select id %d error!!!: %s", idnum, err.Error())})
		return
	}

	//Actual update.
	t := Todo{}

	err = c.ShouldBindJSON(&t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "JSON parsing Error!!! " + err.Error()})
		return
	}

	stmt, err = db.Prepare("UPDATE todos SET title=$2,status=$3 WHERE id=$1")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Prepare SQL for update error!!! " + err.Error()})
		return
	}

	_, err = stmt.Exec(id, t.Title, t.Status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Execute update error!!! " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func main() {
	createTable()

	r := gin.Default()

	r.Use(authMiddleware)

	r.GET("/getcontain", getContainTodo)

	r.GET("getone/:id", getOneTodo)

	r.POST("/insert", postTodo)

	r.DELETE("/delete/:id", deleteTodo)

	r.PUT("/update/:id", updateTodo)

	r.Run(":1234")

	defer db.Close()
}

func createTable() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Connect to database error", err)
		return
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Println("Cannot create table.", err)
		return
	}

	fmt.Println("Successfully create table.")
}
