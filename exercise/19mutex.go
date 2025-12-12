/* Exercise nineteen
19. Create a counter that can be safely incremented by multiple goroutines using a mutex.
*/

package exercise

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
}

func createCounter() *Counter {
	return &Counter{}
}

var mutex = sync.Mutex{}

func slowIncrement(c *Counter) {
	time.Sleep(3 * time.Second)
	mutex.Lock()
	c.count += 1
	mutex.Unlock()
	fmt.Println("Slow increment: Increased count by 1 to ->", c.count)
}

func fastIncrement(c *Counter) {
	time.Sleep(1 * time.Second)
	mutex.Lock()
	c.count += 2
	mutex.Unlock()
	fmt.Println("Fast increment: Increased count by 2 to ->", c.count)
}

func HandleIncrements() {
	slowIncrementChan := make(chan bool, 10)
	fastIncrementChan := make(chan bool, 10)
	counter1 := createCounter()

	for range 10 {
		go func() {
			slowIncrement(counter1)
			slowIncrementChan <- true
		}()

		go func() {
			fastIncrement(counter1)
			fastIncrementChan <- true
		}()
	}

	for range 10 {
		<-slowIncrementChan
		<-fastIncrementChan
	}
}
