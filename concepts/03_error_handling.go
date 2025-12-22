/*
ERROR HANDLING IN GO - The Right Way
============================================================================

Go's philosophy: "Errors are values"
Rather than using exceptions, Go uses return values for errors.
This requires explicit handling but leads to more robust code.

TOPICS COVERED:
- Error interface and error values
- Creating custom errors
- Error wrapping and unwrapping
- Deferred cleanup with errors
- Patterns for error handling
*/

package concepts

import (
	"errors"
	"fmt"
	"strconv"
)

// ============================================================================
// 1. THE ERROR INTERFACE
// ============================================================================

/*
ERROR INTERFACE:
- Any value with an Error() string method implements the error interface
- errors.New() creates a simple string-based error
- fmt.Errorf() creates formatted errors
- nil means no error occurred
*/

func ErrorInterfaceDemo() {
	fmt.Println("\n========== ERROR INTERFACE ==========")

	// Simple errors
	err1 := errors.New("something went wrong")
	fmt.Printf("Error 1: %v\n", err1)

	// Formatted errors
	value := 42
	err2 := fmt.Errorf("invalid value: %d is out of range", value)
	fmt.Printf("Error 2: %v\n", err2)

	// Checking for errors
	if err1 != nil {
		fmt.Println("Error 1 occurred")
	}

	if err2 != nil {
		fmt.Println("Error 2 occurred")
	}

	fmt.Println("\n✓ Error interface demonstrated")
}

// ============================================================================
// 2. CUSTOM ERRORS - Create Your Own
// ============================================================================

/*
CUSTOM ERRORS:
- Create types that implement the error interface
- More information than simple string errors
- Can add context and codes
- Better for programmatic error handling
*/

// ValidationError - custom error type
type ValidationError struct {
	Field   string
	Message string
}

// Implement error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// ParseError - custom error with code
type ParseError struct {
	Line   int
	Column int
	Text   string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse error at line %d, column %d: %s",
		e.Line, e.Column, e.Text)
}

// Function that returns custom error
func ValidateEmail(email string) error {
	if !contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "missing @ symbol",
		}
	}
	if !contains(email, ".") {
		return ValidationError{
			Field:   "email",
			Message: "missing domain extension",
		}
	}
	return nil
}

// Helper function
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		(substr[0] >= 1 && s[0] >= 1) // Simplified for demo
}

func CustomErrorsDemo() {
	fmt.Println("\n========== CUSTOM ERRORS ==========")

	// Valid email
	err := ValidateEmail("user@example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid email")
	}

	// Invalid email
	err = ValidateEmail("invalidemail")
	if err != nil {
		// Type assertion to get specific error type
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("Validation Error - Field: %s, Message: %s\n",
				ve.Field, ve.Message)
		}
	}

	// ParseError
	parseErr := ParseError{Line: 10, Column: 5, Text: "unexpected token"}
	fmt.Println(parseErr.Error())

	fmt.Println("\n✓ Custom errors demonstrated")
}

// ============================================================================
// 3. ERROR WRAPPING - Context and Chains
// ============================================================================

/*
ERROR WRAPPING:
- fmt.Errorf with %w preserves the original error
- Allows adding context while keeping original error
- errors.Is() checks wrapped errors
- errors.As() extracts wrapped errors
*/

// Convert string to integer with context
func ParseInteger(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		// Wrap error with context
		return 0, fmt.Errorf("failed to parse integer from '%s': %w", s, err)
	}
	return num, nil
}

// Function that can fail
func ProcessData(input string) error {
	num, err := ParseInteger(input)
	if err != nil {
		// Add more context
		return fmt.Errorf("data processing failed: %w", err)
	}

	if num < 0 {
		return fmt.Errorf("data processing failed: %w",
			errors.New("number must be positive"))
	}

	return nil
}

func ErrorWrappingDemo() {
	fmt.Println("\n========== ERROR WRAPPING ==========")

	// Successful case
	err := ProcessData("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully processed '42'")
	}

	// Error case
	err = ProcessData("not a number")
	if err != nil {
		fmt.Println("Error:", err)

		// Check if it's a specific error type
		if errors.Is(err, strconv.ErrRange) {
			fmt.Println("This is a range error")
		}
	}

	// Negative number case
	err = ProcessData("-5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n✓ Error wrapping demonstrated")
}

// ============================================================================
// 4. COMMON ERROR PATTERNS
// ============================================================================

/*
PATTERNS:
1. Multiple returns with error as last
2. Check errors immediately
3. Defer cleanup on errors
4. Wrap errors with context
5. Use meaningful error messages
*/

// Pattern 1: Multiple returns
func DivideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Pattern 2: Check errors immediately
func ReadProcessWrite(filename string) error {
	// Simulated functions
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	processed := process(data)

	err = writeFile(filename, processed)
	if err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	return nil
}

// Simulated functions
func readFile(name string) (string, error) {
	return "file content", nil
}

func process(data string) string {
	return data + " (processed)"
}

func writeFile(name, content string) error {
	return nil
}

// Pattern 3: Using defer for cleanup
func ProcessFile() error {
	// Open file (simulated)
	file := "data.txt"

	// Close in defer (always executes)
	defer func() {
		fmt.Printf("Cleaned up: closed %s\n", file)
	}()

	// Do processing
	if file == "" {
		return errors.New("empty filename")
	}

	return nil
}

func ErrorPatternsDemo() {
	fmt.Println("\n========== ERROR PATTERNS ==========")

	// Pattern 1
	result, err := DivideNumbers(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	_, err = DivideNumbers(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Pattern 2
	err = ReadProcessWrite("test.txt")
	if err != nil {
		fmt.Println("Pipeline error:", err)
	}

	// Pattern 3
	err = ProcessFile()
	if err != nil {
		fmt.Println("Processing error:", err)
	}

	fmt.Println("\n✓ Error patterns demonstrated")
}

// ============================================================================
// 5. PANIC AND RECOVER - When Errors Aren't Enough
// ============================================================================

/*
WHEN TO USE PANIC:
- Unrecoverable errors
- Programming errors (not user input errors)
- Use recover only in specific places

WHEN NOT TO USE PANIC:
- Regular error conditions
- User input validation
- Expected failures
*/

func SafeOperation(shouldFail bool) (result string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = "default value"
		}
	}()

	if shouldFail {
		panic("intentional panic for demonstration")
	}

	return "operation successful"
}

func PanicRecoverDemo() {
	fmt.Println("\n========== PANIC AND RECOVER ==========")

	result := SafeOperation(false)
	fmt.Printf("Result: %s\n", result)

	result = SafeOperation(true)
	fmt.Printf("Result after recovery: %s\n", result)

	fmt.Println("\n✓ Panic and recover demonstrated")
}

// ============================================================================
// 6. PRACTICAL EXAMPLE: Form Validation
// ============================================================================

type ValidationErrors struct {
	Errors map[string]string
}

func (ve ValidationErrors) Error() string {
	if len(ve.Errors) == 0 {
		return "no validation errors"
	}

	var msg string
	for field, err := range ve.Errors {
		msg += fmt.Sprintf("%s: %s; ", field, err)
	}
	return msg
}

// User struct
type User struct {
	Name  string
	Email string
	Age   int
}

// ValidateUser checks all user fields
func ValidateUser(user User) error {
	errors := make(map[string]string)

	// Validate name
	if user.Name == "" {
		errors["name"] = "name is required"
	}

	// Validate email
	if user.Email == "" {
		errors["email"] = "email is required"
	}

	// Validate age
	if user.Age < 18 {
		errors["age"] = "must be at least 18"
	}

	if len(errors) > 0 {
		return ValidationErrors{Errors: errors}
	}

	return nil
}

func PracticalErrorHandlingDemo() {
	fmt.Println("\n========== PRACTICAL EXAMPLE: Form Validation ==========")

	// Valid user
	user1 := User{Name: "Alice", Email: "alice@example.com", Age: 25}
	err := ValidateUser(user1)
	if err != nil {
		fmt.Println("Validation errors:", err)
	} else {
		fmt.Println("User validated successfully")
	}

	// Invalid user
	user2 := User{Name: "", Email: "bob@example.com", Age: 16}
	err = ValidateUser(user2)
	if err != nil {
		if ve, ok := err.(ValidationErrors); ok {
			fmt.Println("Validation failed:")
			for field, message := range ve.Errors {
				fmt.Printf("  - %s: %s\n", field, message)
			}
		}
	}

	fmt.Println("\n✓ Practical error handling demonstrated")
}

// ============================================================================
// BEST PRACTICES FOR ERROR HANDLING
// ============================================================================

/*
1. Return errors as values, not nil/true
2. Wrap errors with context using fmt.Errorf with %w
3. Check errors immediately
4. Use custom error types for programmatic checks
5. Don't ignore errors (even if you think you will)
6. Log errors appropriately
7. Use panic only for truly unrecoverable situations
8. Provide meaningful error messages
9. Use defer for cleanup
10. Don't use exceptions for control flow
*/

func BEST_PRACTICES() {
	s := `
ERROR HANDLING BEST PRACTICES:

✓ DO:
  - Return errors from functions
  - Wrap errors with context
  - Check errors immediately
  - Use custom error types
  - Provide meaningful messages
  - Use defer for cleanup
  - Log errors appropriately

✗ DON'T:
  - Ignore errors
  - Use panic for normal flow
  - Create generic error types
  - Return nil for success
  - Create error hierarchies
  - Catch all errors indiscriminately
  - Use errors for control flow
`
	fmt.Println(s)
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunErrorHandlingExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║      ERROR HANDLING IN GO - COMPLETE GUIDE             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	ErrorInterfaceDemo()
	CustomErrorsDemo()
	ErrorWrappingDemo()
	ErrorPatternsDemo()
	PanicRecoverDemo()
	PracticalErrorHandlingDemo()
	BEST_PRACTICES()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     ERROR HANDLING MASTERY - BUILD ROBUST CODE!        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
