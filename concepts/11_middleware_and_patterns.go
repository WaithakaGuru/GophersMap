package concepts

import (
	"fmt"
	_ "net/http"
)

// ============================================================================
// MIDDLEWARE AND DESIGN PATTERNS
// ============================================================================
// This file covers:
// - Middleware concepts and patterns
// - Common middleware types
// - Middleware chains
// - Composition patterns
// - Real-world middleware examples
// ============================================================================

// MiddlewareConceptsDemo shows middleware fundamentals
func MiddlewareConceptsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       MIDDLEWARE CONCEPTS DEMO                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. WHAT IS MIDDLEWARE:")
	fmt.Println(`
// Middleware = Code that runs before/after handler
// Think of it as a "request filter" or "interceptor"

// Without middleware:
Browser Request → Handler → Response

// With middleware:
Browser Request → MW1 → MW2 → MW3 → Handler → Response

// Middleware can:
// - Log requests
// - Add headers to response
// - Authenticate/authorize users
// - Compress responses
// - Add request ID for tracing
// - Rate limiting
// - CORS handling
// - Error recovery
`)

	fmt.Println("\n2. MIDDLEWARE PATTERN:")
	fmt.Println(`
// Signature: func(http.Handler) http.Handler

// Simple middleware structure:
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // BEFORE handler
        fmt.Println("Request coming in")
        
        // Call next handler
        next.ServeHTTP(w, r)
        
        // AFTER handler
        fmt.Println("Response going out")
    })
}

// Usage:
mux := http.NewServeMux()
mux.HandleFunc("/api", myHandler)

// Wrap with middleware
handler := MyMiddleware(mux)

// Start server with wrapped handler
http.ListenAndServe(":8080", handler)
`)

	fmt.Println("✓ Middleware concepts demonstrated")
}

// CommonMiddlewareDemo shows typical middleware examples
func CommonMiddlewareDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       COMMON MIDDLEWARE TYPES DEMO                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. LOGGING MIDDLEWARE:")
	fmt.Println(`
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Record request details
        start := time.Now()
        
        fmt.Printf("[%s] %s %s\n", 
            time.Now().Format("2006-01-02 15:04:05"),
            r.Method,
            r.URL.Path,
        )
        
        // Wrap ResponseWriter to capture status code
        wrapped := &ResponseWriterWrapper{ResponseWriter: w}
        
        // Call next handler
        next.ServeHTTP(wrapped, r)
        
        // Log response details
        duration := time.Since(start)
        fmt.Printf("  Status: %d, Duration: %v\n", wrapped.StatusCode, duration)
    })
}

type ResponseWriterWrapper struct {
    http.ResponseWriter
    StatusCode int
    Written    bool
}

func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
    if !w.Written {
        w.StatusCode = statusCode
        w.Written = true
        w.ResponseWriter.WriteHeader(statusCode)
    }
}
`)

	fmt.Println("\n2. AUTHENTICATION MIDDLEWARE:")
	fmt.Println(`
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get token from header
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Missing authorization", http.StatusUnauthorized)
            return
        }
        
        // Verify token
        claims, err := VerifyToken(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Add claims to request context
        ctx := context.WithValue(r.Context(), "user", claims)
        r = r.WithContext(ctx)
        
        // Continue to next handler
        next.ServeHTTP(w, r)
    })
}

// Usage in handler:
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user")
    fmt.Fprintf(w, "Hello %v", user)
}
`)

	fmt.Println("\n3. CORS MIDDLEWARE:")
	fmt.Println(`
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
        
        // Continue to next handler
        next.ServeHTTP(w, r)
    })
}
`)

	fmt.Println("\n4. RATE LIMITING MIDDLEWARE:")
	fmt.Println(`
type RateLimiter struct {
    visits map[string][]time.Time
    mu     sync.RWMutex
    limit  int           // requests per window
    window time.Duration // time window
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        
        rl.mu.Lock()
        defer rl.mu.Unlock()
        
        now := time.Now()
        
        // Clean old entries outside window
        if visits, ok := rl.visits[ip]; ok {
            rl.visits[ip] = filterRecent(visits, now, rl.window)
        }
        
        // Check limit
        if len(rl.visits[ip]) >= rl.limit {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        // Add current request
        rl.visits[ip] = append(rl.visits[ip], now)
        
        // Continue
        next.ServeHTTP(w, r)
    })
}
`)

	fmt.Println("✓ Common middleware types demonstrated")
}

// MiddlewareChainDemo shows how to chain multiple middleware
func MiddlewareChainDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       MIDDLEWARE CHAINS DEMO                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. CHAINING MIDDLEWARE:")
	fmt.Println(`
// Apply multiple middleware in order

// Manual chaining:
mux := http.NewServeMux()
mux.HandleFunc("/api", myHandler)

handler := LoggingMiddleware(
    AuthenticationMiddleware(
        CORSMiddleware(mux),
    ),
)

http.ListenAndServe(":8080", handler)

// Execution order (for /api request):
// 1. LoggingMiddleware (starts)
// 2. AuthenticationMiddleware (starts)
// 3. CORSMiddleware (starts)
// 4. myHandler
// 5. CORSMiddleware (ends)
// 6. AuthenticationMiddleware (ends)
// 7. LoggingMiddleware (ends)
`)

	fmt.Println("\n2. MIDDLEWARE CHAIN HELPER:")
	fmt.Println(`
// Helper function for cleaner chaining
func Chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}

// Usage:
mux := http.NewServeMux()
mux.HandleFunc("/api", myHandler)

handler := Chain(mux,
    LoggingMiddleware,
    AuthenticationMiddleware,
    CORSMiddleware,
    RateLimitingMiddleware,
)

http.ListenAndServe(":8080", handler)
`)

	fmt.Println("\n3. ROUTE-SPECIFIC MIDDLEWARE:")
	fmt.Println(`
// Apply middleware to specific routes only

type Router struct {
    *http.ServeMux
}

func (r *Router) HandleWithMiddleware(pattern string, handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) {
    // Apply middleware to this specific handler
    h := http.Handler(handler)
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    r.Handle(pattern, h)
}

// Usage:
router := &Router{http.NewServeMux()}

// Public route (no auth middleware)
router.HandleFunc("/login", loginHandler)

// Protected route (with auth middleware)
router.HandleWithMiddleware("/api/protected",
    protectedHandler,
    AuthenticationMiddleware,
    LoggingMiddleware,
)
`)

	fmt.Println("✓ Middleware chains demonstrated")
}

// CompositionPatternsDemo shows design patterns
func CompositionPatternsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       COMPOSITION PATTERNS DEMO                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. DEPENDENCY INJECTION MIDDLEWARE:")
	fmt.Println(`
// Middleware that needs dependencies

type AppMiddleware struct {
    logger   Logger
    db       Database
    config   Config
}

func (a *AppMiddleware) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        a.logger.Info("Authenticating request")
        
        token := r.Header.Get("Authorization")
        user, err := a.db.GetUserByToken(token)
        if err != nil {
            a.logger.Error("Auth failed", err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), "user", user)
        r = r.WithContext(ctx)
        
        next.ServeHTTP(w, r)
    })
}

// Create middleware with dependencies
app := &AppMiddleware{
    logger: newLogger(),
    db:     newDatabase(),
    config: newConfig(),
}

// Use in router
mux := http.NewServeMux()
mux.HandleFunc("/api/protected", protectedHandler)
handler := app.AuthMiddleware(mux)
`)

	fmt.Println("\n2. HIGHER-ORDER FUNCTIONS:")
	fmt.Println(`
// Create middleware factories

func WithTimeout(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()
            
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func WithHeader(key, value string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set(key, value)
            next.ServeHTTP(w, r)
        })
    }
}

// Usage:
handler := Chain(mux,
    WithTimeout(30 * time.Second),
    WithHeader("X-API-Version", "1.0"),
    LoggingMiddleware,
)
`)

	fmt.Println("✓ Composition patterns demonstrated")
}

// ErrorHandlingMiddlewareDemo shows error handling
func ErrorHandlingMiddlewareDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       ERROR HANDLING MIDDLEWARE DEMO                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. PANIC RECOVERY MIDDLEWARE:")
	fmt.Println(`
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Log the panic
                fmt.Printf("Panic: %v\n", err)
                
                // Return error response
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, ` + "`{\"error\": \"Internal server error\"}`" + `)
            }
        }()
        
        // Continue execution
        next.ServeHTTP(w, r)
    })
}
`)

	fmt.Println("\n2. ERROR LOGGING MIDDLEWARE:")
	fmt.Println(`
type LoggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (w *LoggingResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}

func ErrorLoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        wrapped := &LoggingResponseWriter{ResponseWriter: w}
        
        next.ServeHTTP(wrapped, r)
        
        // Log errors (4xx, 5xx responses)
        if wrapped.statusCode >= 400 {
            fmt.Printf("ERROR: %s %s - Status %d\n", 
                r.Method, r.URL.Path, wrapped.statusCode)
        }
    })
}
`)

	fmt.Println("✓ Error handling middleware demonstrated")
}

// RealWorldMiddlewareDemo shows practical example
func RealWorldMiddlewareDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    REAL WORLD MIDDLEWARE EXAMPLE DEMO                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("PRODUCTION API MIDDLEWARE STACK:")
	fmt.Println(`
// Typical production API middleware stack:

1. RecoveryMiddleware
   └─ Catches panics, returns 500 errors

2. LoggingMiddleware
   └─ Logs all requests and responses

3. CORSMiddleware
   └─ Handles cross-origin requests

4. RateLimitingMiddleware
   └─ Prevents abuse

5. AuthenticationMiddleware
   └─ Verifies JWT tokens

6. ErrorLoggingMiddleware
   └─ Logs errors separately

// Implementation:
func setupServer() http.Handler {
    mux := http.NewServeMux()
    
    // Public routes
    mux.HandleFunc("/health", healthHandler)
    mux.HandleFunc("/login", loginHandler)
    
    // Protected routes
    mux.HandleFunc("/api/data", dataHandler)
    
    // Apply middleware stack
    handler := Chain(mux,
        RecoveryMiddleware,
        ErrorLoggingMiddleware,
        LoggingMiddleware,
        CORSMiddleware,
        RateLimitingMiddleware,
        AuthenticationMiddleware,
    )
    
    return handler
}

// Start server
func main() {
    handler := setupServer()
    http.ListenAndServe(":8080", handler)
}
`)

	fmt.Println("✓ Real world middleware example demonstrated")
}

// RunMiddlewareExamples executes all middleware demos
func RunMiddlewareExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   MIDDLEWARE AND PATTERNS - COMPREHENSIVE GUIDE         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	MiddlewareConceptsDemo()
	CommonMiddlewareDemo()
	MiddlewareChainDemo()
	CompositionPatternsDemo()
	ErrorHandlingMiddlewareDemo()
	RealWorldMiddlewareDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL MIDDLEWARE EXAMPLES COMPLETE                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Middleware wraps handlers: func(Handler) -> Handler")
	fmt.Println("✓ Run code before and after handler execution")
	fmt.Println("✓ Chain multiple middleware for layered functionality")
	fmt.Println("✓ Use context to pass data between middleware and handlers")
	fmt.Println("✓ Apply middleware in correct order (reverse of execution)")
	fmt.Println("✓ Common middleware: logging, auth, CORS, rate limiting")
	fmt.Println("✓ Recover from panics to prevent server crashes")
	fmt.Println("✓ Log errors separately for debugging")
	fmt.Println("✓ Use dependency injection for middleware with state")
	fmt.Println("✓ Middleware is the key to clean, layered architecture\n")
}
