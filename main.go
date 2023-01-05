package main

import (
	"fmt"
	"net/http"
)

type user struct {
	name   string
	age    int
	gender string
}

var users []user = []user{
	{"Alice", 25, "Female"},
	{"Bob", 30, "Male"},
	{"Eve", 35, "Female"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	for _, values := range users {
		fmt.Println(values)
	}
}

func main() {

	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8000", nil)
}
