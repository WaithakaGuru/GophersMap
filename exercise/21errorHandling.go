/* Exercise Twenty one
21. Write a function that returns an error if a number is negative.
*/

package exercise

// am going to create my own error function

type Failure struct {
	info string
}

func (f Failure) Error() string {
	return f.info
}

func Fail(info string) *Failure {
	return &Failure{info: info}
}

// implement the function to check and return and error if the number is negative

func CheckNegative[n int | float32 | float64](num n) error {
	if num < 0 {
		return Fail("Failure: the number is negative")
	}
	return nil
}
