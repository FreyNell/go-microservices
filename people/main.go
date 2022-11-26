package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Sex  string `json:"sex"`
}

var DBNAME string = "people"

func getPeople(c *gin.Context) {
	db := createDB("people")

	var people []person
	rows, err := db.Query("SELECT id,name,age,sex FROM people;")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		var per person
		if err := rows.Scan(&per.ID, &per.Name, &per.Age, &per.Sex); err != nil {
			log.Fatal(err)
			return
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, people)
}

func postPeople(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error has occur"})
	}

	db := createDB("people")
	result, err := db.Exec("INSERT INTO people(name,age,sex) VALUES (?,?,?);", newPerson.Name, newPerson.Age, newPerson.Sex)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error has occur"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error has occur"})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Created ID: " + strconv.FormatInt(id, 10)})
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	var per person

	db := createDB("people")
	row := db.QueryRow("SELECT id,name,age,sex FROM people WHERE id = ?", id)
	defer db.Close()
	if err := row.Scan(&per.ID, &per.Name, &per.Age, &per.Sex); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found."})
		}
		log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "An error has ocurr."})
	}
	c.IndentedJSON(http.StatusOK, per)

}

func defaultResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "People Service")
}

func createDB(name string) (db *sql.DB) {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBIP") + ":" + os.Getenv("DBPORT"),
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	cfg = mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBIP") + ":" + os.Getenv("DBPORT"),
		DBName: name,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS people(" +
		"id INT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"name VARCHAR(30) NOT NULL," +
		"age INT," +
		"sex INT" +
		");")
	if err != nil {
		log.Fatal(err)
	}

	return
}

func main() {

	router := gin.Default()

	if os.Getenv("ENV") == "dev" {
		router.Use(cors.Default())
	}

	router.GET("/", defaultResponse)
	router.GET("/people", getPeople)
	router.GET("/people/:id", getPersonByID)
	router.POST("/people", postPeople)

	router.Run("0.0.0.0:8080")
}