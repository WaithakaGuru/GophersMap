/* Exercise Seventeen
17. Run two functions concurrently using goroutines.
*/

package exercise

import (
	"fmt"
	"time"
)

func slowCount() {
	now := time.Now()
	for range 1000 {
		time.Sleep(3 * time.Millisecond)
	}
	fmt.Println("Finished slow count in ", time.Since(now), "time")
}

func slowerCount() {
	now := time.Now()
	for range 10000 {
		time.Sleep(3 * time.Millisecond)
	}
	fmt.Println("Finished slow count in ", time.Since(now), "time")
}

func RunConcurrent() {
	go slowCount()
	go slowerCount()
}
