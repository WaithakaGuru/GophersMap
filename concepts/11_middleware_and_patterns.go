/*
MIDDLEWARE AND DESIGN PATTERNS IN GO
============================================================================

Middleware is a powerful pattern for building layered, maintainable applications.
This file covers:
- Middleware concepts and patterns
- Common middleware types (logging, auth, CORS, rate limiting)
- Middleware chains and composition
- Error handling in middleware
- Real-world middleware examples

Middleware is "code that runs before/after a handler"
Think of it as a request filter or interceptor.
*/

package concepts

import (
	"fmt"
	"net/http"
	"time"
)

// ============================================================================
// 1. MIDDLEWARE FUNDAMENTALS
// ============================================================================

/*
MIDDLEWARE PATTERN:
- Signature: func(http.Handler) http.Handler
- Wraps a handler to add functionality before/after execution
- Returns a new handler that can be wrapped by other middleware

Flow without middleware:
Browser Request → Handler → Response

Flow with middleware:
Browser Request → MW1 → MW2 → MW3 → Handler → Response

What middleware can do:
- Log requests and responses
- Add/modify headers
- Authenticate and authorize users
- Compress responses
- Add request ID for tracing
- Implement rate limiting
- Handle CORS
- Recover from panics
*/

// BasicMiddleware is a simple middleware example
func BasicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Code runs BEFORE handler
		fmt.Printf("Before: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Code runs AFTER handler
		fmt.Println("After: Response sent")
	})
}

// MiddlewareConceptsDemo explains middleware fundamentals
func MiddlewareConceptsDemo() {
	fmt.Println("\n========== MIDDLEWARE CONCEPTS ==========")

	fmt.Println("\n1. MIDDLEWARE PATTERN:")
	fmt.Println("   - Takes http.Handler as input")
	fmt.Println("   - Returns http.Handler as output")
	fmt.Println("   - Signature: func(http.Handler) http.Handler")

	fmt.Println("\n2. MIDDLEWARE FLOW:")
	fmt.Println("   Request → Middleware → Handler → Middleware → Response")
	fmt.Println("   Middleware runs BEFORE and AFTER handler")

	fmt.Println("\n3. USES:")
	fmt.Println("   - Logging requests/responses")
	fmt.Println("   - Authentication and authorization")
	fmt.Println("   - CORS handling")
	fmt.Println("   - Rate limiting")
	fmt.Println("   - Request/response transformation")
	fmt.Println("   - Error recovery")
	fmt.Println("   - Request tracing")
	fmt.Println("   - Compression")

	fmt.Println("\n4. EXAMPLE:")
	exampleHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}
	wrappedHandler := BasicMiddleware(exampleHandler)

	fmt.Println("   - Create handler: exampleHandler")
	fmt.Println("   - Wrap with middleware: BasicMiddleware(exampleHandler)")
	fmt.Println("   - Middleware logs before/after handler execution")
	_ = wrappedHandler

	fmt.Println("\n✓ Middleware concepts demonstrated")
}

// ============================================================================
// 2. COMMON MIDDLEWARE TYPES
// ============================================================================

/*
LOGGING MIDDLEWARE:
- Records request details (method, path, timestamp)
- Measures request duration
- Records response status code
- Useful for debugging and monitoring

AUTHENTICATION MIDDLEWARE:
- Checks for valid credentials/token
- Rejects unauthorized requests (401)
- Adds user info to request context
- Protects sensitive endpoints

CORS MIDDLEWARE:
- Allows cross-origin requests
- Sets Access-Control headers
- Handles preflight (OPTIONS) requests
- Essential for frontend-backend communication

RATE LIMITING MIDDLEWARE:
- Tracks requests per IP/user
- Returns 429 Too Many Requests when limit exceeded
- Prevents abuse and DDoS attacks
- Configurable time windows and limits
*/

// LoggingMiddleware logs request and response information
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record start time
		start := time.Now()

		// Log request
		fmt.Printf("[%s] %s %s\n",
			time.Now().Format("15:04:05"),
			r.Method,
			r.URL.Path,
		)

		// Call next handler
		next.ServeHTTP(w, r)

		// Log response time
		duration := time.Since(start)
		fmt.Printf("   Completed in %v\n", duration)
	})
}

// CORSMiddleware handles cross-origin requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

// AuthenticationMiddleware checks for valid authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// In real app, verify token here
		fmt.Printf("Authenticating with token: %s\n", token)

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

// HeaderMiddleware adds headers to response
func HeaderMiddleware(key, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}

// CommonMiddlewareDemo demonstrates common middleware
func CommonMiddlewareDemo() {
	fmt.Println("\n========== COMMON MIDDLEWARE TYPES ==========")

	fmt.Println("\n1. LOGGING MIDDLEWARE:")
	fmt.Println("   - Records request method, path, timestamp")
	fmt.Println("   - Measures request duration")
	fmt.Println("   - Useful for debugging and monitoring")

	fmt.Println("\n2. CORS MIDDLEWARE:")
	fmt.Println("   - Allows cross-origin requests")
	fmt.Println("   - Sets Access-Control-Allow-* headers")
	fmt.Println("   - Handles OPTIONS preflight requests")

	fmt.Println("\n3. AUTHENTICATION MIDDLEWARE:")
	fmt.Println("   - Checks Authorization header")
	fmt.Println("   - Verifies tokens/credentials")
	fmt.Println("   - Returns 401 Unauthorized if missing/invalid")

	fmt.Println("\n4. HEADER MIDDLEWARE:")
	fmt.Println("   - Adds custom headers to responses")
	fmt.Println("   - Examples: API-Version, Content-Type, etc.")
	fmt.Println("   - Can be created with higher-order functions")

	fmt.Println("\n✓ Common middleware types demonstrated")
}

// ============================================================================
// 3. MIDDLEWARE CHAINS
// ============================================================================

/*
MIDDLEWARE CHAINS:
- Multiple middleware wrapped together
- Executes in order of wrapping (reversed for execution)
- Each middleware can modify request/response or stop execution

Order matters!
- Inner middleware executes before outer middleware
- Wrap helpers apply middleware in reverse order

Example:
Chain(mux,
    Middleware1,  // Outermost (runs first)
    Middleware2,
    Middleware3,  // Innermost (runs closest to handler)
)

Execution order:
1. Middleware1 starts
2. Middleware2 starts
3. Middleware3 starts
4. Handler
5. Middleware3 ends
6. Middleware2 ends
7. Middleware1 ends
*/

// Chain applies multiple middleware to a handler
func Chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	// Apply middleware in reverse order so they execute in declaration order
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

// MiddlewareChainDemo shows how to chain middleware
func MiddlewareChainDemo() {
	fmt.Println("\n========== MIDDLEWARE CHAINS ==========")

	fmt.Println("\n1. CHAINING MIDDLEWARE:")
	fmt.Println("   - Apply multiple middleware to single handler")
	fmt.Println("   - Order matters: inner executes closer to handler")
	fmt.Println("   - Each middleware can skip or modify request")

	fmt.Println("\n2. CHAIN HELPER FUNCTION:")
	fmt.Println("   func Chain(h Handler, mw ...func(Handler) Handler) Handler")
	fmt.Println("   - Takes handler and middleware functions")
	fmt.Println("   - Applies in reverse order for correct execution")
	fmt.Println("   - Returns wrapped handler")

	fmt.Println("\n3. EXECUTION ORDER:")
	fmt.Println("   Chain(handler, MW1, MW2, MW3)")
	fmt.Println("   Execution: MW1 → MW2 → MW3 → handler → MW3 → MW2 → MW1")

	fmt.Println("\n4. EXAMPLE:")
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Response")
	})

	// Create middleware chain
	handler := Chain(mux,
		LoggingMiddleware,
		CORSMiddleware,
		HeaderMiddleware("X-API-Version", "1.0"),
	)

	fmt.Println("   - Create mux with handlers")
	fmt.Println("   - Chain middleware together")
	fmt.Println("   - Pass to http.ListenAndServe")
	_ = handler

	fmt.Println("\n✓ Middleware chains demonstrated")
}

// ============================================================================
// 4. ERROR HANDLING MIDDLEWARE
// ============================================================================

/*
PANIC RECOVERY:
- Middleware can recover from panics
- Prevents entire server from crashing
- Returns proper HTTP error response
- Logs the panic for debugging

ERROR LOGGING:
- Middleware wraps ResponseWriter to capture status code
- Can log errors (4xx, 5xx responses) separately
- Useful for error tracking and monitoring
*/

// RecoveryMiddleware recovers from panics
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"error": "Internal server error"}`)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// ResponseWriterWrapper wraps http.ResponseWriter to capture status code
type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
	Written    bool
}

// WriteHeader captures the status code
func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	if !w.Written {
		w.StatusCode = statusCode
		w.Written = true
		w.ResponseWriter.WriteHeader(statusCode)
	}
}

// ErrorLoggingMiddleware logs error responses
func ErrorLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &ResponseWriterWrapper{ResponseWriter: w, StatusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		// Log errors (4xx, 5xx)
		if wrapped.StatusCode >= 400 {
			fmt.Printf("ERROR: %s %s - Status %d\n",
				r.Method, r.URL.Path, wrapped.StatusCode)
		}
	})
}

// ErrorHandlingMiddlewareDemo shows error handling
func ErrorHandlingMiddlewareDemo() {
	fmt.Println("\n========== ERROR HANDLING MIDDLEWARE ==========")

	fmt.Println("\n1. PANIC RECOVERY:")
	fmt.Println("   - Recovers from handler panics")
	fmt.Println("   - Returns 500 Internal Server Error")
	fmt.Println("   - Prevents server crash")
	fmt.Println("   - Should be outermost middleware")

	fmt.Println("\n2. ERROR LOGGING:")
	fmt.Println("   - Wraps ResponseWriter to capture status")
	fmt.Println("   - Logs 4xx and 5xx responses separately")
	fmt.Println("   - Useful for monitoring and debugging")

	fmt.Println("\n3. PROPER ORDERING:")
	fmt.Println("   RecoveryMiddleware")
	fmt.Println("   └─ ErrorLoggingMiddleware")
	fmt.Println("      └─ LoggingMiddleware")
	fmt.Println("         └─ Handler")

	fmt.Println("\n✓ Error handling middleware demonstrated")
}

// ============================================================================
// 5. REAL-WORLD EXAMPLE
// ============================================================================

// RealWorldMiddlewareDemo shows production middleware setup
func RealWorldMiddlewareDemo() {
	fmt.Println("\n========== PRODUCTION MIDDLEWARE STACK ==========")

	fmt.Println("\n1. TYPICAL MIDDLEWARE ORDER (innermost to outermost):")
	fmt.Println("   1. Recovery         - Catch panics")
	fmt.Println("   2. Error Logging    - Log errors separately")
	fmt.Println("   3. Request Logging  - Log all requests")
	fmt.Println("   4. CORS             - Handle cross-origin")
	fmt.Println("   5. Headers          - Add custom headers")
	fmt.Println("   6. Authentication   - Verify tokens (optional)")

	fmt.Println("\n2. SETUP EXAMPLE:")
	fmt.Println("   mux := http.NewServeMux()")
	fmt.Println("   mux.HandleFunc(\"/health\", healthHandler)")
	fmt.Println("   mux.HandleFunc(\"/api/data\", dataHandler)")
	fmt.Println("")
	fmt.Println("   handler := Chain(mux,")
	fmt.Println("       RecoveryMiddleware,")
	fmt.Println("       ErrorLoggingMiddleware,")
	fmt.Println("       LoggingMiddleware,")
	fmt.Println("       CORSMiddleware,")
	fmt.Println("   )")
	fmt.Println("")
	fmt.Println("   http.ListenAndServe(\":8080\", handler)")

	fmt.Println("\n3. SELECTIVE MIDDLEWARE:")
	fmt.Println("   - Apply different middleware to different routes")
	fmt.Println("   - Public routes: no auth middleware")
	fmt.Println("   - Protected routes: with auth middleware")
	fmt.Println("   - All routes: recovery and logging")

	fmt.Println("\n✓ Production middleware stack demonstrated")
}

// ============================================================================
// 6. MAIN EXECUTION
// ============================================================================

// RunMiddlewareExamples executes all middleware demonstrations
func RunMiddlewareExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   MIDDLEWARE AND PATTERNS - COMPREHENSIVE GUIDE         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	MiddlewareConceptsDemo()
	CommonMiddlewareDemo()
	MiddlewareChainDemo()
	ErrorHandlingMiddlewareDemo()
	RealWorldMiddlewareDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL MIDDLEWARE EXAMPLES COMPLETE                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
	fmt.Println("✓ Middleware wraps handlers: func(Handler) -> Handler")
	fmt.Println("✓ Run code before and after handler execution")
	fmt.Println("✓ Chain multiple middleware for layered functionality")
	fmt.Println("✓ Use context to pass data between middleware and handlers")
	fmt.Println("✓ Apply middleware in correct order (innermost runs closest to handler)")
	fmt.Println("✓ Common middleware: logging, auth, CORS, rate limiting")
	fmt.Println("✓ Recover from panics to prevent server crashes")
	fmt.Println("✓ Log errors separately for debugging")
	fmt.Println("✓ Use Chain() helper for clean middleware composition")
	fmt.Println("✓ Middleware is the key to clean, layered architecture\n")
}
