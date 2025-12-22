/*
GO FUNDAMENTALS - Complete Learning Guide
============================================================================

This file covers the foundational concepts that every Go programmer must
understand. These are the building blocks for everything you'll build in Go.

TOPICS COVERED:
- Variables and Constants
- Data Types (int, float, string, bool)
- Type Conversions
- Functions and Methods
- Control Flow (if, switch, loops)
- Packages and Imports
- Defer, Panic, Recover
*/

package concepts

import (
	"fmt"
)

// ============================================================================
// 1. VARIABLES AND CONSTANTS
// ============================================================================

/*
VARIABLES:
- Declared with 'var', 'const', or ':=' (short declaration)
- Types are explicit or inferred
- Go is statically typed but infers types from values
- Must use declared variables (unused variables cause compilation error)
- Exported (public): PascalCase, Unexported (private): camelCase
*/

func VarAndConstantsDemo() {
	fmt.Println("\n========== VARIABLES & CONSTANTS ==========")

	// 1. Explicit type declaration
	var age int = 25
	var name string = "Alice"
	var isStudent bool = true
	var price float64 = 19.99

	fmt.Printf("Age: %d, Name: %s, Student: %v, Price: %.2f\n",
		age, name, isStudent, price)

	// 2. Type inference - Go figures out the type
	var country = "Kenya"     // Type inferred as string
	var population = 50000000 // Type inferred as int

	fmt.Printf("Country: %s, Population: %d\n", country, population)

	// 3. Short declaration (only inside functions)
	// This is the most idiomatic way in Go
	city := "Nairobi"   // Automatically inferred as string
	temperature := 28.5 // Automatically inferred as float64

	fmt.Printf("City: %s, Temp: %.1f°C\n", city, temperature)

	// 4. Multiple declarations
	var (
		firstName = "John"
		lastName  = "Doe"
		age2      = 30
	)
	fmt.Printf("Full Name: %s %s, Age: %d\n", firstName, lastName, age2)

	// 5. Constants - values that cannot change
	const PI = 3.14159
	const author = "Go Team"
	const maxUsers int = 1000

	fmt.Printf("π ≈ %.5f, Author: %s, Max Users: %d\n", PI, author, maxUsers)

	// 6. Blank identifier - for unused variables
	_ = "This value is ignored"

	fmt.Println("\n✓ Variables and Constants demonstrated")
}

// ============================================================================
// 2. DATA TYPES IN DEPTH
// ============================================================================

/*
GO DATA TYPES:
- Integer: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- Float: float32, float64
- Complex: complex64, complex128
- String: immutable sequence of bytes
- Boolean: true or false
- Rune: single Unicode character (int32)
- Byte: single byte (uint8)
*/

func DataTypesDemo() {
	fmt.Println("\n========== DATA TYPES IN DEPTH ==========")

	// INTEGER TYPES
	fmt.Println("--- INTEGER TYPES ---")
	var signedInt int = -42
	var unsignedInt uint = 42
	var smallInt int8 = 127 // Range: -128 to 127
	var largeInt int64 = 9223372036854775807

	fmt.Printf("Signed: %d, Unsigned: %d, Small: %d, Large: %d\n",
		signedInt, unsignedInt, smallInt, largeInt)

	// FLOAT TYPES - Used for decimal numbers
	fmt.Println("\n--- FLOAT TYPES ---")
	var pi float32 = 3.14
	var euler float64 = 2.718281828

	fmt.Printf("π (float32): %v, e (float64): %v\n", pi, euler)

	// STRING - Immutable sequence of characters
	fmt.Println("\n--- STRING TYPES ---")
	var message string = "Hello, Go!"
	var empty string = ""
	var multiline = `This is
a multiline
string`

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Empty string length: %d\n", len(empty))
	fmt.Printf("Multiline:\n%s\n", multiline)

	// BOOLEAN
	fmt.Println("\n--- BOOLEAN ---")
	var isActive bool = true
	var isEmpty bool = len(message) == 0

	fmt.Printf("Is Active: %v, Is Empty: %v\n", isActive, isEmpty)

	// RUNE - Single Unicode character
	fmt.Println("\n--- RUNE (Unicode Character) ---")
	var char rune = 'A'
	var emoji rune = '😀'
	var number rune = '5'

	fmt.Printf("Char: %c (code: %d)\n", char, char)
	fmt.Printf("Emoji: %c (code: %d)\n", emoji, emoji)
	fmt.Printf("Number: %c (code: %d)\n", number, number)

	// BYTE
	fmt.Println("\n--- BYTE ---")
	var singleByte byte = 65 // ASCII 'A'
	var byteValue byte = 255 // Max value

	fmt.Printf("Byte 65: %c, Max byte: %d\n", singleByte, byteValue)

	fmt.Println("\n✓ Data types demonstrated")
}

// ============================================================================
// 3. TYPE CONVERSIONS - CRITICAL CONCEPT
// ============================================================================

/*
IMPORTANT: Go does NOT allow implicit type conversion.
You must explicitly convert between types.
This prevents subtle bugs that occur in languages with implicit conversions.
*/

func TypeConversionDemo() {
	fmt.Println("\n========== TYPE CONVERSION ==========")

	// Converting between numeric types
	var intValue int = 42
	var floatValue float64 = float64(intValue) // int → float64
	var uint8Value uint8 = uint8(intValue)     // int → uint8

	fmt.Printf("int: %d → float64: %.1f → uint8: %d\n",
		intValue, floatValue, uint8Value)

	// String to numeric
	var str string = "123"
	fmt.Println(str)
	// Note: strconv package is needed for string ↔ number conversion
	// This is shown in a separate example with imports

	// Numeric to string - using fmt
	numToStr := fmt.Sprintf("%d", 456) // Convert int to string
	fmt.Printf("Number as string: %s (type: string)\n", numToStr)

	// ASCII codes and runes
	var asciiCode int = 65
	var character rune = rune(asciiCode)
	fmt.Printf("ASCII %d = %c\n", asciiCode, character)

	// Byte to string
	byteSlice := []byte{72, 101, 108, 108, 111} // "Hello"
	stringFromBytes := string(byteSlice)
	fmt.Printf("Bytes %v → String: %s\n", byteSlice, stringFromBytes)

	fmt.Println("\n✓ Type conversion demonstrated")
}

// ============================================================================
// 4. FUNCTIONS - THE BUILDING BLOCKS
// ============================================================================

/*
FUNCTION DECLARATION:
func functionName(param1 type, param2 type) returnType {
    // function body
}

KEY CONCEPTS:
- Functions are first-class citizens (can be passed as values)
- Multiple return values are supported
- Named return values
- Variadic parameters (...args)
- Defer - execute at function exit
*/

// Simple function - takes parameters, returns single value
func add(a, b int) int {
	return a + b
}

// Multiple return values - very common in Go
func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return numerator / denominator, nil
}

// Named return values - values are initialized to zero values
func multiplyWithoutError(a, b int) (product int, quadruple int) {
	product = a * b // Named return value
	quadruple = product * 4
	return // Can return without specifying values if named
}

// Variadic function - accepts variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that takes another function as parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Anonymous function (closure) - function without a name
func counterFactory() func() int {
	count := 0
	return func() int { // Anonymous function that has access to 'count'
		count++
		return count
	}
}

func FunctionsDemo() {
	fmt.Println("\n========== FUNCTIONS ==========")

	// Basic function
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// Named return values
	product, quad := multiplyWithoutError(4, 5)
	fmt.Printf("4 × 5 = %d, × 4 = %d\n", product, quad)

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)

	// Function as parameter (Higher-order function)
	multiply := func(a, b int) int { return a * b }
	result = applyOperation(6, 7, multiply)
	fmt.Printf("applyOperation(6, 7, multiply) = %d\n", result)

	// Closure - function with access to outer scope
	counter := counterFactory()
	fmt.Printf("Counter: %d, %d, %d\n", counter(), counter(), counter())

	fmt.Println("\n✓ Functions demonstrated")
}

// ============================================================================
// 5. CONTROL FLOW - if, switch, for
// ============================================================================

/*
CONTROL FLOW:
- if/else if/else - conditional execution
- switch - multiple conditions
- for - the only loop in Go (no while, do-while)
- for range - iterate over collections
- break/continue - loop control
*/

func ControlFlowDemo() {
	fmt.Println("\n========== CONTROL FLOW ==========")

	// IF/ELSE
	fmt.Println("--- IF/ELSE ---")
	age := 25
	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teen")
	} else if age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// IF with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	}

	// SWITCH - cleaner than multiple if/else
	fmt.Println("\n--- SWITCH ---")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	// Switch with fallthrough
	fmt.Println("\n--- SWITCH FALLTHROUGH ---")
	grade := 'B'
	switch grade {
	case 'A', 'B': // Multiple cases
		fmt.Println("Excellent")
		fallthrough // Execute next case too
	case 'C':
		fmt.Println("Good")
	case 'D', 'F':
		fmt.Println("Needs improvement")
	}

	// FOR LOOP - basic
	fmt.Println("\n--- FOR LOOP ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// FOR as while loop
	fmt.Println("\n--- FOR AS WHILE ---")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// FOR RANGE - iterate over collections
	fmt.Println("\n--- FOR RANGE ---")
	numbers := []int{10, 20, 30, 40}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Using _ to ignore index
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// BREAK and CONTINUE
	fmt.Println("\n--- BREAK & CONTINUE ---")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skip iteration
		}
		if i == 5 {
			break // Exit loop
		}
		fmt.Printf("Number: %d\n", i)
	}

	fmt.Println("\n✓ Control flow demonstrated")
}

// ============================================================================
// 6. DEFER - Cleanup and Finalization
// ============================================================================

/*
DEFER:
- Schedules a function call to run when the surrounding function returns
- Very useful for cleanup operations (closing files, releasing locks, etc.)
- Deferred calls are executed in LIFO order (Last In, First Out)
- Arguments are evaluated immediately, but the function call is deferred
*/

func DeferDemo() {
	fmt.Println("\n========== DEFER ==========")

	// Basic defer
	fmt.Println("--- Basic Defer ---")
	defer fmt.Println("This executes last")
	fmt.Println("This executes first")
	fmt.Println("This executes second")

	// Multiple defers - LIFO order
	fmt.Println("\n--- Multiple Defers (LIFO) ---")
	defer fmt.Println("3. Third defer")
	defer fmt.Println("2. Second defer")
	defer fmt.Println("1. First defer")

	// Real-world use case: File handling
	fmt.Println("\n--- Defer for Cleanup ---")
	filename := "test.txt"
	fmt.Printf("Simulating: Opening file '%s'\n", filename)
	defer fmt.Printf("Closing file '%s'\n", filename) // Automatically called at end
	fmt.Println("Working with file...")

	// Defer with function call
	fmt.Println("\n--- Defer with Arguments ---")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("Defer called with i=%d\n", i) // i is captured immediately
	}

	fmt.Println("\n✓ Defer demonstrated")
}

// ============================================================================
// 7. PANIC AND RECOVER
// ============================================================================

/*
PANIC:
- Stops normal execution and starts unwinding the stack
- Deferred functions still execute during unwinding
- Used for unrecoverable errors

RECOVER:
- Catches panics
- Only works inside deferred functions
- Returns nil if no panic occurred
- Must be used carefully
*/

func safeDivide(a, b int) (result int) {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Set a safe default
		}
	}()

	// This will panic if b == 0
	if b == 0 {
		panic("division by zero!")
	}

	return a / b
}

func PanicAndRecoverDemo() {
	fmt.Println("\n========== PANIC AND RECOVER ==========")

	// Safe division that recovers from panic
	result1 := safeDivide(10, 2)
	fmt.Printf("10 / 2 = %d\n", result1)

	result2 := safeDivide(10, 0) // This will panic but recover
	fmt.Printf("10 / 0 = %d (recovered)\n", result2)

	fmt.Println("\n✓ Panic and recover demonstrated")
}

// ============================================================================
// 8. SCOPE AND VISIBILITY
// ============================================================================

/*
SCOPE:
- Block scope: Variables exist within { }
- Function scope: Variables accessible throughout function
- Package scope: Exported (Capital) vs Unexported (lowercase)
- Shadowing: Inner scope can redeclare outer variables

VISIBILITY:
- Exported: Starts with uppercase letter (MyFunction)
- Unexported: Starts with lowercase letter (myFunction)
- Only applies at package level
*/

// Exported function - can be imported from other packages
func PublicFunction() {
	fmt.Println("This function can be called from other packages")
}

// Unexported function - only available in this package
func privateFunction() {
	fmt.Println("This function is private to this package")
}

func ScopeDemo() {
	fmt.Println("\n========== SCOPE AND VISIBILITY =========")

	// Package-level variable (if declared outside functions)
	// packageVar := "Package scope"

	// Function-level variable
	functionVar := "Function scope"
	fmt.Println(functionVar)

	// Block scope
	{
		blockVar := "Block scope"
		fmt.Println(blockVar)
		// fmt.Println(innerVar) // Would be error - innerVar not in this block
	}
	// fmt.Println(blockVar) // Would be error - blockVar out of scope

	// Shadowing - inner scope redefines outer variable
	x := "outer"
	fmt.Printf("x (outer): %s\n", x)
	{
		x := "inner" // Shadows outer x
		fmt.Printf("x (inner): %s\n", x)
	}
	fmt.Printf("x (outer again): %s\n", x) // Back to outer

	fmt.Println("\n✓ Scope and visibility demonstrated")
}

// ============================================================================
// 9. BEST PRACTICES FOR FUNDAMENTALS
// ============================================================================

/*
BEST PRACTICES:
1. Use meaningful variable names (not x, y, z unless in math)
2. Constants for values that don't change
3. Multiple return values (especially error as last return)
4. Short variable scopes (declare close to use)
5. Use for loops, not while loops
6. Prefer switch over multiple if/else
7. Use defer for cleanup operations
8. Handle errors explicitly
9. Export only what's needed
10. Use composition over inheritance
*/

// Good practice example: Clear function with error handling
func parseAge(ageStr string) (int, error) {
	// Meaningful names, clear intent
	age := 0

	// This would use strconv.Atoi in real code
	_, err := fmt.Sscanf(ageStr, "%d", &age)
	if err != nil {
		return 0, fmt.Errorf("invalid age format: %w", err)
	}

	if age < 0 || age > 150 {
		return 0, fmt.Errorf("age out of reasonable range: %d", age)
	}

	return age, nil
}

func BestPracticesDemo() {
	fmt.Println("\n========== BEST PRACTICES ==========")

	age, err := parseAge("25")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Successfully parsed age: %d\n", age)

	// Invalid age
	_, err = parseAge("-5")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	fmt.Println("\n✓ Best practices demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunFundamentalsExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║        GO FUNDAMENTALS - COMPLETE GUIDE               ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	VarAndConstantsDemo()
	DataTypesDemo()
	TypeConversionDemo()
	FunctionsDemo()
	ControlFlowDemo()
	DeferDemo()
	PanicAndRecoverDemo()
	ScopeDemo()
	BestPracticesDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         FUNDAMENTALS MASTERY - YOU'RE READY!           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
