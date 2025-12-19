/* Exercise Twenty Eight
28. Create a RESTful API for managing a list of tasks (CRUD operations).
*/

package exercise

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var TaskLock = sync.Mutex{}
var TaskReadLock = sync.RWMutex{}

var MockTaskDB map[int]task = map[int]task{
	1: {Id: 1, Title: "Walk in the park", Complete: false},
	2: {Id: 2, Title: "Clean utensils", Complete: false},
	3: {Id: 3, Title: "Read the bible", Complete: false},
	4: {Id: 4, Title: "Code a task to lear GOlang", Complete: false},
}

type task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"completed"`
}

// func to mimic data fetching from a DB
func getTasks() *map[int]task {
	TaskReadLock.RLock()
	defer TaskReadLock.RUnlock()
	return &MockTaskDB
}

// default/ primary route
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to task Controller API"))
}

// funct to get all tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	jsString, err := json.Marshal(*getTasks())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusFound)
	w.Write(jsString)
}

// func to get a specific task by its id
func getTask(w http.ResponseWriter, r *http.Request) {

}

// Func to add a task
// func addTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	ta := task{}
// 	err := json.NewDecoder(r.Body).Decode(&ta)

// 	if err != nil {
// 		http.Error(
// 			w, err.Error(), http.StatusBadRequest,
// 		)
// 	}

// 	tasks := getTasks()
// 	ta = *tasks
// 	TaskLock.Lock()

// 	tasks
// 	TaskLock.Unlock()

// 	tJson, _ := json.Marshal(&tasks)
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(tJson)
// }

// func to handle delete endpoint
func deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
	}
	tasks := []task{}
	for _, item := range MockTaskDB {
		if item.Id == id {
			continue
		}
		tasks = append(tasks, item)
	}

	// set JSON params
	w.Header().Set("Content-Type", "application/json")

	str, stringfyError := json.Marshal(tasks)
	if stringfyError != nil {
		http.Error(w, stringfyError.Error(), http.StatusInternalServerError)
	}

	// write a successful delete status to the header
	w.WriteHeader(http.StatusNoContent)

	// return the updated body
	w.Write(str)
}

// func to handle task status updating
// func updateTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	taskIdToUpdate, err :=  strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	newTask := task{}

// 	decodeError := json.NewDecoder(r.Body).Decode(&newTask)

// }

// will use built in http module to create a multiplexer
func TasksServer() {
	var mux = http.NewServeMux()

	mux.HandleFunc("GET /", hello)
	mux.HandleFunc("GET /tasks", getAllTasks)
	mux.HandleFunc("GET /task/{id}", getTask)
	// mux.HandleFunc("POST /task", addTask)
	mux.HandleFunc("DELETE /task/{id}", deleteTask)
	// mux.HandleFunc("PATCH /task/{id}", updateTask)

	http.ListenAndServe(":5050", mux)
}
