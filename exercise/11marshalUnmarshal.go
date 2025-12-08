/* Exercise Eleven
11. Marshal a `User` struct to JSON and unmarshal it back.
*/

package exercise

import "encoding/json"

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func Marshal(data User) ([]byte, error) {
	result, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func unmarshal(jsStr []byte) (User, error) {
	resultStruct := User{}
	err := json.Unmarshal(jsStr, &resultStruct)
	if err != nil {
		return resultStruct, err
	}
	return resultStruct, nil
}
