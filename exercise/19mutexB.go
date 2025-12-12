// Similar implementation of the mutex but in a much simpler way using waitgroup

package exercise

import (
	"fmt"
	"sync"
	"time"
)

var waitG = sync.WaitGroup{}
var mut = sync.Mutex{}

func slowInc(c *Counter) {
	time.Sleep(3 * time.Second)
	mut.Lock()
	defer mut.Unlock()
	c.count += 1
	fmt.Println("Added '1' to count -> ", c.count)
	waitG.Done()
}

func fastInc(c *Counter) {
	time.Sleep(1 * time.Second)
	mut.Lock()
	defer mut.Unlock()
	c.count += 2
	fmt.Println("Added '2' to count -> ", c.count)
	waitG.Done()
}

func CounterIncrementHandler() {
	var counter2 = createCounter()
	waitG.Add(20)

	for range 10 {
		go func() {
			slowInc(counter2)
		}()

		go func() {
			fastInc(counter2)
		}()
	}
	waitG.Wait()
}
