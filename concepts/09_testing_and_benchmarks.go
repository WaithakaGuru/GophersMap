/*
TESTING AND BENCHMARKING IN GO - Complete Learning Guide
============================================================================

Testing is critical for reliable software. Go provides excellent built-in
testing support with the "testing" package. Benchmarking helps identify
performance bottlenecks.

TOPICS COVERED:
- Unit testing fundamentals and conventions
- Table-driven testing patterns
- Error and edge case testing
- Mocking and stub strategies
- Benchmarking and performance analysis
- Testing best practices and organization
*/

package concepts

import (
	"fmt"
)

// ============================================================================
// 1. UNIT TESTING FUNDAMENTALS
// ============================================================================

/*
GO TESTING CONVENTIONS:
- Test files named *_test.go
- Test function signature: func Test{FunctionName}(t *testing.T)
- Run tests: go test, go test -v, go test -run TestName
- Assertion pattern: if result != expected { t.Errorf(...) }

Testing methods on *testing.T:
- t.Errorf(format, args...)   : Fail but continue
- t.Error(args...)            : Fail but continue (no format)
- t.Fatalf(format, args...)   : Fail and stop
- t.Fatal(args...)            : Fail and stop
- t.Log(args...)              : Log message (verbose only)
- t.Logf(format, args...)     : Log with format
- t.Skip(args...)             : Skip this test
- t.Skipf(format, args...)    : Skip with reason
- t.SkipNow()                 : Skip and stop

Test organization:
- Arrange: Set up test data
- Act: Call function being tested
- Assert: Verify result
*/

func TestingBasicsDemo() {
	fmt.Println("\n========== UNIT TESTING FUNDAMENTALS ==========")

	// Example test function
	testAdd := func(t interface{}) {
		// Arrange
		a, b := 2, 3
		expected := 5

		// Act
		result := a + b

		// Assert
		if result != expected {
			// This is how you'd fail a test in a real test file
			// t.Errorf("Add(%d, %d) = %d, want %d", a, b, result, expected)
		}
	}

	testWithError := func(t interface{}) {
		// Test error handling
		result := 10

		// Check value
		if result < 0 {
			// t.Errorf("Got %d, want >= 0", result)
		}

		// Fatal stops the test immediately
		if result == -999 {
			// t.Fatal("Value is invalid - stop testing")
		}
	}

	fmt.Println("\n1. TEST FUNCTION SIGNATURE:")
	fmt.Println("   func Test{FunctionName}(t *testing.T) {")
	fmt.Println("       // Test code here")
	fmt.Println("   }")

	fmt.Println("\n2. ASSERTION PATTERN (Arrange-Act-Assert):")
	fmt.Println("   Arrange: Set up test data and dependencies")
	fmt.Println("   Act:     Call function being tested")
	fmt.Println("   Assert:  Verify result matches expected")
	_ = testAdd

	fmt.Println("\n3. ERROR REPORTING:")
	fmt.Println("   - t.Errorf(msg)   : Report failure, continue test")
	fmt.Println("   - t.Fatalf(msg)   : Report failure, stop test")
	fmt.Println("   - t.Log(msg)       : Log (verbose: go test -v)")
	_ = testWithError

	fmt.Println("\n4. RUNNING TESTS:")
	fmt.Println("   go test              : Run all tests")
	fmt.Println("   go test -v           : Verbose output")
	fmt.Println("   go test -run TestAdd : Run specific test")
	fmt.Println("   go test -race        : Detect race conditions")

	fmt.Println("\n✓ Unit testing fundamentals demonstrated")
}

// ============================================================================
// 2. TABLE-DRIVEN TESTING
// ============================================================================

/*
TABLE-DRIVEN TESTING:
- Define test cases as array of structs
- Each struct contains input and expected output
- Loop through test cases and run assertions
- Best practice in Go for testing multiple scenarios

Benefits:
- Easy to add test cases (just add rows)
- Reduces code duplication
- Clear documentation of behavior
- Easy to identify failing cases
- Good for edge cases and boundary conditions
*/

func TableDrivenTestsDemo() {
	fmt.Println("\n========== TABLE-DRIVEN TESTING ==========")

	// Example table-driven test
	tableDrivenExample := func() {
		// Define test cases in a table
		testCases := []struct {
			name     string
			a, b     int
			expected int
		}{
			{"positive numbers", 2, 3, 6},
			{"with zero", 5, 0, 0},
			{"negative numbers", -2, 3, -6},
			{"large numbers", 1000, 1000, 1000000},
		}

		// Run each test case
		for _, tc := range testCases {
			result := tc.a * tc.b
			if result != tc.expected {
				fmt.Printf("FAIL: %s - got %d, want %d\n", tc.name, result, tc.expected)
			}
		}
	}

	// Complex table-driven example
	complexExample := func() {
		testCases := []struct {
			name      string
			input     string
			expected  string
			wantErr   bool
			errSubstr string
		}{
			{
				name:     "valid input",
				input:    "hello",
				expected: "HELLO",
				wantErr:  false,
			},
			{
				name:      "empty input",
				input:     "",
				wantErr:   true,
				errSubstr: "empty",
			},
			{
				name:     "special chars",
				input:    "hello123",
				expected: "HELLO123",
				wantErr:  false,
			},
		}

		for _, tc := range testCases {
			_ = tc // Use test case
		}
	}

	fmt.Println("\n1. BASIC TABLE-DRIVEN TEST:")
	fmt.Println("   tests := []struct{")
	fmt.Println("       name string")
	fmt.Println("       input string")
	fmt.Println("       expected string")
	fmt.Println("   } {")
	fmt.Println("       {\"case1\", \"input1\", \"output1\"},")
	fmt.Println("       {\"case2\", \"input2\", \"output2\"},")
	fmt.Println("   }")
	_ = tableDrivenExample

	fmt.Println("\n2. COMPLEX TABLE-DRIVEN TEST:")
	fmt.Println("   Include:")
	fmt.Println("   - Test description (name)")
	fmt.Println("   - Input data")
	fmt.Println("   - Expected output")
	fmt.Println("   - Whether error is expected (wantErr)")
	fmt.Println("   - Error message to check (errSubstr)")
	_ = complexExample

	fmt.Println("\n3. RUNNING TEST CASES:")
	fmt.Println("   for _, tc := range testCases {")
	fmt.Println("       result := functionToTest(tc.input)")
	fmt.Println("       if result != tc.expected {")
	fmt.Println("           t.Errorf(\"%s: got %v, want %v\",")
	fmt.Println("               tc.name, result, tc.expected)")
	fmt.Println("       }")
	fmt.Println("   }")

	fmt.Println("\n4. SUB-TESTS FOR ORGANIZATION:")
	fmt.Println("   for _, tc := range testCases {")
	fmt.Println("       t.Run(tc.name, func(t *testing.T) {")
	fmt.Println("           // Test tc here")
	fmt.Println("       })")
	fmt.Println("   }")
	fmt.Println("   Output: TestFunc/case1, TestFunc/case2, etc.")

	fmt.Println("\n✓ Table-driven testing demonstrated")
}

// ============================================================================
// 3. ERROR AND EDGE CASE TESTING
// ============================================================================

/*
TESTING ERROR CASES:
- Always test error conditions
- Check nil errors when success expected
- Check non-nil errors when failure expected
- Verify error messages contain expected text
- Test boundary conditions and edge cases

Edge cases to test:
- Empty inputs
- Nil values
- Large values
- Negative values
- Zero values
- Invalid types/formats
- Concurrent access (with -race)
*/

func ErrorHandlingTestingDemo() {
	fmt.Println("\n========== ERROR AND EDGE CASE TESTING ==========")

	testErrorCases := func() {
		// Test structure for error cases
		testCases := []struct {
			name        string
			input       string
			wantErr     bool
			errContains string
		}{
			{"valid", "hello", false, ""},
			{"empty", "", true, "empty"},
			{"invalid", "!!!invalid!!!", true, "invalid"},
		}

		for _, tc := range testCases {
			if tc.wantErr {
				// Error expected
				// err := functionThatErrors(tc.input)
				// if err == nil {
				//     t.Error("Expected error, got nil")
				// }
			} else {
				// No error expected
				// _, err := functionThatMightError(tc.input)
				// if err != nil {
				//     t.Errorf("Unexpected error: %v", err)
				// }
			}
		}
	}

	fmt.Println("\n1. ERROR TESTING PATTERNS:")
	fmt.Println("   if err != nil { t.Errorf(...) }     : Unexpected error")
	fmt.Println("   if err == nil { t.Error(...) }      : Expected error")
	fmt.Println("   if !strings.Contains(...) { ... }   : Check error message")

	fmt.Println("\n2. EDGE CASES TO TEST:")
	fmt.Println("   - Empty inputs (empty string, empty slice)")
	fmt.Println("   - Nil values")
	fmt.Println("   - Boundary values (min, max, zero)")
	fmt.Println("   - Negative numbers")
	fmt.Println("   - Very large numbers")
	fmt.Println("   - Invalid types/formats")
	_ = testErrorCases

	fmt.Println("\n3. EXAMPLE TABLE:")
	fmt.Println("   {\"zero\", 0, true, \"cannot divide by zero\"},")
	fmt.Println("   {\"negative\", -5, false, \"\"},")
	fmt.Println("   {\"very large\", 999999999, false, \"\"},")

	fmt.Println("\n✓ Error and edge case testing demonstrated")
}

// ============================================================================
// 4. MOCKING AND TESTING WITH INTERFACES
// ============================================================================

/*
MOCKING STRATEGY:
- Use interfaces to decouple dependencies
- Create mock implementations for testing
- Avoid testing external services (network, database)
- Inject dependencies via constructors or fields

Benefits of mocks:
- Fast tests (no network/database calls)
- Predictable results
- Test error conditions easily
- Test rare conditions (network timeouts)
- Parallel test execution
*/

func MockingAndStubsDemo() {
	fmt.Println("\n========== MOCKING AND STUBS ==========")

	// Example: Interface-based design
	// type UserRepository interface {
	//     GetUser(id int) (*User, error)
	// }

	// Real implementation
	// type RealUserRepository struct{}
	// func (r *RealUserRepository) GetUser(id int) (*User, error) {
	//     // Database call
	// }

	// Mock for testing
	// type MockUserRepository struct {
	//     users map[int]*User
	// }
	// func (m *MockUserRepository) GetUser(id int) (*User, error) {
	//     user, exists := m.users[id]
	//     if !exists {
	//         return nil, fmt.Errorf("not found")
	//     }
	//     return user, nil
	// }

	// Service uses interface
	// type UserService struct {
	//     repo UserRepository
	// }
	// func (s *UserService) GetEmail(id int) (string, error) {
	//     user, err := s.repo.GetUser(id)
	//     // ...
	// }

	// Test with mock
	// func TestGetEmail(t *testing.T) {
	//     mock := &MockUserRepository{
	//         users: map[int]*User{
	//             1: &User{Email: "test@example.com"},
	//         },
	//     }
	//     service := &UserService{repo: mock}
	//     email, _ := service.GetEmail(1)
	//     if email != "test@example.com" {
	//         t.Errorf("got %s, want test@example.com", email)
	//     }
	// }

	fmt.Println("\n1. INTERFACE-BASED TESTING:")
	fmt.Println("   Define interface for dependency:")
	fmt.Println("   type Repository interface {")
	fmt.Println("       Get(id int) (*Data, error)")
	fmt.Println("   }")

	fmt.Println("\n2. INJECT INTERFACE, NOT CONCRETE TYPE:")
	fmt.Println("   type Service struct {")
	fmt.Println("       repo Repository  // Interface!")
	fmt.Println("   }")

	fmt.Println("\n3. CREATE MOCK IMPLEMENTATION:")
	fmt.Println("   type MockRepository struct { ... }")
	fmt.Println("   func (m *MockRepository) Get(...) { ... }")

	fmt.Println("\n4. USE MOCK IN TESTS:")
	fmt.Println("   mock := &MockRepository{ ... }")
	fmt.Println("   service := &Service{repo: mock}")
	fmt.Println("   result := service.someMethod()")

	fmt.Println("\n✓ Mocking and stubs demonstrated")
}

// ============================================================================
// 5. BENCHMARKING
// ============================================================================

/*
BENCHMARKING:
- Measure function performance
- Benchmark function signature: func Benchmark{Name}(b *testing.B)
- Run benchmarks: go test -bench=.
- Go automatically determines iterations (b.N)

Benchmark output:
- BenchmarkName-8    1000000    1234 ns/op   128 B/op    4 allocs/op
  - Name and cores
  - Iterations (b.N)
  - Time per operation (ns/op)
  - Bytes allocated per op (B/op)
  - Allocations per op (allocs/op)

Best practices:
- Use b.ResetTimer() after setup
- Run with -benchmem to see allocations
- Compare before/after for optimizations
- Focus on allocation reduction (affects GC)
*/

func BenchmarkingDemo() {
	fmt.Println("\n========== BENCHMARKING ==========")

	// Example benchmark
	exampleBench := func() {
		iterations := 1000000
		operations := 0
		for i := 0; i < iterations; i++ {
			// Simulate operation
			operations++
		}
	}

	// Benchmark with setup
	benchWithSetup := func() {
		// Setup phase (not timed)
		data := make([]int, 0, 1000)

		// Simulate reset timer
		// b.ResetTimer()

		// Benchmark phase
		for i := 0; i < 1000; i++ {
			data = append(data, i)
		}
	}

	fmt.Println("\n1. BENCHMARK FUNCTION:")
	fmt.Println("   func BenchmarkAdd(b *testing.B) {")
	fmt.Println("       for i := 0; i < b.N; i++ {")
	fmt.Println("           Add(2, 3)")
	fmt.Println("       }")
	fmt.Println("   }")

	fmt.Println("\n2. BENCHMARK COMMANDS:")
	fmt.Println("   go test -bench=.              : Run all benchmarks")
	fmt.Println("   go test -bench=BenchmarkAdd   : Run specific")
	fmt.Println("   go test -bench=. -benchmem    : Show memory")
	fmt.Println("   go test -bench=. -count=5     : Run 5 times")

	_ = exampleBench
	_ = benchWithSetup

	fmt.Println("\n3. BENCHMARK OUTPUT INTERPRETATION:")
	fmt.Println("   BenchmarkAdd-8    1000000000    1.03 ns/op    0 B/op    0 allocs/op")
	fmt.Println("   - Name and cores (-8 = 8 cores)")
	fmt.Println("   - Iterations (1 billion times)")
	fmt.Println("   - Time per operation (1.03 nanoseconds)")
	fmt.Println("   - Bytes per operation (0)")
	fmt.Println("   - Allocations per operation (0)")

	fmt.Println("\n4. OPTIMIZATION STRATEGY:")
	fmt.Println("   - Reduce allocations (allocs/op)")
	fmt.Println("   - Reduce memory usage (B/op)")
	fmt.Println("   - Optimize hot paths (high frequency)")
	fmt.Println("   - Profile with pprof for details")

	fmt.Println("\n✓ Benchmarking demonstrated")
}

// ============================================================================
// 6. TESTING BEST PRACTICES
// ============================================================================

/*
TESTING BEST PRACTICES:
- Test public API, not private functions
- Test behavior, not implementation details
- Use table-driven tests for multiple scenarios
- Keep tests independent (no shared state)
- Use mocks for external dependencies
- Test error paths thoroughly
- Use clear, descriptive test names
- Check coverage (aim for 70%+ on critical code)
- Run tests with -race to detect races
- Use -timeout to catch hanging tests
*/

func TestBestPracticesDemo() {
	fmt.Println("\n========== TESTING BEST PRACTICES ==========")

	fmt.Println("\n1. DO:")
	fmt.Println("   ✓ Test public functions/methods")
	fmt.Println("   ✓ Test behavior, not implementation")
	fmt.Println("   ✓ Use table-driven tests")
	fmt.Println("   ✓ Keep tests independent")
	fmt.Println("   ✓ Mock external dependencies")
	fmt.Println("   ✓ Test error cases")
	fmt.Println("   ✓ Use descriptive names")
	fmt.Println("   ✓ Run with -race flag")

	fmt.Println("\n2. DON'T:")
	fmt.Println("   ✗ Test private functions")
	fmt.Println("   ✗ Test implementation details")
	fmt.Println("   ✗ Create interdependent tests")
	fmt.Println("   ✗ Call external services")
	fmt.Println("   ✗ Write massive tests")
	fmt.Println("   ✗ Skip error cases")
	fmt.Println("   ✗ Hardcode test data")

	fmt.Println("\n3. TEST NAMING CONVENTION:")
	fmt.Println("   TestFunctionName          : Basic test")
	fmt.Println("   TestFunctionNameErrorCase : Error handling")
	fmt.Println("   TestFunctionNameEdgeCase  : Edge cases")
	fmt.Println("   go test -v shows clear test names")

	fmt.Println("\n4. CODE COVERAGE:")
	fmt.Println("   go test -cover          : Show coverage %")
	fmt.Println("   go test -coverprofile=coverage.out")
	fmt.Println("   go tool cover -html=coverage.out")
	fmt.Println("   Target: 70%+ for important code")

	fmt.Println("\n5. RUNNING TESTS WITH OPTIONS:")
	fmt.Println("   go test -timeout 10s     : Timeout after 10s")
	fmt.Println("   go test -parallel 4      : Run 4 tests in parallel")
	fmt.Println("   go test -failfast        : Stop on first failure")
	fmt.Println("   go test -count=3         : Run 3 times")

	fmt.Println("\n✓ Testing best practices demonstrated")
}

// ============================================================================
// 7. REAL-WORLD TESTING EXAMPLE
// ============================================================================

/*
COMPLETE TESTING EXAMPLE:
- Function to test: Calculator with Add and Divide
- Test organization: Basic tests + edge cases
- Table-driven for multiple scenarios
- Error handling for invalid inputs
- Benchmarks to measure performance
*/

func RealWorldTestingExampleDemo() {
	fmt.Println("\n========== REAL-WORLD TESTING EXAMPLE ==========")

	// Simple function to test
	add := func(a, b int) int {
		return a + b
	}

	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	}

	// Example test cases structure
	testCases := []struct {
		name     string
		a, b     int
		expected int
		wantErr  bool
	}{
		{"positive numbers", 2, 3, 5, false},
		{"zero case", 5, 0, 5, false},
		{"negative", -2, 3, 1, false},
		{"large numbers", 1000, 2000, 3000, false},
	}

	fmt.Println("\n1. FUNCTION UNDER TEST:")
	fmt.Println("   func Add(a, b int) int {")
	fmt.Println("       return a + b")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func Divide(a, b int) (int, error) {")
	fmt.Println("       if b == 0 {")
	fmt.Println("           return 0, fmt.Errorf(\"division by zero\")")
	fmt.Println("       }")
	fmt.Println("       return a / b, nil")
	fmt.Println("   }")

	fmt.Println("\n2. TABLE-DRIVEN TEST:")
	fmt.Println("   func TestAdd(t *testing.T) {")
	fmt.Println("       tests := []struct{ ... }")
	fmt.Println("       for _, tt := range tests {")
	fmt.Println("           t.Run(tt.name, func(t *testing.T) { ... })")
	fmt.Println("       }")
	fmt.Println("   }")

	fmt.Println("\n3. TEST CASES INCLUDED:")
	for _, tc := range testCases {
		result := add(tc.a, tc.b)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("   %s %s: Add(%d, %d) = %d (expected %d)\n",
			status, tc.name, tc.a, tc.b, result, tc.expected)
	}

	fmt.Println("\n4. ERROR HANDLING TEST:")
	fmt.Println("   func TestDivide(t *testing.T) {")
	fmt.Println("       result, err := Divide(10, 0)")
	fmt.Println("       if err == nil {")
	fmt.Println("           t.Error(\"Expected error, got nil\")")
	fmt.Println("       }")
	fmt.Println("   }")

	testDivideResult, testDivideErr := divide(10, 0)
	_ = testDivideResult
	if testDivideErr != nil {
		fmt.Printf("   ✓ Division by zero correctly returns error: %v\n", testDivideErr)
	}

	fmt.Println("\n5. BENCHMARK EXAMPLE:")
	fmt.Println("   func BenchmarkAdd(b *testing.B) {")
	fmt.Println("       for i := 0; i < b.N; i++ {")
	fmt.Println("           Add(123, 456)")
	fmt.Println("       }")
	fmt.Println("   }")

	fmt.Println("\n✓ Real-world testing example demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

// RunTestingAndBenchmarkExamples executes all testing demonstrations
func RunTestingAndBenchmarkExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    TESTING AND BENCHMARKING - COMPREHENSIVE GUIDE      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	TestingBasicsDemo()
	TableDrivenTestsDemo()
	ErrorHandlingTestingDemo()
	MockingAndStubsDemo()
	BenchmarkingDemo()
	TestBestPracticesDemo()
	RealWorldTestingExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALL TESTING AND BENCHMARKING EXAMPLES COMPLETE      ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Println("\nKEY TAKEAWAYS:")
	fmt.Println("✓ Test files named *_test.go with func Test{Name}(t *testing.T)")
	fmt.Println("✓ Use table-driven tests for multiple scenarios")
	fmt.Println("✓ Always test error cases and edge cases")
	fmt.Println("✓ Mock external dependencies using interfaces")
	fmt.Println("✓ Keep tests independent and focused")
	fmt.Println("✓ Use benchmarks to measure performance")
	fmt.Println("✓ Focus on allocation reduction in benchmarks")
	fmt.Println("✓ Run tests with -race to detect race conditions")
	fmt.Println("✓ Aim for 70%+ code coverage on critical paths")
	fmt.Println("✓ Test behavior, not implementation details\n")
}
