/* Exercise Twenty Eight
28. Create a RESTful API for managing a list of tasks (CRUD operations).
*/

package exercise

import (
	"encoding/json"
	"net/http"
)

var MockTaskDB []task = []task{
	{Id: 1, Title: "Walk in the park", Complete: false},
	{Id: 2, Title: "Clean utensils", Complete: false},
	{Id: 3, Title: "Read the bible", Complete: false},
	{Id: 4, Title: "Code a task to lear GOlang", Complete: false},
}

type task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"completed"`
}

// func to mimic data fetching from a DB
func getTasks() []task {
	var Tasks []task = MockTaskDB
	return Tasks
}

// default/ primary route
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to task Controller API"))
}

// funct to get all tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	jsString, _ := json.Marshal(getTasks())
	w.Write(jsString)
}

// Func to add a task
func addTask(w http.ResponseWriter, r *http.Request) {
	ta := task{}
	err := json.NewDecoder(r.Body).Decode(&ta)

	if err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
	}

	tasks := getTasks()
	tasks = append(tasks, ta)
	tJson, _ := json.Marshal(tasks)
	w.Write(tJson)
}

// will use built in http module to create a multiplexer
func TasksServer() {
	var mux = http.NewServeMux()

	mux.HandleFunc("GET /", hello)
	mux.HandleFunc("GET /tasks", getAllTasks)
	mux.HandleFunc("POST /task", addTask)

	http.ListenAndServe(":5050", mux)
}
