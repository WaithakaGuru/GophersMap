/* Exercise Twelve
12. Read a text file and print its contents.
*/

package exercise

import (
	"fmt"
	"os"
)

func ReadFile(path string) error {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return err
	}
	fmt.Println("File content: ", string(fileContent))
	return nil
}
