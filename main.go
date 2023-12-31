package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home page")
	fmt.Printf("hit the home page endpoint")
}

type User struct {
	ID int `json:"id"`
	Name string `json:"string"`
}

func getUsers() []*User {
	// Open up db connection (data source name for sql.Open should be the same as docker-compose.yml)
	db, err := sql.Open("mysql", "test_user:topSecretPassword@tcp(db:3306)/test_database")

	// chef if error with db connection
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Execute query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User

		// Scan each result
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}

	return users
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Hit the useres page endpoint")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", usersPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}