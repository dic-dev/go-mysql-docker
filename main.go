package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getUsers() []*User {
	db, err := sql.Open("mysql", "tester:password@tcp(db:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User
		err := results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}
	return users
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	json.NewEncoder(w).Encode(users)
}

func main() {
  // echo
  // e := echo.New()
  // e.GET("/", func(c echo.Context) error {
  //   return c.String(http.StatusOK, "Hello, World!")
  // })
  // e.Logger.Fatal(e.start(":8080"))

	http.HandleFunc("/", handler)
	http.HandleFunc("/users", usersPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
