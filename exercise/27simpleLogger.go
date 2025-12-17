/*
	Exercise Twenty seven

27. Implement a simple logger that writes logs to a file with timestamps.
*/
package exercise

import (
	"io"
	"log"
	"os"
	"time"
)

// create a log that writes to multiple destinations - A FILE and THE CONSOLE
func Logger(txt, filePath string, i io.Writer) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open file %s", filePath)
	}

	defer file.Close()

	multW := io.MultiWriter(i, file)
	logger := log.New(multW, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)

	var RandomWords []string = []string{
		"Word", "is", "a", "The", "Food", "Road", "Area", "bag", "Time",
		"Month", "Great", "Movies", "Thief", "Field", "book",
	}

	logger.Println(txt)
	logger.Println("Here are your lucky words today: ")
	for i := range RandomWords {
		time.Sleep(time.Second)
		logger.Println(RandomWords[i])
	}

	return nil
}
