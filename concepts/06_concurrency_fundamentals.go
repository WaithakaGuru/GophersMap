/*
CONCURRENCY IN GO - Goroutines and Channels Fundamentals
============================================================================

Concurrency is Go's secret weapon. It's what makes Go special.
Goroutines are lightweight threads managed by the Go runtime.
Channels allow safe communication between goroutines.

TOPICS COVERED:
- Goroutines: creating and running
- Channels: sending and receiving data
- WaitGroup: synchronizing goroutines
- Race conditions and solutions
- Common concurrency patterns
- Goroutine leaks prevention
*/

package concepts

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// 1. GOROUTINES - Lightweight Concurrency
// ============================================================================

/*
GOROUTINES:
- Extremely lightweight threads (thousands per program)
- Managed by Go runtime (not OS threads)
- Cost: minimal memory overhead
- Creation: just add 'go' keyword before function call
- Go runtime handles scheduling across CPU cores

IMPORTANT:
- Main function is also a goroutine
- Program exits when main goroutine returns
- Other goroutines are killed when main exits
- Always ensure goroutines complete before main returns
*/

func printNumbers(name string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("%s: Done!\n", name)
}

func GoroutinesBasicsDemo() {
	fmt.Println("\n========== GOROUTINES BASICS ==========")

	fmt.Println("--- Sequential Execution (Slow) ---")
	start := time.Now()
	fmt.Println("Starting sequential execution...")
	printNumbers("Task1", 3)
	printNumbers("Task2", 3)
	elapsed := time.Since(start)
	fmt.Printf("Total time: %.2f seconds\n\n", elapsed.Seconds())

	fmt.Println("--- Concurrent Execution (Fast) ---")
	start = time.Now()
	fmt.Println("Starting concurrent execution...")
	// Create goroutines with 'go' keyword
	go printNumbers("Task1", 3)
	go printNumbers("Task2", 3)

	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	elapsed = time.Since(start)
	fmt.Printf("Total time: %.2f seconds\n", elapsed.Seconds())
	fmt.Println("Notice: concurrent is faster!")

	fmt.Println("\n✓ Goroutines basics demonstrated")
}

// ============================================================================
// 2. CHANNELS - Communication Between Goroutines
// ============================================================================

/*
CHANNELS:
- Typed conduit for communication between goroutines
- Created with make(chan Type)
- Send: ch <- value
- Receive: value := <-ch
- Both send and receive block until operation completes
- Safer than shared memory (no race conditions)

BUFFERED VS UNBUFFERED:
- Unbuffered: ch := make(chan int)
  └─ Goroutines block on send/receive
  └─ Must be matched with receiver/sender

- Buffered: ch := make(chan int, 5)
  └─ Can hold 5 elements
  └─ Send blocks only when full
  └─ Receive blocks only when empty
*/

func channelSender(ch chan string, count int) {
	for i := 1; i <= count; i++ {
		message := fmt.Sprintf("Message %d", i)
		fmt.Printf("Sending: %s\n", message)
		ch <- message // Send to channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Always close when done
}

func ChannelsBasicsDemo() {
	fmt.Println("\n========== CHANNELS BASICS ==========")

	// 1. Unbuffered channel
	fmt.Println("--- Unbuffered Channel ---")
	ch := make(chan string)

	go channelSender(ch, 3)

	// Receive from channel
	for message := range ch {
		fmt.Printf("Received: %s\n", message)
	}

	// 2. Buffered channel
	fmt.Println("\n--- Buffered Channel ---")
	buffered := make(chan int, 3) // Can hold 3 elements

	// Send multiple values without receiver
	fmt.Println("Sending to buffered channel...")
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("All sent! (no blocking)")

	// Now receive
	fmt.Println("Receiving from buffered channel...")
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	close(buffered)

	fmt.Println("\n✓ Channels basics demonstrated")
}

// ============================================================================
// 3. WAITGROUP - Synchronizing Goroutines
// ============================================================================

/*
WAITGROUP:
- Counts how many goroutines are in progress
- Add(n): increment counter
- Done(): decrement counter (call when goroutine finishes)
- Wait(): block until counter reaches zero

IMPORTANT:
- Better than manually managing channels for simple cases
- Prevents main from exiting before goroutines complete
- Clean and idiomatic
*/

func workerTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Automatically called when goroutine exits

	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d: task %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Worker %d: finished\n", id)
}

func WaitGroupDemo() {
	fmt.Println("\n========== WAITGROUP ==========")

	var wg sync.WaitGroup

	// Create 4 workers
	numWorkers := 4

	// Add 4 to the counter
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Launch goroutine
		go workerTask(i, &wg)
	}

	fmt.Println("Main: waiting for all workers to complete...")

	// Block until counter reaches zero
	wg.Wait()

	fmt.Println("Main: all workers completed!")

	fmt.Println("\n✓ WaitGroup demonstrated")
}

// ============================================================================
// 4. RACE CONDITIONS - The Danger of Shared Memory
// ============================================================================

/*
RACE CONDITIONS:
- Occur when multiple goroutines access shared data simultaneously
- Can cause unpredictable behavior
- Hard to debug because timing-dependent

SOLUTIONS:
1. Use channels (preferred in Go)
2. Use sync.Mutex for critical sections
3. Use atomic operations
4. Use channels to isolate data
*/

// PROBLEM: Race condition with shared counter
func RaceConditionProblemDemo() {
	fmt.Println("\n--- PROBLEM: Race Condition ---")

	var counter int
	var wg sync.WaitGroup

	// Launch 5 goroutines that increment counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Read, increment, write (not atomic!)
			counter++ // RACE CONDITION!
		}()
	}

	wg.Wait()

	// Expected: 5, but might be less due to race condition
	fmt.Printf("Counter (should be 5): %d\n", counter)
	fmt.Println("Notice: might not be 5 due to race condition!")
}

// SOLUTION 1: Using Mutex
func RaceConditionMutexSolution() {
	fmt.Println("\n--- SOLUTION 1: Using Mutex ---")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // Protected by mutex
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Printf("Counter (should be 5): %d\n", counter)
	fmt.Println("✓ Mutex ensures correct value!")
}

// SOLUTION 2: Using channels (preferred Go way)
func RaceConditionChannelSolution() {
	fmt.Println("\n--- SOLUTION 2: Using Channels (Go way) ---")

	counterChan := make(chan int, 5) // Buffered channel
	var wg sync.WaitGroup

	// Senders: increment and send
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterChan <- 1 // Send increment
		}()
	}

	// Go wait and close channel
	go func() {
		wg.Wait()
		close(counterChan)
	}()

	// Sum all increments
	counter := 0
	for increment := range counterChan {
		counter += increment
	}

	fmt.Printf("Counter (should be 5): %d\n", counter)
	fmt.Println("✓ Channels isolate data access!")
}

func RaceConditionDemo() {
	fmt.Println("\n========== RACE CONDITIONS ==========")

	RaceConditionProblemDemo()
	RaceConditionMutexSolution()
	RaceConditionChannelSolution()

	fmt.Println("\n✓ Race conditions and solutions demonstrated")
}

// ============================================================================
// 5. GOROUTINE LEAKS - A Common Problem
// ============================================================================

/*
GOROUTINE LEAKS:
- Goroutines that never terminate
- Caused by blocked channels or waiting goroutines
- Lead to resource exhaustion
- Hard to debug

PREVENTION:
1. Always provide exit condition
2. Use context for cancellation
3. Ensure channels are closed
4. Avoid deadlocks
*/

// LEAK: Goroutine never terminates
func leakyFunction(ch chan int) {
	value := <-ch // Waits forever if no send
	fmt.Println(value)
}

// FIXED: Goroutine with timeout
func fixedFunction(ch chan int) {
	select {
	case value := <-ch:
		fmt.Println(value)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: goroutine exiting")
	}
}

func GoroutineLeaksDemo() {
	fmt.Println("\n========== GOROUTINE LEAKS ==========")

	fmt.Println("--- Example: Fixed with Timeout ---")
	ch := make(chan int)

	go fixedFunction(ch)

	time.Sleep(2 * time.Second)

	fmt.Println("Main: finished (goroutine handled timeout)")

	fmt.Println("\n✓ Goroutine leak prevention demonstrated")
}

// ============================================================================
// 6. COMMON CONCURRENCY PATTERNS
// ============================================================================

// Pattern 1: Producer-Consumer
func producerConsumerDemo() {
	fmt.Println("\n--- Pattern 1: Producer-Consumer ---")

	numbers := make(chan int)

	// Producer: sends numbers
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Consumer: receives and processes
	for num := range numbers {
		fmt.Printf("Processing: %d\n", num)
	}
}

// Pattern 2: Fan-out / Fan-in
func fanOutFanInDemo() {
	fmt.Println("\n--- Pattern 2: Fan-out / Fan-in ---")

	// Fan-out: distribute work
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Launch 3 workers (fan-out)
	for w := 1; w <= 3; w++ {
		go func() {
			for job := range jobs {
				results <- job * 2
			}
		}()
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results (fan-in)
	fmt.Print("Results: ")
	for range 5 {
		fmt.Printf("%d ", <-results)
	}
	fmt.Println()
}

// Pattern 3: Worker pool
type WorkerPool struct {
	jobs    chan int
	results chan int
	workers int
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan int),
		results: make(chan int),
		workers: numWorkers,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go func(id int) {
			for job := range wp.jobs {
				wp.results <- job * job // Square the number
			}
		}(i)
	}
}

func (wp *WorkerPool) Submit(job int) {
	wp.jobs <- job
}

func (wp *WorkerPool) Shutdown() {
	close(wp.jobs)
}

func (wp *WorkerPool) GetResults(count int) []int {
	results := make([]int, count)
	for i := 0; i < count; i++ {
		results[i] = <-wp.results
	}
	return results
}

func workerPoolDemo() {
	fmt.Println("\n--- Pattern 3: Worker Pool ---")

	pool := NewWorkerPool(3)
	pool.Start()

	// Submit jobs
	jobs := []int{1, 2, 3, 4, 5}
	for _, job := range jobs {
		pool.Submit(job)
	}
	pool.Shutdown()

	// Get results
	results := pool.GetResults(len(jobs))
	fmt.Printf("Results: %v\n", results)
}

func PatternsDemo() {
	fmt.Println("\n========== CONCURRENCY PATTERNS ==========")

	producerConsumerDemo()
	fanOutFanInDemo()
	workerPoolDemo()

	fmt.Println("\n✓ Concurrency patterns demonstrated")
}

// ============================================================================
// BEST PRACTICES
// ============================================================================

func CONCURRENCY_BEST_PRACTICES() {
	s := `
CONCURRENCY BEST PRACTICES:

1. GOROUTINES:
  ✓ Use 'go' keyword to launch concurrent execution
  ✓ Always ensure goroutines complete before main exits
  ✓ Use WaitGroup to synchronize
  ✓ Avoid goroutine leaks with timeouts or channels
  ✗ Don't create unbounded goroutines

2. CHANNELS:
  ✓ Use for communication between goroutines
  ✓ Close channels to signal completion
  ✓ Use for-range to receive until closed
  ✓ Use select for multiple channel operations
  ✗ Don't send on closed channels
  ✗ Don't close channels from receiver side

3. SHARED DATA:
  ✓ Prefer channels over shared memory
  ✓ Use sync.Mutex to protect critical sections
  ✓ Keep critical sections small
  ✓ Use sync.RWMutex for read-heavy workloads
  ✗ Don't share memory between goroutines

4. SYNCHRONIZATION:
  ✓ Use WaitGroup for waiting
  ✓ Use context for cancellation
  ✓ Use sync.Once for one-time initialization
  ✓ Use sync.Pool for object reuse

5. PATTERNS:
  ✓ Producer-Consumer for data pipelines
  ✓ Fan-out/Fan-in for parallel processing
  ✓ Worker pools for rate limiting
  ✓ Select for multiplexing
  ✗ Don't over-complicate with too many patterns

KEY RULE: "Don't communicate by sharing memory; share memory by communicating"
`
	fmt.Println(s)
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunConcurrencyExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    CONCURRENCY FUNDAMENTALS - COMPLETE GUIDE           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	GoroutinesBasicsDemo()
	ChannelsBasicsDemo()
	WaitGroupDemo()
	RaceConditionDemo()
	GoroutineLeaksDemo()
	PatternsDemo()
	CONCURRENCY_BEST_PRACTICES()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║  CONCURRENCY MASTERY - BUILD POWERFUL GO PROGRAMS!     ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
