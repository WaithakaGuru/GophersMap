/* Exercise Twenty six
26. Read and print an environment variable.
*/

package exercise

import (
	"fmt"

	// import a third party package for loading the environment variables
	// E "godotenv"

	"os"

	E "github.com/joho/godotenv"
)

func LoadEnv() {
	// use a third party package to load the env
	E.Load(".env")
	// use os package to get the env
	port, found := os.LookupEnv("PORT")
	if !found {
		fmt.Println("Variable not found")
		return
	}
	fmt.Println("Variable:", port)
}
