package concepts

import (
	"fmt"
)

// ============================================================================
// TESTING AND BENCHMARKING IN GO
// ============================================================================
// This file covers:
// - Unit testing fundamentals
// - Table-driven tests
// - Testing patterns and best practices
// - Mocking and stubs
// - Benchmarking
// - Test fixtures
// ============================================================================

// TestingBasicsDemo shows basic testing concepts
func TestingBasicsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         TESTING BASICS DEMO                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. UNIT TEST STRUCTURE:")
	fmt.Println(`
// File: math_test.go (naming convention: _test.go)
package mypackage

import "testing"

// Test function signature: func TestFunctionName(t *testing.T)
func TestAdd(t *testing.T) {
    // Arrange: Set up test data
    a := 2
    b := 3
    expected := 5
    
    // Act: Call function being tested
    result := Add(a, b)
    
    // Assert: Verify result
    if result != expected {
        t.Errorf("Add(%d, %d) = %d, want %d", a, b, result, expected)
    }
}

// Test naming convention:
// Test{FunctionName}{Scenario}
func TestAddWithNegativeNumbers(t *testing.T) {
    result := Add(-2, -3)
    if result != -5 {
        t.Errorf("Expected -5, got %d", result)
    }
}

// Run tests:
// go test              (run all tests)
// go test -v           (verbose - show individual tests)
// go test -run TestAdd (run specific test)
// go test -race        (detect race conditions)
`)

	fmt.Println("\n2. ASSERTION PATTERNS:")
	fmt.Println(`
func TestWithErrors(t *testing.T) {
    value, err := parseInteger("123")
    
    // Check for errors
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    // Check value
    if value != 123 {
        t.Errorf("Got %d, want 123", value)
    }
    
    // Fail test immediately
    if value < 0 {
        t.Fatal("Value cannot be negative")  // Stops test
    }
    
    // Log helpful information
    t.Logf("Test passed with value: %d", value)
}
`)

	fmt.Println("✓ Testing basics demonstrated")
}

// TableDrivenTestsDemo shows table-driven test pattern
func TableDrivenTestsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       TABLE-DRIVEN TESTS DEMO                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. TABLE-DRIVEN TEST PATTERN:")
	fmt.Println(`
// Define test cases in a table
func TestMultiplyTableDriven(t *testing.T) {
    // Test cases: describe what should happen
    tests := []struct {
        name     string // Description
        a, b     int    // Inputs
        expected int    // Expected output
    }{
        {"positive numbers", 2, 3, 6},
        {"with zero", 5, 0, 0},
        {"negative numbers", -2, 3, -6},
        {"large numbers", 1000, 1000, 1000000},
    }
    
    // Run each test case
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Multiply(%d, %d) = %d, want %d",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Benefits:
// - Easy to add test cases (just add rows)
// - Reduces code duplication
// - Clear expected behavior
// - Good documentation of edge cases
`)

	fmt.Println("\n2. COMPLEX TABLE-DRIVEN TESTS:")
	fmt.Println(`
func TestParseJSONTableDriven(t *testing.T) {
    tests := []struct {
        name      string        // Test description
        input     string        // Input data
        wantData  interface{}   // Expected parsed data
        wantErr   bool          // Whether error is expected
        wantErrMsg string       // Error message substring
    }{
        {
            name:     "valid JSON",
            input:    ` + "`{\"name\": \"John\"}`" + `,
            wantData: map[string]string{"name": "John"},
            wantErr:  false,
        },
        {
            name:       "invalid JSON",
            input:      ` + "`{invalid}`" + `,
            wantErr:    true,
            wantErrMsg: "syntax error",
        },
        {
            name:     "empty input",
            input:    "",
            wantErr:  true,
            wantErrMsg: "empty",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            data, err := ParseJSON(tt.input)
            
            if tt.wantErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            if tt.wantErr && err != nil && !contains(err.Error(), tt.wantErrMsg) {
                t.Errorf("Error message doesn't contain %q", tt.wantErrMsg)
            }
            
            if !tt.wantErr && data != tt.wantData {
                t.Errorf("Got %v, want %v", data, tt.wantData)
            }
        })
    }
}
`)

	fmt.Println("✓ Table-driven tests demonstrated")
}

// MockingAndStubsDemo shows testing with mocks
func MockingAndStubsDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       MOCKING AND STUBS DEMO                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. INTERFACE-BASED TESTING (DEPENDENCY INJECTION):")
	fmt.Println(`
// Real implementation uses actual database
type RealUserRepository struct{}

func (r *RealUserRepository) GetUser(id int) (*User, error) {
    // Actual database call
    return queryDatabase(id)
}

// Mock implementation for testing
type MockUserRepository struct {
    users map[int]*User
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

// Service uses interface, not concrete type
type UserService struct {
    repo UserRepository // Interface, not concrete type
}

func (s *UserService) GetUserEmail(id int) (string, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Email, nil
}

// Interface definition
type UserRepository interface {
    GetUser(id int) (*User, error)
}

// Test with mock
func TestGetUserEmail(t *testing.T) {
    mock := &MockUserRepository{
        users: map[int]*User{
            1: &User{ID: 1, Email: "john@example.com"},
        },
    }
    
    service := &UserService{repo: mock}
    email, err := service.GetUserEmail(1)
    
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if email != "john@example.com" {
        t.Errorf("Got %s, want john@example.com", email)
    }
}
`)

	fmt.Println("\n2. STUB RESPONSES:")
	fmt.Println(`
// Stub HTTP client for testing
type StubHTTPClient struct {
    Response *http.Response
    Error    error
}

func (s *StubHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return s.Response, s.Error
}

// Use stub in tests
func TestFetchDataWithStub(t *testing.T) {
    stubResponse := &http.Response{
        StatusCode: 200,
        Body:       ioutil.NopCloser(strings.NewReader(` + "`{\"data\": \"test\"}`" + `)),
    }
    
    client := &StubHTTPClient{Response: stubResponse}
    fetcher := &DataFetcher{client: client}
    
    data, err := fetcher.Fetch("http://api.example.com")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if data != "test" {
        t.Errorf("Got %s, want test", data)
    }
}
`)

	fmt.Println("✓ Mocking and stubs demonstrated")
}

// BenchmarkingDemo shows performance testing
func BenchmarkingDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       BENCHMARKING DEMO                                ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. BENCHMARK STRUCTURE:")
	fmt.Println(`
// Benchmark function: func Benchmark{FunctionName}(b *testing.B)
func BenchmarkAdd(b *testing.B) {
    // Run the function b.N times
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

// Go automatically determines b.N for reasonable runtime
// Run benchmarks:
// go test -bench=.                (run all benchmarks)
// go test -bench=BenchmarkAdd    (run specific benchmark)
// go test -bench=. -benchmem      (show memory allocations)

// Example output:
// BenchmarkAdd-8    1000000000   1.03 ns/op
// (1 billion iterations, 1.03 nanoseconds per operation)
`)

	fmt.Println("\n2. BENCHMARK WITH SETUP/TEARDOWN:")
	fmt.Println(`
func BenchmarkComplexOperation(b *testing.B) {
    // Setup (runs once before benchmark)
    data := generateLargeDataset()
    
    // Reset benchmark timer (exclude setup from timing)
    b.ResetTimer()
    
    // Run b.N times
    for i := 0; i < b.N; i++ {
        ProcessData(data)
    }
}

// Benchmark with sub-benchmarks
func BenchmarkStringConcatenation(b *testing.B) {
    tests := []struct {
        name string
        fn   func([]string) string
    }{
        {"using +", concatenateWithPlus},
        {"using fmt.Sprint", concatenateWithSprint},
        {"using strings.Join", concatenateWithJoin},
    }
    
    strings := generateStrings(100)
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                tt.fn(strings)
            }
        })
    }
}

// Benchmarks show which approach is fastest
`)

	fmt.Println("\n3. BENCHMARK ANALYSIS:")
	fmt.Println(`
// Benchmark comparison:
go test -bench=. -benchmem

// Output shows:
// - Number of iterations (b.N)
// - Time per iteration (ns/op, µs/op, ms/op)
// - Allocations per iteration (allocs/op)
// - Bytes allocated per iteration (B/op)

// Example:
// BenchmarkAdd-8               1000000000   1.03 ns/op   0 B/op   0 allocs/op
// BenchmarkStrConcat-8         10000000    156 ns/op   32 B/op   1 allocs/op
// BenchmarkJSONMarshal-8        100000   10523 ns/op   512 B/op  12 allocs/op

// Interpretation:
// - Lower time is better
// - Fewer allocations is better (less garbage collection)
// - Less memory allocated is better
`)

	fmt.Println("✓ Benchmarking demonstrated")
}

// TestBestPracticesDemo shows testing best practices
func TestBestPracticesDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       TESTING BEST PRACTICES DEMO                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("1. TESTING BEST PRACTICES:")
	fmt.Println(`
// ✓ DO:
// - Test public API, not private functions
// - Test behavior, not implementation
// - Use table-driven tests for multiple scenarios
// - Mock external dependencies
// - Test error cases
// - Use clear, descriptive test names
// - Keep tests focused and independent
// - Run tests with -race flag to detect races

// ✗ DON'T:
// - Test private functions directly
// - Test implementation details
// - Create interdependent tests
// - Test external services (use mocks)
// - Write huge tests with many assertions
// - Ignore error cases
// - Leave hardcoded test data
`)

	fmt.Println("\n2. TEST FILE ORGANIZATION:")
	fmt.Println(`
// Project structure:
// myapp/
//   ├── user.go          (implementation)
//   ├── user_test.go     (tests for user.go)
//   ├── order.go         (implementation)
//   ├── order_test.go    (tests for order.go)
//   └── testdata/        (test fixtures and data files)
//       ├── valid.json
//       └── invalid.json

// Test data organization:
// testdata/
//   ├── fixtures/        (test fixtures)
//   ├── golden/          (expected outputs)
//   └── seeds/           (seed data)
`)

	fmt.Println("\n3. COVERAGE TESTING:")
	fmt.Println(`
// Check code coverage:
go test -cover

// Example output:
// ok  	mypackage	0.512s	coverage: 78.5% of statements

// Generate coverage report:
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

// Aim for:
// - Critical code: 90%+ coverage
// - Normal code: 70%+ coverage
// - Simple utilities: 50%+ coverage

// Note: 100% coverage doesn't mean bug-free!
// Focus on testing important behaviors, not just lines.
`)

	fmt.Println("✓ Testing best practices demonstrated")
}

// RealWorldTestingExampleDemo shows complete testing example
func RealWorldTestingExampleDemo() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    REAL WORLD TESTING EXAMPLE DEMO                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("COMPLETE TESTING EXAMPLE:")
	fmt.Println(`
// File: calculator.go
package calculator

type Calculator struct {
    lastResult float64
}

func (c *Calculator) Add(a, b float64) float64 {
    c.lastResult = a + b
    return c.lastResult
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    c.lastResult = a / b
    return c.lastResult, nil
}

// File: calculator_test.go
package calculator

import "testing"

func TestCalculatorAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive", 2, 3, 5},
        {"with zero", 5, 0, 5},
        {"negative", -2, 3, 1},
        {"decimals", 1.5, 2.5, 4.0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result := calc.Add(tt.a, tt.b)
            
            if result != tt.expected {
                t.Errorf("Add(%f, %f) = %f, want %f",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestCalculatorDivide(t *testing.T) {
    tests := []struct {
        name      string
        a, b      float64
        expected  float64
        wantErr   bool
    }{
        {"normal", 10, 2, 5, false},
        {"by zero", 10, 0, 0, true},
        {"decimals", 7.5, 2.5, 3, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            calc := &Calculator{}
            result, err := calc.Divide(tt.a, tt.b)
            
            if tt.wantErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            if !tt.wantErr && result != tt.expected {
                t.Errorf("Got %f, want %f", result, tt.expected)
            }
        })
    }
}

func BenchmarkAdd(b *testing.B) {
    calc := &Calculator{}
    for i := 0; i < b.N; i++ {
        calc.Add(123.45, 678.90)
    }
}

func BenchmarkDivide(b *testing.B) {
    calc := &Calculator{}
    for i := 0; i < b.N; i++ {
        calc.Divide(1000, 3)
    }
}
`)

	fmt.Println("✓ Real world testing example demonstrated")
}

// RunTestingAndBenchmarkExamples executes all testing demos
func RunTestingAndBenchmarkExamples() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    TESTING AND BENCHMARKING - COMPREHENSIVE GUIDE       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	TestingBasicsDemo()
	TableDrivenTestsDemo()
	MockingAndStubsDemo()
	BenchmarkingDemo()
	TestBestPracticesDemo()
	RealWorldTestingExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL TESTING AND BENCHMARKING EXAMPLES COMPLETE       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("✓ Test functions must be in *_test.go files")
	fmt.Println("✓ Use table-driven tests for multiple scenarios")
	fmt.Println("✓ Test error cases and edge cases")
	fmt.Println("✓ Mock external dependencies using interfaces")
	fmt.Println("✓ Benchmarks measure performance")
	fmt.Println("✓ Use -race flag to detect race conditions")
	fmt.Println("✓ Aim for meaningful coverage, not 100%")
	fmt.Println("✓ Keep tests independent and focused")
	fmt.Println("✓ Use subtests for organization")
	fmt.Println("✓ Benchmarks help find performance bottlenecks\n")
}
