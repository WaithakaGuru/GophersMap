/*	Exercise Thirteen
13. Write a string to a new file.
*/

package exercise

import (
	"errors"
	"os"
)

func WriteToNewFile(newPath, content string) error {
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, writeError := newFile.Write([]byte(content))
	if writeError != nil {
		return errors.New("Failed to write to file -> " + writeError.Error())
	}
	return nil
}
