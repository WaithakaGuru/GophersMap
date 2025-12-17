/* Exercise Twenty Five
25. Write a function that takes an interface and demonstrate dependency injection.
*/

package exercise

import (
	"fmt"
	"time"
)

// Dependency injection is where by a type like a struct uses info (dependency) from an external source

// Mock database
var WorkersDB []userInfo = []userInfo{
	{id: 101, name: "Waithaka Amos", age: 24, occupation: "System Engineer", salary: 240_567.44, hobbies: []string{
		"Reading Tech Blogs", "Roasting bad coding practices", "Reviewing public codebases", "Playing Scrabble",
	}},
	{id: 201, name: "Mutheu Anitah", age: 21, occupation: "Software Developer", salary: 188_455.90, hobbies: []string{
		"Writing code blogs", "Mountain Hiking", "Singing", "Teaching how to code",
	}},
	{id: 306, name: "Steve Hawkings", age: 35, occupation: "System designer", salary: 340_677.65, hobbies: []string{
		"Write best coding practices", "Play Guitar", "Host Tech Events",
	}},
	{id: 202, name: "Jane Switt", age: 29, occupation: "UI/UX designer", salary: 155_673.90, hobbies: []string{
		"Draw interfaces", "Paint arts", "Riding bicycle", "Attending Hackatons",
	}},
	{id: 303, name: "Johan Grant", age: 22, occupation: "Backend developer", salary: 270_567.40, hobbies: []string{
		"Write Go blogs", "Create custom packages", "Discuss new tech on podcasts", "Reading Hero books",
	}},
	{},
}

// user infomation description
type userInfo struct {
	id         int
	name       string
	age        int
	occupation string
	salary     float32
	hobbies    []string
}

type UserData interface {
	getInfo(id int) userInfo
}

// a struct that implements the UserData interface
type Fetcher struct{}

func (f Fetcher) getInfo(id int) userInfo {
	fmt.Println("---- Fetching Data for ID: ", id, "-----")
	time.Sleep(1 * time.Second)
	for _, item := range WorkersDB {
		if id == item.id {
			return item
		}
	}
	return userInfo{}
}

// a struct that requires/uses/consumes  UserData interface
type WorkerModel struct {
	id   int
	info UserData
}

// WorkerModel methods
func (model WorkerModel) DisplayWorkerInfo() {
	var fetchedInfo = model.info.getInfo(model.id)
	fmt.Println("Worker Id: ", model.id)
	fmt.Println("Worker Name: ", fetchedInfo.name)
	fmt.Println("Worker Age: ", fetchedInfo.age)
	fmt.Println("Worker Occupation: ", fetchedInfo.occupation)
	fmt.Println("Worker Salary: ", fetchedInfo.salary)
	fmt.Printf("Worker '%s' loves to do the following:\n ", fetchedInfo.name)
	fmt.Printf("%v", fetchedInfo.hobbies)
}

// injecting the dependency via a constructor method
func NewWorker(i UserData, id int) *WorkerModel {
	return &WorkerModel{info: i, id: id}
}
