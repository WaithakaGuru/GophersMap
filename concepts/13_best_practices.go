package concepts

import (
	"fmt"
)

// ============================================================================
// BEST PRACTICES AND ADVANCED PATTERNS
// ============================================================================
// This file covers:
// - Code organization and structure
// - Package design
// - Naming conventions
// - Documentation standards
// - Common pitfalls and solutions
// - Performance optimization tips
// ============================================================================

// CodeOrganizationDemo shows how to structure Go projects
func CodeOrganizationDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       CODE ORGANIZATION AND STRUCTURE DEMO             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. RECOMMENDED PROJECT STRUCTURE:")
	fmt.Println(`
myapp/
├── main.go                 # Entry point
├── go.mod                  # Module definition
├── go.sum                  # Dependency hashes
├── README.md               # Project documentation
├── LICENSE                 # License file
│
├── cmd/                    # Executables
│   ├── server/
│   │   └── main.go        # Server executable
│   └── cli/
│       └── main.go        # CLI executable
│
├── internal/               # Private packages (not importable)
│   ├── models/
│   │   └── user.go        # Data models
│   ├── handlers/
│   │   └── user.go        # HTTP handlers
│   ├── services/
│   │   └── user.go        # Business logic
│   ├── repositories/
│   │   └── user.go        # Database access
│   └── middleware/
│       └── auth.go        # HTTP middleware
│
├── pkg/                    # Public packages (importable by others)
│   ├── logger/
│   │   └── logger.go
│   └── config/
│       └── config.go
│
├── tests/                  # Integration/end-to-end tests
│   └── api_test.go
│
├── migrations/             # Database migrations
│   └── 001_create_users.sql
│
├── docs/                   # Documentation
│   ├── API.md
│   └── ARCHITECTURE.md
│
└── .gitignore              # Git ignore rules

// Package naming rules:
// - Lowercase only
// - Single word preferred (handlers, not handler_logic)
// - Short and descriptive
// - Avoid generic names (util, helper, common)
`)

	fmt.Println("\n2. PACKAGE ORGANIZATION PRINCIPLES:")
	fmt.Println(`
// Go philosophy: organize by FUNCTIONALITY, not LAYER

// ✗ Organize by type (don't do this):
models/
  user.go
  product.go
handlers/
  user.go
  product.go
services/
  user.go
  product.go

// ✓ Organize by feature (do this):
users/
  user.go
  handler.go
  service.go
  repository.go
products/
  product.go
  handler.go
  service.go
  repository.go

// Benefits:
// ✓ Related code is together
// ✓ Easy to find what you need
// ✓ Easier to delete features
// ✓ Clear dependencies
`)

	fmt.Println("✓ Code organization demonstrated")
}

// NamingConventionsDemo shows Go naming standards
func NamingConventionsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       NAMING CONVENTIONS DEMO                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. VARIABLE AND FUNCTION NAMING:")
	fmt.Println(`
// Variables: Use camelCase, be descriptive
var (
    maxRetries = 3
    userName = "john"           // Good
    user_name = "john"          // Bad (snake_case)
    n = "john"                  // Bad (too short)
    userNameForAuthentication = "john"  // Too long
)

// Constants: Use CamelCase, not SCREAMING_SNAKE_CASE
const (
    DefaultPort = 8080       // Good
    DEFAULT_PORT = 8080      // Bad
    TIMEOUT = 30             // Bad
    DefaultTimeout = 30      // Good
)

// Functions: Use camelCase, start with lowercase
func getUserByID(id int) (*User, error) { }  // Good
func GetUserByID(id int) (*User, error) { }  // Exported (public)
func get_user_by_id(id int) (*User, error) { }  // Bad

// Single letter names (only for loops/contexts):
for i := 0; i < 10; i++ { }       // OK
ctx := context.Background()       // OK
buf := bytes.NewBuffer(data)      // Good abbreviation

// Acronyms: Use CamelCase, not ALL_CAPS
type HTTPHandler struct { }        // Good
type HttpHandler struct { }        // Bad
type HTTP_Handler struct { }       // Bad
`)

	fmt.Println("\n2. INTERFACE NAMING:")
	fmt.Println(`
// Single method interfaces: use er suffix
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Multi-method interfaces: descriptive names
type Storage interface {
    Save(key string, value interface{}) error
    Get(key string) (interface{}, error)
    Delete(key string) error
}

// Keep interfaces small (1-3 methods is ideal)
`)

	fmt.Println("\n3. PACKAGE NAMING:")
	fmt.Println(`
// Package names should be lowercase, no underscores
package models          // Good
package model           // Good (singular preferred)
package data_models     // Bad
package DataModels      // Bad
package models_v2       // Bad

// Avoid:
// - Plurals (imported as models.models - redundant)
// - Generic names (util, helper, common)
// - Names that conflict with stdlib

// Good package names:
encoding/json           // Clear, descriptive
net/http                // Clear, descriptive
time                    # Clear, descriptive
strings                 // Clear, descriptive

// Bad package names:
util                    // Too generic
helper                  // Too generic
common                  // Too generic
`)

	fmt.Println("✓ Naming conventions demonstrated")
}

// DocumentationDemo shows code documentation
func DocumentationDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       CODE DOCUMENTATION DEMO                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. GODOC COMMENTS:")
	fmt.Println(`
// Package comments describe package purpose
// Appears at beginning of file before package declaration
// Should be complete sentence ending with period

// File: user.go
// Package users handles user-related operations
// It provides user management, authentication, and profile features.
package users

// Exported function comment
// GetUser retrieves a user by ID from the database.
// It returns nil if the user doesn't exist.
// Error is returned if database connection fails.
func GetUser(id int) (*User, error) {
    // ...
}

// Exported type comment
// User represents a user in the system
type User struct {
    // ID is the unique identifier for the user
    ID int
    
    // Name is the user's full name
    Name string
}

// Method comment
// Close closes the connection and releases resources
func (s *Service) Close() error {
    // ...
}

// Generate documentation:
// go doc package
// godoc -http=:6060  (local web server)
`)

	fmt.Println("\n2. EXAMPLE TESTS FOR DOCUMENTATION:")
	fmt.Println(`
// Example tests serve as executable documentation
// Appear in godoc output

// File: user_test.go

func ExampleGetUser() {
    user, err := GetUser(1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(user.Name)
    // Output: John Doe
}

func ExampleNewService() {
    service := NewService("localhost:5432")
    fmt.Println(service != nil)
    // Output: true
}

// Benefits:
// ✓ Docs are executable (always up to date)
// ✓ Shows expected behavior
// ✓ Tests basic functionality
// ✓ Appears in godoc
`)

	fmt.Println("✓ Documentation demonstrated")
}

// CommonPitfallsDemo shows mistakes to avoid
func CommonPitfallsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       COMMON PITFALLS DEMO                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. COMMON MISTAKES AND SOLUTIONS:")
	fmt.Println(`
// 1. Not checking errors
result, err := someFunction()  // ✗ Bad
if result != nil {             // Never checks error
    // ...
}

if err != nil {                // ✓ Good
    return fmt.Errorf("failed: %w", err)
}

// 2. Goroutine leaks
go someFunction()  // ✗ Bad - no control when it stops
                   // ✗ If someFunction loops forever, leaks goroutine

go func() {        // ✓ Good - has clear exit condition
    for {
        select {
        case <-ctx.Done():
            return
        case x := <-ch:
            process(x)
        }
    }
}()

// 3. Closing channels incorrectly
ch <- 1
close(ch)  // ✗ If ch still has receivers, panics
           // ✗ If other goroutines write, panics

// ✓ Good:
// - Only sender closes
// - Coordinator closes (both sender and receiver close)
// - Use context for cancellation instead

// 4. Race conditions
var counter int
go func() { counter++ }  // ✗ Data race
go func() { counter++ }  // ✗ Both access counter

var counter int
var mu sync.Mutex        // ✓ Mutex protects access
go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()

// 5. Defer in loops
for i := 0; i < 100; i++ {
    f, _ := os.Open(filename)
    defer f.Close()  // ✗ Bad! All defers run at end of function
                     // 99 files still open!
}

// ✓ Good: Use function
for i := 0; i < 100; i++ {
    process(filename)
}

func process(filename string) {
    f, _ := os.Open(filename)
    defer f.Close()  // Runs immediately
}
`)

	fmt.Println("\n2. PERFORMANCE ANTI-PATTERNS:")
	fmt.Println(`
// 1. Excessive allocations
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i)  // ✗ Allocates many times
}

result := make([]int, 0, 1000)  // ✓ Preallocate
for i := 0; i < 1000; i++ {
    result = append(result, i)
}

// 2. String concatenation in loops
var result string
for i := 0; i < 1000; i++ {
    result += fmt.Sprintf("Item %d\n", i)  // ✗ Inefficient
}

var buf bytes.Buffer              // ✓ Better
for i := 0; i < 1000; i++ {
    fmt.Fprintf(&buf, "Item %d\n", i)
}

// 3. Unmarshaling into interface{}
var data interface{}
json.Unmarshal(bytes, &data)     // ✗ Slow, weak typing

type Response struct {           // ✓ Better
    Name string ` + "`json:\"name\"`" + `
    Age  int    ` + "`json:\"age\"`" + `
}
var data Response
json.Unmarshal(bytes, &data)
`)

	fmt.Println("✓ Common pitfalls demonstrated")
}

// PerformanceOptimizationDemo shows performance tips
func PerformanceOptimizationDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       PERFORMANCE OPTIMIZATION DEMO                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. PROFILING AND BENCHMARKING:")
	fmt.Println(`
// Run benchmarks
go test -bench=. -benchmem

// Output shows:
// BenchmarkFunction-8    1000000    1234 ns/op    512 B/op    3 allocs/op
//                        iterations  time/op      memory/op   allocations

// Useful benchmarks:
// - Time per operation (ns/op)
// - Bytes allocated (B/op)
// - Allocations per operation (allocs/op)

// CPU profiling:
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

// Memory profiling:
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

// Trace:
go test -trace=trace.out
go tool trace trace.out
`)

	fmt.Println("\n2. OPTIMIZATION TECHNIQUES:")
	fmt.Println(`
// 1. Preallocate collections
data := make([]int, 0, expectedSize)  // Reserve capacity
data := make([]int, size)             // Exact size

// 2. Avoid unnecessary allocations
// Use pointers for large structs (reduces copying)
// Pass small structs by value (no allocation)

// 3. Use sync.Pool for temporary objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

buf := bufferPool.Get().(*bytes.Buffer)
defer func() {
    buf.Reset()
    bufferPool.Put(buf)
}()

// 4. Cache results
type Service struct {
    cache map[string]Result
    mu    sync.RWMutex
}

func (s *Service) GetResult(key string) Result {
    s.mu.RLock()
    if result, ok := s.cache[key]; ok {
        s.mu.RUnlock()
        return result
    }
    s.mu.RUnlock()
    
    result := compute(key)
    
    s.mu.Lock()
    s.cache[key] = result
    s.mu.Unlock()
    
    return result
}

// 5. Use efficient data structures
// Maps for lookups (O(1) average)
// Slices for sequential access
// Linked lists rarely needed (memory overhead)
`)

	fmt.Println("\n3. CONCURRENCY OPTIMIZATION:")
	fmt.Println(`
// 1. Use worker pools for many tasks
type WorkerPool struct {
    jobs    chan Job
    workers int
}

// 2. Buffer channels appropriately
ch := make(chan int)        // Unbuffered (0)
ch := make(chan int, 1)     // Small buffer (1)
ch := make(chan int, 100)   // Larger buffer

// Use unbuffered for synchronization
// Use buffered to decouple sender/receiver

// 3. Use context for cancellation
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
case result := <-ch:
    return result
}
`)

	fmt.Println("✓ Performance optimization demonstrated")
}

// BestPracticesSummaryDemo provides final summary
func BestPracticesSummaryDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       BEST PRACTICES SUMMARY                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("GO PHILOSOPHY SUMMARY:")
	fmt.Println(`
1. SIMPLICITY
   ✓ Code should be clear, not clever
   ✓ Prefer explicit over implicit
   ✓ Readability counts
   ✓ Simple is often better than optimal

2. COMPOSITION
   ✓ Use interfaces for abstraction
   ✓ Embed, don't inherit
   ✓ Small, focused types
   ✓ Compose large systems from small parts

3. CONCURRENCY
   ✓ Use goroutines for parallelism
   ✓ Use channels for synchronization
   ✓ "Share memory by communicating"
   ✓ Avoid shared mutable state

4. ERROR HANDLING
   ✓ Errors are values, not exceptions
   ✓ Handle errors explicitly
   ✓ Wrap errors with context
   ✓ Don't panic (except unrecoverable errors)

5. TESTING
   ✓ Write tests alongside code
   ✓ Table-driven tests for multiple scenarios
   ✓ Use mocks for isolation
   ✓ Benchmark to find bottlenecks

6. STANDARDS
   ✓ Follow Go conventions
   ✓ Use gofmt for formatting
   ✓ Use golint for style
   ✓ Write clear documentation

7. PACKAGES
   ✓ Small, focused packages
   ✓ Clear package boundaries
   ✓ Avoid circular dependencies
   ✓ Explicit is better than implicit

8. PERFORMANCE
   ✓ Profile before optimizing
   ✓ Avoid premature optimization
   ✓ Use benchmarks to measure
   ✓ Readability > micro-optimizations
`)

	fmt.Println("\nCHECKLIST FOR PRODUCTION CODE:")
	fmt.Println(`
□ All errors are checked and handled
□ All resources are cleaned up (defer, close)
□ Code has no race conditions
□ Code is documented (comments, examples)
□ Code has tests (unit, integration)
□ Performance is acceptable (benchmarked)
□ Logging is present for debugging
□ Errors have context and are wrapped
□ Goroutines are managed (no leaks)
□ Graceful shutdown implemented
□ Rate limiting/backpressure handled
□ Metrics/monitoring in place
□ Security best practices followed
□ Code follows Go style conventions
□ Dependencies are managed (go.mod)
`)

	fmt.Println("✓ Best practices summary demonstrated")
}

// RunBestPracticesExamples executes all best practices demos
func RunBestPracticesExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║   BEST PRACTICES AND PATTERNS - COMPREHENSIVE GUIDE     ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	CodeOrganizationDemo()
	NamingConventionsDemo()
	DocumentationDemo()
	CommonPitfallsDemo()
	PerformanceOptimizationDemo()
	BestPracticesSummaryDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL BEST PRACTICES EXAMPLES COMPLETE                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Organize code by feature, not by type")
	fmt.Println("✓ Use lowercase package names")
	fmt.Println("✓ Follow Go naming conventions (camelCase)")
	fmt.Println("✓ Write clear documentation (godoc)")
	fmt.Println("✓ Always check and handle errors")
	fmt.Println("✓ Avoid common pitfalls (goroutine leaks, races)")
	fmt.Println("✓ Use benchmarks and profiling for optimization")
	fmt.Println("✓ Readability over cleverness")
	fmt.Println("✓ Simplicity is a feature")
	fmt.Println("✓ Test your code comprehensively\n")
}
