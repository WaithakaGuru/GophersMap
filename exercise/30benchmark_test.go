/* Exercise Thirty
30. Benchmark a function using Go's testing package.
*/

package exercise_test

import (
	tst "testing"
)

// ----- bechmarking run commands---
// cd exercise
// go test -bench

// func one with time conplexity of O(n^2) - using recursion
func FibonnaciA(n int) int {
	if n < 1 {
		return n
	}
	return FibonnaciA(n-1) + FibonnaciA(n-2)
}

// func with time complexity of O(n) - using dynamic programming
func fibonacciB(n int) int {
	a, b := 0, 1
	for range n - 1 {
		A := a
		a, b = b, A+b
	}
	return b
}

// benchmarking measures a function's performance in terms of memory and time not correctness

func BenchmarkFibbonacciA(b *tst.B) {
	// to reset timer so as not to count time used for setup
	b.ResetTimer()

	// to report memory allocations
	b.ReportAllocs()
	//replace 'for i := 0; i < b.N;' i++  with 'for b.loop() {}'
	for b.Loop() {
		// function to test goes here
		FibonnaciA(10)
	}
}

func BenchmarkFibbonacciB(b *tst.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		fibonacciB(10)
	}
}
