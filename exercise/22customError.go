// to define a custom Error we just have to implement the Error() method of the error{} interface

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

func Divide[n int | float32 | float64](a, b n) n {
	if b == 0 {
		Fail("Division Error: Cannot divide by Zero")
		return 0
	}
	return a / b
}
