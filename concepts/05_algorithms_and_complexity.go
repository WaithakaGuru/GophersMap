/*
SORTING AND SEARCHING ALGORITHMS IN GO
============================================================================

Understanding algorithms is fundamental for writing efficient code.
Go's standard library provides excellent sorting capabilities,
but understanding how they work is crucial.

TOPICS COVERED:
- Time and space complexity (Big O)
- Sorting algorithms (bubble, quick, merge)
- Searching algorithms (linear, binary)
- Practical examples and comparisons
- When to use which algorithm
*/

package concepts

import (
	"fmt"
	"sort"
	"time"
)

// ============================================================================
// 1. UNDERSTANDING BIG O NOTATION
// ============================================================================

/*
BIG O NOTATION - How code scales with input size

Common Complexities (best to worst):
O(1)      - Constant: 1 operation regardless of input size (array access)
O(log n)  - Logarithmic: halves with each iteration (binary search)
O(n)      - Linear: operations grow with input (linear search)
O(n log n)- Linearithmic: n × log(n) (good sorting: quicksort, mergesort)
O(n²)     - Quadratic: nested loops (bubble sort, insertion sort)
O(n³)     - Cubic: triple nested loops (rare in practice)
O(2ⁿ)     - Exponential: doubles with each input (very bad)
O(n!)     - Factorial: extremely bad (very rare)

SPACE COMPLEXITY:
O(1) - Constant space (in-place algorithms)
O(n) - Linear space (need array for result)
O(log n) - Logarithmic space (recursion depth)
*/

func ComplexityDemo() {
	fmt.Println("\n========== BIG O COMPLEXITY ANALYSIS ==========\n")

	fmt.Println("Time Complexity Examples:")
	fmt.Println("  O(1) - Array access: arr[5]")
	fmt.Println("  O(log n) - Binary search: find element in sorted array")
	fmt.Println("  O(n) - Linear search: find element by scanning all")
	fmt.Println("  O(n log n) - Quicksort, mergesort: efficient general sorting")
	fmt.Println("  O(n²) - Bubble sort, selection: simple but slow for large arrays")
	fmt.Println("  O(2ⁿ) - Brute force: avoid at all costs")

	fmt.Println("\nPractical Implications:")
	sizes := []int{10, 100, 1000, 10000, 100000}
	for _, n := range sizes {
		fmt.Printf("  n=%d: O(n)=%d, O(n log n)≈%d, O(n²)=%d\n",
			n, n, n*7, n*n)
	}

	fmt.Println("\n✓ Complexity analysis demonstrated")
}

// ============================================================================
// 2. SORTING ALGORITHMS
// ============================================================================

// Bubble Sort - Simple but slow O(n²)
func BubbleSort(arr []int) []int {
	n := len(arr)
	// Create copy to not modify original
	result := make([]int, n)
	copy(result, arr)

	// Outer loop: number of passes
	for i := 0; i < n-1; i++ {
		// Inner loop: compare adjacent elements
		for j := 0; j < n-i-1; j++ {
			// If current is greater than next, swap
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// Selection Sort - Simple O(n²)
func SelectionSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		// Find minimum in remaining array
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		// Swap with current position
		result[i], result[minIdx] = result[minIdx], result[i]
	}
	return result
}

// Quick Sort - Efficient O(n log n) average
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Choose pivot (middle element)
	pivot := arr[len(arr)/2]
	var left, middle, right []int

	// Partition array into 3 parts
	for _, x := range arr {
		switch {
		case x < pivot:
			left = append(left, x)
		case x == pivot:
			middle = append(middle, x)
		case x > pivot:
			right = append(right, x)
		}
	}

	// Recursively sort and combine
	result := append(QuickSort(left), middle...)
	result = append(result, QuickSort(right)...)
	return result
}

// Merge Sort - Guaranteed O(n log n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge
	return merge(left, right)
}

// Merge two sorted arrays
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	// Copy remaining elements
	copy(result[i+j:], left[i:])
	copy(result[i+j:], right[j:])

	return result
}

func SortingAlgorithmsDemo() {
	fmt.Println("\n========== SORTING ALGORITHMS ==========\n")

	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	fmt.Printf("Original: %v\n\n", data)

	// Bubble Sort
	fmt.Println("--- Bubble Sort (O(n²)) ---")
	bubbleSorted := BubbleSort(data)
	fmt.Printf("Result: %v\n", bubbleSorted)

	// Selection Sort
	fmt.Println("\n--- Selection Sort (O(n²)) ---")
	selectSorted := SelectionSort(data)
	fmt.Printf("Result: %v\n", selectSorted)

	// Quick Sort
	fmt.Println("\n--- Quick Sort (O(n log n)) ---")
	quickSorted := QuickSort(data)
	fmt.Printf("Result: %v\n", quickSorted)

	// Merge Sort
	fmt.Println("\n--- Merge Sort (O(n log n)) ---")
	mergeSorted := MergeSort(data)
	fmt.Printf("Result: %v\n", mergeSorted)

	// Built-in sort
	fmt.Println("\n--- Go's Built-in Sort (Highly optimized) ---")
	builtinData := make([]int, len(data))
	copy(builtinData, data)
	sort.Ints(builtinData)
	fmt.Printf("Result: %v\n", builtinData)

	fmt.Println("\n✓ Sorting algorithms demonstrated")
}

// ============================================================================
// 3. SEARCHING ALGORITHMS
// ============================================================================

// Linear Search - Simple O(n)
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // Found at index i
		}
	}
	return -1 // Not found
}

// Binary Search - Fast but requires sorted array O(log n)
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid // Found
		} else if arr[mid] < target {
			left = mid + 1 // Search right half
		} else {
			right = mid - 1 // Search left half
		}
	}

	return -1 // Not found
}

// Recursive binary search
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
}

func SearchingAlgorithmsDemo() {
	fmt.Println("\n========== SEARCHING ALGORITHMS ==========\n")

	unsorted := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	sorted := []int{11, 12, 22, 25, 34, 45, 50, 64, 88, 90}

	target := 45

	// Linear Search
	fmt.Println("--- Linear Search (O(n)) ---")
	fmt.Printf("Searching for %d in unsorted array\n", target)
	idx := LinearSearch(unsorted, target)
	if idx != -1 {
		fmt.Printf("Found at index: %d\n", idx)
	} else {
		fmt.Println("Not found")
	}

	// Binary Search
	fmt.Println("\n--- Binary Search (O(log n)) ---")
	fmt.Printf("Searching for %d in sorted array\n", target)
	idx = BinarySearch(sorted, target)
	if idx != -1 {
		fmt.Printf("Found at index: %d\n", idx)
	} else {
		fmt.Println("Not found")
	}

	// Recursive Binary Search
	fmt.Println("\n--- Recursive Binary Search (O(log n)) ---")
	idx = BinarySearchRecursive(sorted, target, 0, len(sorted)-1)
	fmt.Printf("Found at index: %d\n", idx)

	// Compare performance
	fmt.Println("\n--- Performance Comparison ---")
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = i
	}

	searchTarget := 9999

	// Linear search
	start := time.Now()
	LinearSearch(largeArray, searchTarget)
	linearTime := time.Since(start)

	// Binary search
	start = time.Now()
	BinarySearch(largeArray, searchTarget)
	binaryTime := time.Since(start)

	fmt.Printf("Array size: %d\n", len(largeArray))
	fmt.Printf("Linear search time: %v\n", linearTime)
	fmt.Printf("Binary search time: %v\n", binaryTime)
	fmt.Printf("Binary search is ~%.0fx faster\n",
		float64(linearTime)/float64(binaryTime))

	fmt.Println("\n✓ Searching algorithms demonstrated")
}

// ============================================================================
// 4. CUSTOM SORTING WITH SORT.INTERFACE
// ============================================================================

// Person struct for custom sorting
type PersonAlgo struct {
	Name string
	Age  int
}

// PersonSlice allows us to sort Person by age
type PersonSlice []PersonAlgo

// Implement sort.Interface
func (p PersonSlice) Len() int {
	return len(p)
}

func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func CustomSortingDemo() {
	fmt.Println("\n========== CUSTOM SORTING ==========\n")

	people := PersonSlice{
		{Name: "Alice", Age: 28},
		{Name: "Bob", Age: 35},
		{Name: "Charlie", Age: 22},
		{Name: "Diana", Age: 30},
	}

	fmt.Println("Before sorting:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// Sort using custom less function
	sort.Sort(people)

	fmt.Println("\nAfter sorting by age:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// Sort by name using Slice function (Go 1.8+)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("\nAfter sorting by name:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	fmt.Println("\n✓ Custom sorting demonstrated")
}

// ============================================================================
// 5. PRACTICAL EXAMPLE: Finding k Largest Elements
// ============================================================================

func FindKLargest(arr []int, k int) []int {
	if k > len(arr) {
		k = len(arr)
	}

	// Sort in descending order
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j] // Descending
	})

	return sorted[:k]
}

func PracticalAlgorithmDemo() {
	fmt.Println("\n========== PRACTICAL EXAMPLE: K Largest Elements ==========\n")

	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	k := 3

	fmt.Printf("Array: %v\n", data)
	fmt.Printf("Find %d largest elements\n\n", k)

	largest := FindKLargest(data, k)
	fmt.Printf("Result: %v\n", largest)

	fmt.Println("\n✓ Practical algorithm demonstrated")
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

func RunAlgorithmExamples() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║    ALGORITHMS & COMPLEXITY - COMPLETE GUIDE            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	ComplexityDemo()
	SortingAlgorithmsDemo()
	SearchingAlgorithmsDemo()
	CustomSortingDemo()
	PracticalAlgorithmDemo()

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║       ALGORITHMS MASTERY - WRITE EFFICIENT CODE!       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")
}
