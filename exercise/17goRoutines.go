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
	fmt.Println("  Finished slow count in ", time.Since(now))
}

func slowerCount() {
	now := time.Now()
	for range 1000 {
		time.Sleep(4 * time.Millisecond)
	}
	fmt.Println("  Finished slower count in ", time.Since(now))
}

func RunConcurrent() {
	// make a channel to ensure programs await execution of all routines
	chan1 := make(chan bool)
	chan2 := make(chan bool)

	// sequential function execution
	now := time.Now()
	fmt.Println("Running functions sequentially")
	slowCount()
	slowerCount()
	fmt.Println("Total sequential time: ", time.Since(now))

	// concurrent function execution
	better := time.Now()
	fmt.Println("Running function concurrently")
	go func() {
		slowCount()
		chan1 <- true
	}()
	go func() {
		slowerCount()
		chan2 <- true
	}()
	defer close(chan1)
	defer close(chan2)
	<-chan1
	<-chan2

	fmt.Println("Total concurrency time: ", time.Since(better))
}
