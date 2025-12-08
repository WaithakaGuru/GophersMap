/* Exercise Two
2. Create a program that takes two numbers and an operator (+, -, *, /) and prints the result.
*/

package exercise

import "fmt"

type CustomError struct {
	errorMessage string
}

func (e CustomError) Error() string {
	return e.errorMessage
}

func Calculate(num1, num2 float64, op string) {
	result, err := calc(num1, num2, op)
	if err != nil {
		fmt.Println("An error occured: ", err.Error())
		return
	}

	fmt.Printf("The result of '%f %s %f' is %f ", num1, op, num2, result)
}

func calc(num1, num2 float64, op string) (float64, error) {
	var result float64 = 0
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			divError := &CustomError{"DivisionError: Cannot divide by zero"}
			return 0, divError
		}
		result = num1 / num2
	default:
		goto exitOnInvalidOP
	}

	return result, nil
exitOnInvalidOP:
	invalidOpError := &CustomError{"OperatorError: Invalid operator"}
	return 0, invalidOpError
}
