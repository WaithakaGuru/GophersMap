/* Exercise Twenty nine
29. Write unit tests for a function (e.g., palindrome checker).
*/

package exercise_test

import (
	"testing"

	exer "github.com/WaithakaGuru/gophersmap/exercise"
)

// to run the test :
// cd exercise
//
//	go test -v
func TestPalindrome(t *testing.T) {
	got := exer.IsPalindrome("racecar")
	expected := true

	if got != expected {
		t.Errorf("test failed: expected %v but got %v", expected, got)
	}
}
