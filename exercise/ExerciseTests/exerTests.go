/*
	Testing Area
*/

package main

import (
	"fmt"

	exer "github.com/WaithakaGuru/gophersmap/exercise"
)

// num is the exercise number for which you are testing
// func Test[T any | []any](num int, test func(T), arg T) func(T) {
// 	return func(arg T) {
// 		fmt.Println("\nTest results for Exercise ", num)
// 		test(arg)
// 		fmt.Println()
// 	}
// }

func testInfo(num int) {
	fmt.Printf("\n\nTest results for Exercise %v \n", num)
}

func NonArgsTest(num int, test func()) {
	fmt.Println("Test results for Exercise ", num)
	test()
	fmt.Println()
}

func main() {
	// test exercise 1
	NonArgsTest(1, exer.SayHello)

	testInfo(2)
	exer.Calculate(23.445, 456, "+")

	testInfo(3)
	exer.IsOddOrEven(34)

	// test exercise 4
	NonArgsTest(4, exer.FizzBuzz)

	testInfo(5)
	fmt.Println("Reversed string:", exer.ReverseString("Golang"))

	testInfo(6)
	fmt.Println("is palindrome:", exer.IsPalindrome("racecar"))
	fmt.Println("is palindrome:", exer.IsPalindrome("mother"))

	testInfo(7)
	fmt.Println(exer.GetMaxVal(23, 56, 67, 32, 88, 90, 43, 41))

	testInfo(8)
	fmt.Println(exer.CountOccurences("Waithaka is a Golang coder, developer and a system architect"))

	testInfo(9)
	fmt.Println(exer.RemoveDuplicates("waith", "kim", "munga", "lucy", "kim", "Amo", "waith"))
	fmt.Println(exer.RemoveDuplicates(23, 4, 5, 5, 23, -1, 0, 45))

	testInfo(10)
	exer.CreatePrintPerson("Waithaka", "a@b.c", 23)

	testInfo(11)
	var testUser exer.User = exer.User{Name: "Waithaka", Role: "Coder"}
	info, err := exer.Unmarshal(exer.Marshal(testUser))

	if err == nil {
		fmt.Println(info)
	} else {
		fmt.Println(err)
	}

	testInfo(12)
	exer.ReadFile("test.txt")

	testInfo(13)
	exer.WriteToNewFile("test.txt", "New information")
	exer.ReadFile("test.txt")

	testInfo(14)
	exer.PrintCLIArguments()

	testInfo(15)
	// commented since it is a blocking function
	// exer.CreateServer()

	testInfo(16)
	// the data is too huge clogging the terminal
	// exer.GetUsers()

	testInfo(17)
	exer.RunConcurrent()

	testInfo(18)
	exer.HandleChannels()

	testInfo(19)
	exer.HandleIncrements()
}
