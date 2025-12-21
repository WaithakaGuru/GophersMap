/*	Thirty Four
34. Consume a JSON API, parse the response, and display selected fields.
*/

package exercise

import (
	"log"
	"net/http"
	"os"
)

func readJSONData(path string) ([]byte, bool, error) {
	byteInfo, err := os.ReadFile(path)

	if err != nil {
		return []byte{}, false, err
	}
	return byteInfo, true, nil
}

func basePath(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the Root Endpoint: Electronic Data in JSON format"))
}

func getFullInfo(w http.ResponseWriter, r *http.Request) {
	// read the json file from storage
	byteInfo, success, err := readJSONData("./exercise/34data.json")

	if !success {
		http.Error(w, "Failed to read JSON file -> "+err.Error(), http.StatusInternalServerError)
		return
	}

	// write json data as response
	w.WriteHeader(http.StatusFound)
	w.Write(byteInfo)
}

// func getSetOfInfo(w http.ResponseWriter, r *http.Request) {
// 	dataCount, err := strconv.Atoi(r.PathValue("count"))
// 	if err != nil {
// 		http.Error(w, err.Error() + " -> Invalid count value", http.StatusBadRequest)
// 		return
// 	}

// 	byteInfo, _, _ := readJSONData("./exercise/34data.json")
// 	g := string(byteInfo)
// 	json.Indent()
// }

func DataAPI() {
	// server the API via a multiplexer
	dataMux := http.NewServeMux()

	dataMux.HandleFunc("/", basePath)
	dataMux.HandleFunc("GET /data", getFullInfo)
	// dataMux.HandleFunc("GET /data/{count}", getSetOfInfo)

	log.Fatal(http.ListenAndServe(":3220", dataMux))
}
