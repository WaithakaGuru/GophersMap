/* Exercise Three
3. Write a function that checks if a number is odd or even.
*/

package exercise

import "fmt"

func IsOddOrEven(n int) {
	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
