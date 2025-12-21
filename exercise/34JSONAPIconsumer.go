/*
consumes the JSON Data API
*/

package exercise

import (
	"fmt"
	"net/http"
	"os"
)

func GetData(url string) {
	// Switch on / spin up / start the data server

	// make a Get request to the server
	info, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := *info
	data.Write(os.Stdout)
	fmt.Println()
}
