package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// #1 creating ToDo model
type ToDo struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Date        string `json: "date"`
}

// #2 creating ToDo slice
var todos []ToDo

// #5 Creating getItems function
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// #6 Creating deleteItems function
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range todos {

		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

// #7 Cretaing getitem function
func getItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range todos{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
// #8 Creating createitem function
func createItem (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var todo ToDo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// #9 Creating update method
func updateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range todos{
		if item.ID == params["id"]{
			todos = append(todos [:index], todos[index+1:]...)
			var todo ToDo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func main() {

	// #3 Register url paths and handlers
	r := mux.NewRouter()

	// Creating local data to todo lice
	todos = append(todos, ToDo{ID: "1", Title: "Task1", Description: "Cut the grass", Date: "2021/2/1"})
	todos = append(todos, ToDo{ID: "2", Title: "Task2", Description: "Buy anniversary gift", Date: "2022/6/2"})
	r.HandleFunc("/todo", createItem).Methods("POST")
	r.HandleFunc("/todo/{id}", getItem).Methods("GET")
	r.HandleFunc("/todo/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteItem).Methods("DELETE")
	r.HandleFunc("/todo", getItems).Methods("GET")

	// #4 Created local server
	fmt.Printf("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
