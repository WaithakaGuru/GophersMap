/* Exercise Sixteen
16. Make an HTTP GET request to a public API and print the response.
*/

package exercise

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/users/"

func GetUsers() {
	// done := make(chan bool, )
	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("Fetch request to", URL, "failed")
		fmt.Println("Error: ", err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Something went wrong!!")
		fmt.Println("Error Status: ", response.Status)
		return
	}

	info, readError := io.ReadAll(response.Body)

	if readError != nil {
		fmt.Println("Error reading the reaponse: ", readError.Error())
		return
	}

	// var userData []map[string]any
	// json.Unmarshal(info, &userData)
	log.Print(string(info))
	// fmt.Println(userData)
}
