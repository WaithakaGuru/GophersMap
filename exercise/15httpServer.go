/* Exercise Fifteen
15. Create an HTTP server that responds with "Welcome to Go!".
*/

package exercise

import (
	"net/http"
)

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go!"))
}

func CreateServer() {
	http.HandleFunc("/", handleWelcome)

	http.ListenAndServe(":8000", nil)
}
