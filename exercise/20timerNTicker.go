/* Exercise Twenty
20. Demonstrate the use of `time.Timer` and `time.Ticker`.

KEY CONCEPTS:
- time.Timer: A one-time event that fires after a specified duration
- time.Ticker: Repeating event that fires at regular intervals
- Both use channels to communicate when the time is reached
- Must call Stop() to prevent goroutine leaks
*/

package exercise

import (
	"fmt"
	"time"
)

// ============================================================================
// TIMER DEMONSTRATION
// ============================================================================
// Timer is used for ONE-TIME delayed actions
// It sends a single event on its channel after the duration expires

func DemoTimer() {
	fmt.Println("\n========== TIMER DEMO ==========")
	fmt.Println("Timer fires ONCE after a specified duration")

	// Create a timer that will fire after 3 seconds
	// Timer has a channel that will receive the time.Time when it fires
	timer := time.NewTimer(3 * time.Second)

	fmt.Println("Timer started. Waiting 3 seconds...")
	// This line blocks until the timer fires
	// The <-timer.C reads from the timer's channel
	fireTime := <-timer.C
	fmt.Printf("Timer fired at: %v\n", fireTime)

	fmt.Println("\n--- Timer with Early Stop ---")
	// Sometimes we want to cancel a timer before it fires
	timer2 := time.NewTimer(5 * time.Second)

	// Stop the timer before it fires (after 1 second)
	time.Sleep(1 * time.Second)
	if timer2.Stop() {
		fmt.Println("Timer stopped successfully before it fired!")
	}

	// If timer is stopped, the channel won't send anything
	// So we need to drain it to avoid goroutine leak
	select {
	case <-timer2.C:
		fmt.Println("Timer fired")
	default:
		fmt.Println("Timer was stopped, no signal on channel")
	}
}

// ============================================================================
// TICKER DEMONSTRATION
// ============================================================================
// Ticker fires repeatedly at a fixed interval
// Useful for polling, checking status, periodic tasks, heartbeats, etc.

func DemoTicker() {
	fmt.Println("\n========== TICKER DEMO ==========")
	fmt.Println("Ticker fires REPEATEDLY at fixed intervals")

	// Create a ticker that fires every 1 second
	// Ticker sends the current time on its channel each interval
	ticker := time.NewTicker(1 * time.Second)

	// IMPORTANT: Always defer Stop() to clean up the goroutine
	defer ticker.Stop()

	fmt.Println("Ticker started. Firing every 1 second...")

	// We'll collect 5 tick events then stop
	tickCount := 0
	for tick := range ticker.C {
		tickCount++
		fmt.Printf("Tick #%d at: %v\n", tickCount, tick.Format("15:04:05"))

		// Exit after 5 ticks
		if tickCount == 5 {
			fmt.Println("\nStopping ticker after 5 ticks")
			ticker.Stop()
			break
		}
	}
}

// ============================================================================
// TIMER WITH SELECT (Async Pattern)
// ============================================================================
// Using select allows non-blocking interaction with timers
// Very useful in real applications where you need to handle multiple events

func DemoTimerWithSelect() {
	fmt.Println("\n========== TIMER WITH SELECT ==========")
	fmt.Println("Non-blocking timer using select statement")

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	// Select allows us to wait on multiple channels
	// This is non-blocking - we can do other work too
	fmt.Println("Timer started, doing other work meanwhile...")

	for {
		select {
		case <-timer.C:
			fmt.Println("Timer fired! Exiting.")
			return
		default:
			fmt.Println("  → Doing other work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ============================================================================
// TICKER WITH SELECT (Async Pattern)
// ============================================================================
// This is the preferred way to use tickers in production code
// Allows you to handle multiple events and graceful shutdown

func DemoTickerWithSelect() {
	fmt.Println("\n========== TICKER WITH SELECT ==========")
	fmt.Println("Non-blocking ticker using select statement")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Create a channel to signal when we're done
	done := make(chan bool)

	// Run ticker in a goroutine to demonstrate concurrent handling
	go func() {
		tickCount := 0
		for {
			select {
			case <-done:
				fmt.Println("Ticker goroutine shutting down gracefully")
				return

			case tick := <-ticker.C:
				tickCount++
				fmt.Printf("  Tick #%d at: %v\n", tickCount, tick.Format("15:04:05.000"))

				// Stop after 5 ticks
				if tickCount >= 5 {
					done <- true
				}
			}
		}
	}()

	// Wait for the goroutine to finish
	<-done
	fmt.Println("Ticker demo complete")
}

// ============================================================================
// REAL-WORLD EXAMPLE: Health Check Monitor
// ============================================================================
// A practical example showing how to use ticker for monitoring

type HealthStatus struct {
	IsHealthy bool
	Message   string
	CheckTime time.Time
}

func DemoHealthCheckMonitor() {
	fmt.Println("\n========== HEALTH CHECK MONITOR ==========")
	fmt.Println("Periodic health check every 1 second")

	// Create a ticker for health checks every 1 second
	healthTicker := time.NewTicker(1 * time.Second)
	defer healthTicker.Stop()

	// Create a stop channel
	stop := make(chan bool)

	checkCount := 0

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Health monitor stopped")
				return

			case tick := <-healthTicker.C:
				checkCount++
				// Simulate health check logic
				status := performHealthCheck(tick)
				fmt.Printf("Check #%d - [%v] %s at %v\n",
					checkCount,
					status.IsHealthy,
					status.Message,
					status.CheckTime.Format("15:04:05"))

				// Stop after 4 checks
				if checkCount >= 4 {
					stop <- true
				}
			}
		}
	}()

	// Wait for monitor to complete
	<-stop
}

// Simulates a health check (e.g., checking if a service is up)
func performHealthCheck(t time.Time) HealthStatus {
	// In real world, this would check database, API endpoints, etc.
	healthy := time.Now().Second()%2 == 0 // Alternates healthy/unhealthy

	var msg string
	if healthy {
		msg = "✓ Service is healthy"
	} else {
		msg = "✗ Service is degraded"
	}

	return HealthStatus{
		IsHealthy: healthy,
		Message:   msg,
		CheckTime: t,
	}
}

// ============================================================================
// TIMER RESET EXAMPLE
// ============================================================================
// Shows how to reset a timer to fire after a new duration

func DemoTimerReset() {
	fmt.Println("\n========== TIMER RESET ==========")
	fmt.Println("Resetting a timer to extend/change its duration")

	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	fmt.Println("First timer set for 2 seconds")
	time.Sleep(1 * time.Second)
	fmt.Println("After 1 second, resetting timer for 2 more seconds...")

	// Reset extends the timer
	timer.Reset(2 * time.Second)

	fmt.Println("Total wait will be ~3 seconds now")
	start := time.Now()
	<-timer.C
	fmt.Printf("Timer fired after %.1f seconds from reset\n", time.Since(start).Seconds())
}

// ============================================================================
// KEY DIFFERENCES: Timer vs Ticker
// ============================================================================
/*
TIMER:
  - Fires ONCE after duration
  - Use case: Timeouts, delays, one-time events
  - Example: Timeout for API call, rate limiting, session expiry
  - Create: time.NewTimer(duration)
  - Access: <-timer.C (blocks until fired)
  - Stop: timer.Stop() (returns true if successful, false if already fired)
  - Resource: Always call Stop() to prevent goroutine leak

TICKER:
  - Fires REPEATEDLY at intervals
  - Use case: Polling, monitoring, heartbeats, periodic tasks
  - Example: Health checks, metrics collection, cache refresh
  - Create: time.NewTicker(duration)
  - Access: <-ticker.C (blocks until next tick)
  - Stop: ticker.Stop() (no return value)
  - Resource: Always call Stop() to prevent goroutine leak

COMMON MISTAKES:
  1. Forgetting to Stop() - causes goroutine leaks
  2. Not handling the fired timer/ticker properly
  3. Using timer in long loops without Reset
  4. Using blocking channels without select for multiple operations
*/

// ============================================================================
// MAIN EXECUTION FUNCTION
// ============================================================================
func RunTimerAndTickerExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║        TIMER AND TICKER COMPREHENSIVE EXAMPLES             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// Run all examples in sequence
	DemoTimer()
	DemoTicker()
	DemoTimerWithSelect()
	DemoTickerWithSelect()
	DemoHealthCheckMonitor()
	DemoTimerReset()

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    ALL EXAMPLES COMPLETE                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

// Legacy function for compatibility
func TimerAndTicker() {
	RunTimerAndTickerExamples()
}
