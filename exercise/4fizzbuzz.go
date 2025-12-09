/*
	Exercise Four

4. Print numbers from 1 to 100. For multiples of 3 print "Fizz", for multiples of 5 print "Buzz", and for both print "FizzBuzz".
*/
package exercise

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
