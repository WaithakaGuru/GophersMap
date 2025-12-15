/* Exercise Twenty one
21. Write a function that returns an error if a number is negative.
*/

package exercise

import "errors"

// implement the function to check and return and error if the number is negative
// handling errors using the 'errors' built-in go package
func CheckNegative[n int | float32 | float64](num n) error {
	if num < 0 {
		return errors.New("Failure: the number is negative")
	}
	return nil
}
