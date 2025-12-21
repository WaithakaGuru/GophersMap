/* Exercise Thirty two
32. Add a logging middleware to your HTTP server.
*/

package exercise

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var userInfoMockDB = map[int]user{
	202: {202, 33, "Works at google"},
	300: {300, 29, "Works At Software Engineering Institute"},
	313: {313, 22, "Am a freelancer software Engineer"},
	529: {529, 23, "Graduate in Computer Science"},
}

type user struct {
	UserId   int    `json:"id"`
	UserAge  int    `json:"age"`
	UserInfo string `json:"info"`
}

func checkIdFieldMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// perform all the necessary checks here
		i := r.PathValue("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			s := err.Error() + " Incorrect request id"
			http.Error(w, s, http.StatusBadRequest)
			return
		}

		if i == "" {
			http.Error(w, "Id cannot be empty", http.StatusBadRequest)
			return
		}

		if _, found := userInfoMockDB[id]; !found {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		// call the next function if the checks are passed
		next(w, r)

	}
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	info, stringifyError := json.Marshal(userInfoMockDB[id])
	if stringifyError != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(info)
}

func StartUserServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the simple user server used to implement middleware in Golang")
	})

	http.HandleFunc("GET /info/{id}", checkIdFieldMiddleware(getUserInfo))

	log.Fatal(http.ListenAndServe(":4040", nil))
}
