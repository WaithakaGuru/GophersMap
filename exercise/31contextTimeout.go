/* Exercise Thirty One
31. Use `context.WithTimeout` to cancel a long-running operation.

KEY CONCEPTS:
- Context: Carries deadlines, cancellation signals, and values across API boundaries
- context.WithTimeout: Creates a context that cancels after a duration
- context.Background: Base context (use for main functions)
- Goroutines should ALWAYS check ctx.Done() to respect cancellation
- Use select with <-ctx.Done() for non-blocking cancellation checks
*/

package exercise

import (
	"context"
	"fmt"
	"time"
)

// ============================================================================
// BASIC EXAMPLE: Long Running Operation without Context
// ============================================================================
// This example shows what happens when we DON'T respect context

func basicLongOperation() {
	fmt.Println("\n========== BASIC LONG OPERATION (No Context) ==========")
	fmt.Println("Running 10 iterations, each takes 1 second...")

	count := 0
	for range 10 {
		time.Sleep(time.Second)
		count++
		fmt.Printf("  → Progress: %d/10 completed\n", count)
	}
	fmt.Printf("✓ Operation completed after %d seconds\n", count)
}

// ============================================================================
// IMPROVED EXAMPLE: Using Context to Respect Cancellation
// ============================================================================
// This is the CORRECT way - the function respects context cancellation

func contextAwareLongOperation(ctx context.Context, operationName string) error {
	fmt.Printf("\n========== %s (Context-Aware) ==========\n", operationName)
	fmt.Println("Running operation with context timeout...")

	count := 0
	// Run the operation in a loop, checking context at each iteration
	for i := 1; i <= 10; i++ {
		// CRITICAL: Always check if context is cancelled
		// This allows graceful shutdown when timeout occurs
		select {
		case <-ctx.Done():
			// Context cancelled or deadline exceeded
			err := ctx.Err()
			fmt.Printf("\n✗ Operation cancelled after %d iterations: %v\n", count, err)
			return err

		default:
			// Context not cancelled, continue processing
			time.Sleep(time.Second)
			count++
			fmt.Printf("  → Iteration %d/%d completed\n", i, 10)
		}
	}

	fmt.Printf("✓ Operation completed successfully after %d seconds\n", count)
	return nil
}

// ============================================================================
// CONTEXT WITH TIMEOUT
// ============================================================================
// Creates a context that automatically cancels after a specified duration

func DemoWithTimeout() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      CONTEXT WITH TIMEOUT DEMONSTRATION             ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	// background is the root context - typically used in main()
	ctx := context.Background()

	// Create a context that will timeout after 5 seconds
	// IMPORTANT: Always defer cancel() to release resources
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	fmt.Println("\nContext created with 5 second timeout")
	fmt.Println("Operation will run until timeout or completion")

	// Pass the timeout context to the operation
	// The operation will stop when timeout occurs OR completes
	err := contextAwareLongOperation(timeoutCtx, "OPERATION WITH 5 SECOND TIMEOUT")

	if err != nil {
		fmt.Printf("Error returned: %v\n", err)
	}
}

// ============================================================================
// MULTIPLE OPERATIONS WITH SHARED TIMEOUT
// ============================================================================
// Shows how multiple goroutines can share the same timeout context

func multipleOperationsWithTimeout() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      MULTIPLE OPERATIONS WITH SHARED TIMEOUT        ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	fmt.Println("\nStarting 2 operations with shared 6 second timeout")

	// Channel to collect results
	done := make(chan string, 2)

	// First operation
	go func() {
		fmt.Println("Operation 1: Starting...")
		err := contextAwareLongOperation(timeoutCtx, "OPERATION 1")
		if err != nil {
			done <- fmt.Sprintf("Op1 failed: %v", err)
		} else {
			done <- "Op1 succeeded"
		}
	}()

	// Second operation
	go func() {
		fmt.Println("\nOperation 2: Starting...")
		err := contextAwareLongOperation(timeoutCtx, "OPERATION 2")
		if err != nil {
			done <- fmt.Sprintf("Op2 failed: %v", err)
		} else {
			done <- "Op2 succeeded"
		}
	}()

	// Collect results from both operations
	for i := 0; i < 2; i++ {
		result := <-done
		fmt.Printf("Result: %s\n", result)
	}
}

// ============================================================================
// CONTEXT WITH DEADLINE (Alternative to Timeout)
// ============================================================================
// WithDeadline sets an absolute time when context expires
// WithTimeout is actually just WithDeadline(time.Now() + duration)

func DemoWithDeadline() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      CONTEXT WITH DEADLINE DEMONSTRATION            ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	ctx := context.Background()

	// Set deadline to 4 seconds from now
	deadline := time.Now().Add(4 * time.Second)
	deadlineCtx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	fmt.Printf("\nDeadline set to: %v\n", deadline.Format("15:04:05"))
	fmt.Println("Operation will run until deadline is reached")

	err := contextAwareLongOperation(deadlineCtx, "OPERATION WITH DEADLINE")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// ============================================================================
// CONTEXT WITH VALUES
// ============================================================================
// Contexts can also carry values across function boundaries

func DemoContextWithValues() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      CONTEXT WITH VALUES                             ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	ctx := context.Background()

	// Define custom key type (prevents collisions with other packages)
	type userIDKey string
	const userID userIDKey = "userID"

	// Add a value to the context
	valueCtx := context.WithValue(ctx, userID, "user123")

	// Retrieve the value
	if id, ok := valueCtx.Value(userID).(string); ok {
		fmt.Printf("\nUser ID from context: %s\n", id)
	}

	// This is useful for passing request IDs, authentication info, etc.
	// across functions without explicit parameters
}

// ============================================================================
// CONTEXT ERROR TYPES
// ============================================================================
// Understanding different context errors

func DemoContextErrors() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      CONTEXT ERROR TYPES                             ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	ctx := context.Background()

	// Create a context that times out immediately
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()

	// Wait a bit to let it timeout
	time.Sleep(100 * time.Millisecond)

	// Check the error type
	err := timeoutCtx.Err()
	fmt.Printf("\nContext error: %v\n", err)
	fmt.Printf("Error type: %T\n", err)

	// In Go, you can use context.DeadlineExceeded to check
	if err == context.DeadlineExceeded {
		fmt.Println("✓ This is a deadline exceeded error")
	}
}

// ============================================================================
// REAL-WORLD EXAMPLE: HTTP Request with Timeout
// ============================================================================
// Simulating a realistic scenario

func simulatedHTTPRequest(ctx context.Context, url string) (string, error) {
	fmt.Printf("\nMaking HTTP request to %s...", url)

	// Simulate network delay
	select {
	case <-time.After(3 * time.Second):
		// Request would complete normally
		return "Response from " + url, nil

	case <-ctx.Done():
		// Context cancelled before request completed
		return "", ctx.Err()
	}
}

func DemoRealisticHTTPTimeout() {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║      REALISTIC HTTP REQUEST WITH TIMEOUT             ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	ctx := context.Background()

	// Set a 2 second timeout for the request
	// Since the "request" takes 3 seconds, it will timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	fmt.Println("\nAttempting request with 2 second timeout...")
	response, err := simulatedHTTPRequest(timeoutCtx, "https://api.example.com")

	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request succeeded: %s\n", response)
	}
}

// ============================================================================
// BEST PRACTICES
// ============================================================================
/*
1. ALWAYS DEFER CANCEL:
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()

2. CHECK CONTEXT IN LOOPS:
   select {
   case <-ctx.Done():
       return ctx.Err()
   default:
       // do work
   }

3. PASS CONTEXT AS FIRST PARAMETER:
   func DoSomething(ctx context.Context, other string) error { }

4. USE CONTEXT.BACKGROUND() AT ENTRY POINTS:
   func main() {
       ctx := context.Background()
   }

5. NEVER STORE CONTEXT IN STRUCT:
   ❌ type MyType struct { ctx context.Context }
   ✓ Pass context as parameter instead

6. USE DIFFERENT CONTEXT TYPES FOR DIFFERENT NEEDS:
   - WithTimeout: For operations with time limits
   - WithDeadline: For absolute time limits
   - WithCancel: For manual cancellation
   - WithValue: For passing request-scoped values
*/

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunContextExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    CONTEXT WITH TIMEOUT COMPREHENSIVE EXAMPLES         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	// Show the problem: operation without context respect
	basicLongOperation()

	// Show the solution: with timeout
	DemoWithTimeout()

	// Multiple operations sharing timeout
	multipleOperationsWithTimeout()

	// Alternative: deadline instead of timeout
	DemoWithDeadline()

	// Passing values through context
	DemoContextWithValues()

	// Understanding error types
	DemoContextErrors()

	// Realistic scenario
	DemoRealisticHTTPTimeout()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                ALL EXAMPLES COMPLETE                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}

// Legacy function for compatibility
func TimeOut(ctx context.Context) {
	RunContextExamples()
}

// Original slow operation with fixes for context awareness
func slowOperation(ctx context.Context) {
	count := 0
	fmt.Println("\n--- Started long running operation ---")
	for i := 1; i <= 10; i++ {
		// FIXED: Check context for cancellation
		select {
		case <-ctx.Done():
			fmt.Printf("Operation interrupted after %d iterations: %v\n", count, ctx.Err())
			return
		default:
			time.Sleep(time.Second)
			count++
			fmt.Printf("--Current progress: COUNT -> %d/9\n", count)
		}
	}
	fmt.Println("---END: Finished long running operation---")
}
