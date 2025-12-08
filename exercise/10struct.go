/* Exercise Ten
10. Define a `Person` struct with name, age, and email. Create and print a few instances.
*/

package exercise

import "fmt"

type Person struct {
	name  string
	email string
	age   uint
}

func CreatePrintPerson(name, email string, age uint) {
	newPerson := Person{
		name, email, age,
	}

	fmt.Println(newPerson)
}
