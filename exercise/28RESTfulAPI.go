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

// func to help automate the response stringifying and writing
// Marshals the data and handle any errors that may occur otherwise write response of the new data to the client
func HelperWriteResponse[info map[int]task | task | any](w http.ResponseWriter, data info, status int) {
	// set JSON params
	w.Header().Set("Content-Type", "application/json")

	dataStr, dataStringifyErr := json.Marshal(data)
	if dataStringifyErr != nil {
		http.Error(w, dataStringifyErr.Error(), http.StatusInternalServerError)
		return
	}

	// write a successful http status code to the header
	w.WriteHeader(status)

	// return the updated body to the client 'browser'
	w.Write(dataStr)
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
	tasks := getTasks()
	HelperWriteResponse(w, *tasks, http.StatusFound)
}

// func to get a specific task by its id
func getTask(w http.ResponseWriter, r *http.Request) {
	taskId, idConvertErr := strconv.Atoi(r.PathValue("id"))
	if idConvertErr != nil {
		http.Error(w, idConvertErr.Error(), http.StatusBadRequest)
		return
	}

	TaskLock.Lock()
	tasks := *getTasks()
	task, ok := tasks[taskId]
	TaskLock.Unlock()

	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	HelperWriteResponse(w, task, http.StatusFound)
}

// Func to add a task
func addTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ta := task{}
	if err := json.NewDecoder(r.Body).Decode(&ta); err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
		return
	}

	t := getTasks()
	tasks := *t

	TaskLock.Lock()
	tasks[len(tasks)+1] = ta
	t = &tasks
	TaskLock.Unlock()

	HelperWriteResponse(w, tasks, http.StatusCreated)
}

// func to handle delete endpoint
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
	}

	t := getTasks()
	ta := *t
	if _, found := ta[id]; !found {
		http.Error(w, "Task to be deleted was not Found", http.StatusBadRequest)
		return
	}

	TaskLock.Lock()
	delete(ta, id)
	t = &ta
	TaskLock.Unlock()

	HelperWriteResponse(w, ta, http.StatusNoContent)
}

// func to handle task status updating
func updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type status struct {
		Complete bool `json:"complete"`
	}
	newTaskStatus := status{}

	if decodeError := json.NewDecoder(r.Body).Decode(&newTaskStatus); decodeError != nil {
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
	}

	TaskLock.Lock()
	t := getTasks()
	ta := *t

	taskToUpdate, found := ta[taskId]
	if !found {
		http.Error(w, "Task to be updated was not found!", http.StatusNotModified|http.StatusNotFound)
	}

	taskToUpdate.Complete = newTaskStatus.Complete
	ta[taskId] = taskToUpdate
	t = &ta
	TaskLock.Unlock()

	HelperWriteResponse(w, ta, http.StatusOK)
}

// will use built in http module to create a multiplexer
func TasksServer() {
	var mux = http.NewServeMux()

	mux.HandleFunc("GET /", hello)
	mux.HandleFunc("GET /tasks", getAllTasks)
	mux.HandleFunc("GET /task/{id}", getTask)
	mux.HandleFunc("POST /task", addTask)
	mux.HandleFunc("DELETE /task/{id}", deleteTask)
	mux.HandleFunc("PATCH /task/{id}", updateTaskStatus)

	http.ListenAndServe(":5050", mux)
}
