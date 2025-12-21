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
	// make a Get request to the server
	info, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := *info
	data.Write(os.Stdout)
}
