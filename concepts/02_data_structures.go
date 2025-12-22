/*
DATA STRUCTURES IN GO - Arrays, Slices, Maps, Structs
============================================================================

Go provides powerful data structures that are essential for any program.
Understanding these is crucial before moving to more advanced topics.

TOPICS COVERED:
- Arrays: Fixed-size collections
- Slices: Dynamic collections (very important!)
- Maps: Key-value stores
- Structs: Custom types
- Methods on types
*/

package concepts

import (
	"fmt"
)

// ============================================================================
// 1. ARRAYS - Fixed Size Collections
// ============================================================================

/*
ARRAYS:
- Fixed size at compile time
- Size is part of the type
- [5]int is different from [10]int
- Rarely used in practice (slices are preferred)
- Useful for fixed-size collections (like coordinates)
- Access: O(1), add/remove: Not efficient
*/

func ArraysDemo() {
	fmt.Println("\n========== ARRAYS ==========")

	// Array declaration with size
	var numbers [5]int
	fmt.Printf("Empty array: %v\n", numbers) // Zero values: [0 0 0 0 0]

	// Array with initialization
	scores := [4]int{85, 90, 78, 95}
	fmt.Printf("Scores: %v\n", scores)

	// Array with ... for length inference
	fruits := [...]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits (%d items): %v\n", len(fruits), fruits)

	// Accessing elements
	fmt.Printf("Second fruit: %s\n", fruits[1])

	// Modifying elements
	fruits[0] = "mango"
	fmt.Printf("Updated fruits: %v\n", fruits)

	// Iterating over arrays
	fmt.Println("Iterating with for range:")
	for i, fruit := range fruits {
		fmt.Printf("  [%d] = %s\n", i, fruit)
	}

	// Multi-dimensional arrays
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix:\n%v\n", matrix)

	fmt.Println("\n✓ Arrays demonstrated")
}

// ============================================================================
// 2. SLICES - Dynamic Collections (Most Important!)
// ============================================================================

/*
SLICES:
- Dynamic size (no size in type)
- Built on top of arrays (hidden backing array)
- Very efficient for most operations
- Pass by reference (changes affect original)
- Three components: pointer, length, capacity

KEY CONCEPTS:
- len(slice): Number of elements
- cap(slice): Capacity of backing array
- append(): Add elements (may reallocate)
- make(): Create slice with specific length/capacity
*/

func SlicesDemo() {
	fmt.Println("\n========== SLICES ==========")

	// 1. Creating slices
	fmt.Println("--- Creating Slices ---")

	// Slice literal
	colors := []string{"red", "green", "blue"}
	fmt.Printf("Colors: %v (len=%d, cap=%d)\n", colors, len(colors), cap(colors))

	// Using make - creates slice with specific length and capacity
	numbers := make([]int, 5, 10) // length=5, capacity=10
	fmt.Printf("Made slice: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

	// Empty slice
	var empty []string
	fmt.Printf("Empty slice: %v (len=%d, cap=%d)\n", empty, len(empty), cap(empty))

	// 2. Slice operations
	fmt.Println("\n--- Slice Operations ---")

	// Append - adds elements
	colors = append(colors, "yellow")
	colors = append(colors, "purple", "orange")
	fmt.Printf("After append: %v\n", colors)

	// Slice operations - slice of slice
	fmt.Println("\n--- Slice of Slice ---")
	fmt.Printf("colors[1:3]: %v\n", colors[1:3]) // Elements at index 1,2
	fmt.Printf("colors[2:]: %v\n", colors[2:])   // From index 2 to end
	fmt.Printf("colors[:3]: %v\n", colors[:3])   // From start to index 3
	fmt.Printf("colors[:]: %v\n", colors[:])     // Entire slice

	// 3. Modifying slices
	fmt.Println("\n--- Modifying Slices ---")
	data := []int{1, 2, 3, 4, 5}
	data[2] = 30 // Change element
	fmt.Printf("Modified: %v\n", data)

	// Copy slice
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)
	dataCopy[0] = 999
	fmt.Printf("Original: %v, Copy: %v\n", data, dataCopy)

	// 4. Capacity and growth
	fmt.Println("\n--- Capacity and Growth ---")
	slice := make([]int, 0, 5)
	fmt.Printf("Start: len=%d, cap=%d\n", len(slice), cap(slice))

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(slice), cap(slice))
	}

	// 5. Practical slice operations
	fmt.Println("\n--- Practical Operations ---")

	// Remove element from slice
	original := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	result := append(original[:indexToRemove], original[indexToRemove+1:]...)
	fmt.Printf("Original: %v, After removing index 2: %v\n", original, result)

	// Reverse slice
	toReverse := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(toReverse)-1; i < j; i, j = i+1, j-1 {
		toReverse[i], toReverse[j] = toReverse[j], toReverse[i]
	}
	fmt.Printf("Reversed: %v\n", toReverse)

	fmt.Println("\n✓ Slices demonstrated")
}

// ============================================================================
// 3. MAPS - Key-Value Stores
// ============================================================================

/*
MAPS:
- Hash table implementation
- Keys must be comparable (int, string, etc.)
- Values can be any type
- Unordered (iteration order is random)
- Pass by reference (like slices)
- Access: O(1) average case
- Keys must be unique
*/

func MapsDemo() {
	fmt.Println("\n========== MAPS ==========")

	// 1. Creating maps
	fmt.Println("--- Creating Maps ---")

	// Map literal
	person := map[string]string{
		"name": "Alice",
		"city": "Nairobi",
		"job":  "Engineer",
	}
	fmt.Printf("Person: %v\n", person)

	// Using make
	scores := make(map[string]int)
	fmt.Printf("Empty map: %v\n", scores)

	// 2. Adding and modifying
	fmt.Println("\n--- Adding and Modifying ---")

	scores["alice"] = 85
	scores["bob"] = 90
	scores["charlie"] = 78
	fmt.Printf("Scores: %v\n", scores)

	// Update existing key
	scores["alice"] = 95
	fmt.Printf("After update: %v\n", scores)

	// 3. Accessing values
	fmt.Println("\n--- Accessing Values ---")

	// Safe access - check if key exists
	value, exists := scores["bob"]
	if exists {
		fmt.Printf("Bob's score: %d\n", value)
	}

	// Accessing non-existent key returns zero value
	notFound := scores["david"]
	fmt.Printf("David's score: %d (not found returns zero)\n", notFound)

	// 4. Deleting entries
	fmt.Println("\n--- Deleting Entries ---")

	delete(scores, "charlie")
	fmt.Printf("After delete: %v\n", scores)

	// 5. Iterating over maps
	fmt.Println("\n--- Iterating Over Maps ---")

	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// 6. Map of slices
	fmt.Println("\n--- Map of Slices ---")

	classGrades := map[string][]int{
		"Class A": {85, 90, 88},
		"Class B": {92, 95, 89},
		"Class C": {78, 80, 82},
	}

	for class, grades := range classGrades {
		fmt.Printf("%s: %v\n", class, grades)
	}

	// 7. Counting occurrences - practical example
	fmt.Println("\n--- Counting Occurrences ---")

	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("Word counts: %v\n", wordCount)

	fmt.Println("\n✓ Maps demonstrated")
}

// ============================================================================
// 4. STRUCTS - Custom Types
// ============================================================================

/*
STRUCTS:
- Group related data together
- Define custom types
- Can have methods
- Exported fields: Capitalize first letter
- Unexported fields: lowercase first letter
- Embedded/nested structs for composition
- Go uses composition, not inheritance
*/

// Person struct - groups related data
type Person struct {
	Name  string
	Age   int
	Email string
	City  string
}

// Employee struct - embedding Person (composition)
type Employee struct {
	Person     // Embedded struct
	EmployeeID string
	Department string
	Salary     float64
}

// Address struct - for nested data
type Address struct {
	Street string
	City   string
	ZIP    string
}

// Company struct - with nested struct
type Company struct {
	Name      string
	Founded   int
	HQ        Address
	Employees int
}

func StructsDemo() {
	fmt.Println("\n========== STRUCTS ==========")

	// 1. Creating structs
	fmt.Println("--- Creating Structs ---")

	// Method 1: Field by field
	person1 := Person{
		Name:  "Alice",
		Age:   28,
		Email: "alice@example.com",
		City:  "Nairobi",
	}
	fmt.Printf("Person 1: %v\n", person1)

	// Method 2: Positional (not recommended - hard to read)
	person2 := Person{"Bob", 35, "bob@example.com", "Mombasa"}
	fmt.Printf("Person 2: %v\n", person2)

	// Method 3: Partial initialization (rest are zero values)
	person3 := Person{Name: "Charlie", Age: 42}
	fmt.Printf("Person 3: %v\n", person3)

	// 2. Accessing fields
	fmt.Println("\n--- Accessing Fields ---")

	fmt.Printf("Name: %s, Age: %d\n", person1.Name, person1.Age)

	// Modifying fields
	person1.Age = 29
	fmt.Printf("Updated age: %d\n", person1.Age)

	// 3. Pointers to structs
	fmt.Println("\n--- Pointers to Structs ---")

	var ptr *Person = &person1
	fmt.Printf("Via pointer: %s\n", ptr.Name)

	// Go automatically dereferences struct pointers
	ptr.Age = 30 // Same as (*ptr).Age = 30
	fmt.Printf("After modification via pointer: %d\n", person1.Age)

	// 4. Nested/Embedded structs
	fmt.Println("\n--- Embedded Structs ---")

	employee := Employee{
		Person:     Person{"Diana", 32, "diana@example.com", "Kampala"},
		EmployeeID: "EMP001",
		Department: "Engineering",
		Salary:     120000,
	}

	fmt.Printf("Employee: %v\n", employee)
	fmt.Printf("Name (embedded): %s\n", employee.Name)
	fmt.Printf("Department: %s\n", employee.Department)

	// 5. Nested structs
	fmt.Println("\n--- Nested Structs ---")

	company := Company{
		Name:    "TechCorp",
		Founded: 2010,
		HQ: Address{
			Street: "123 Tech Street",
			City:   "San Francisco",
			ZIP:    "94102",
		},
		Employees: 500,
	}

	fmt.Printf("Company: %s\n", company.Name)
	fmt.Printf("HQ: %s, %s\n", company.HQ.City, company.HQ.ZIP)

	// 6. Struct tags - metadata about fields
	fmt.Println("\n--- Struct Tags ---")

	type User struct {
		ID   int    `json:"id" db:"user_id"`
		Name string `json:"name" db:"user_name"`
	}

	user := User{ID: 1, Name: "Eve"}
	fmt.Printf("User: %v\n", user)

	fmt.Println("\n✓ Structs demonstrated")
}

// ============================================================================
// 5. METHODS - Functions with Receivers
// ============================================================================

/*
METHODS:
- Functions attached to types
- Receiver type can be value or pointer
- Pointer receiver: can modify the value
- Value receiver: works on a copy
- Enables object-like syntax
*/

// Method on Person struct - value receiver (doesn't modify original)
func (p Person) Greeting() string {
	return fmt.Sprintf("Hello, I'm %s from %s", p.Name, p.City)
}

// Method on Person struct - pointer receiver (can modify)
func (p *Person) HaveBirthday() {
	p.Age++
}

// Method on slice
type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, num := range s {
		total += num
	}
	return total
}

// Method on custom type
type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func MethodsDemo() {
	fmt.Println("\n========== METHODS ==========")

	person := Person{"Frank", 30, "frank@example.com", "Dar es Salaam"}

	// Call method
	fmt.Println(person.Greeting())

	// Pointer receiver method
	fmt.Printf("Age before birthday: %d\n", person.Age)
	person.HaveBirthday()
	fmt.Printf("Age after birthday: %d\n", person.Age)

	// Method on slice
	numbers := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("Sum of %v = %d\n", numbers, numbers.Sum())

	// Method on custom type
	temp := Temperature(25)
	fmt.Printf("Temperature: %.1f°C = %.1f°F\n", temp.Celsius(), temp.Fahrenheit())

	fmt.Println("\n✓ Methods demonstrated")
}

// ============================================================================
// 6. INTERFACES - Contracts and Abstraction
// ============================================================================

/*
INTERFACES:
- Define method sets
- Types satisfy interfaces implicitly (duck typing)
- Enable polymorphism
- empty interface{} matches any type
- Very powerful for design
*/

// Shape interface - defines what shapes must do
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// Function that works with any Shape
func PrintShapeInfo(s Shape, name string) {
	fmt.Printf("%s: Area=%.2f, Perimeter=%.2f\n", name, s.Area(), s.Perimeter())
}

func InterfacesDemo() {
	fmt.Println("\n========== INTERFACES ==========")

	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	PrintShapeInfo(rect, "Rectangle")
	PrintShapeInfo(circle, "Circle")

	// Polymorphism - store different types in interface
	shapes := []Shape{rect, circle}
	fmt.Println("\nAll shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: Area=%.2f\n", i+1, shape.Area())
	}

	fmt.Println("\n✓ Interfaces demonstrated")
}

// ============================================================================
// 7. PRACTICAL EXAMPLE: Todo List Manager
// ============================================================================

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TodoList struct {
	items  []Todo
	nextID int
}

// Add a todo
func (tl *TodoList) Add(title string) {
	tl.nextID++
	tl.items = append(tl.items, Todo{
		ID:        tl.nextID,
		Title:     title,
		Completed: false,
	})
}

// Mark as complete
func (tl *TodoList) Complete(id int) error {
	for i := range tl.items {
		if tl.items[i].ID == id {
			tl.items[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("todo not found")
}

// Get all todos
func (tl *TodoList) GetAll() []Todo {
	return tl.items
}

// Get active todos
func (tl *TodoList) GetActive() []Todo {
	var active []Todo
	for _, todo := range tl.items {
		if !todo.Completed {
			active = append(active, todo)
		}
	}
	return active
}

func PracticalExampleDemo() {
	fmt.Println("\n========== PRACTICAL EXAMPLE: Todo List ==========")

	todoList := TodoList{}

	// Add todos
	todoList.Add("Learn Go")
	todoList.Add("Build a project")
	todoList.Add("Read documentation")

	// Display all
	fmt.Println("All todos:")
	for _, todo := range todoList.GetAll() {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s ID:%d - %s\n", status, todo.ID, todo.Title)
	}

	// Complete one
	todoList.Complete(1)

	// Display active
	fmt.Println("\nActive todos:")
	for _, todo := range todoList.GetActive() {
		fmt.Printf("- %s\n", todo.Title)
	}

	fmt.Println("\n✓ Practical example demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunDataStructuresExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║     DATA STRUCTURES IN GO - COMPLETE GUIDE             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	ArraysDemo()
	SlicesDemo()
	MapsDemo()
	StructsDemo()
	MethodsDemo()
	InterfacesDemo()
	PracticalExampleDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    DATA STRUCTURES MASTERY - BUILD ANYTHING!           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
