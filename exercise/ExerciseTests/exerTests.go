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

	// commented since it is a blocking function
	// NonArgsTest(15, exer.CreateServer)

	// commented because the data is too huge clogging the terminal
	// NonArgsTest(16, exer.GetUsers)

	NonArgsTest(17, exer.RunConcurrent)

	NonArgsTest(18, exer.HandleChannels)

	NonArgsTest(19, exer.HandleIncrements)

	NonArgsTest(19, exer.CounterIncrementHandler)

	NonArgsTest(20, exer.RunTimerAndTickerExamples)

	testInfo(21)
	fmt.Println(exer.CheckNegative(-25))

	testInfo(22)
	result, divErr := exer.Divide(22, 4)
	fmt.Println(result, divErr)

	testInfo(23)
	rect1 := exer.Rectangle{Height: 21, Width: 43}
	// passing rectangle as a shape since it implements all Shpe methods
	exer.DescribeShape(rect1)

	testInfo(24)
	fmt.Println(exer.SortSlice(exer.People))

	// inject dependency via the Fetcher struct which implements the UserData interface
	// 	 which is required by the Workermodel for Data fetching
	var worker1 = exer.NewWorker(exer.Fetcher{}, 303)
	NonArgsTest(25, worker1.DisplayWorkerInfo)

	NonArgsTest(26, exer.LoadEnv)

	// commented since it writes to file
	testInfo(27)
	// exer.Logger("Today is a good day", "simpleLogger.txt", os.Stdout)

	// commented since it creates a serve which is a blocking operation
	// NonArgsTest(28, exer.TasksServer)

	// for exercise 29: see the command commentted on code file '29unitTesting_test.go'

	// for exercise 30: see the command commented on code file '30benchmarkTest_test.go'

	// commented since its a blocking operation as it creates a 'http server'
	// NonArgsTest(32, exer.StartUserServer)

	// for exercise 33 run this test to spin up the go server
	// then use the '33.html' UI to upload a file to the server
	// NonArgsTest(33, exer.StartUploadHandlerServer)

	// 'NonArgsTest(34, exer.DataAPI)' is commented since it is a blocking operation as it starts a http server
	/*
		1. Uncomment the commented 'exer.DataAPI' line and comment the 'exer.GetData("http://localhost:3220/data")' then run the file
		2. after starting the server, comment the 'exer.DataAPI' line (to prevent restating the server), and uncomment the 'exer.GetData("http://localhost:3220/data")' line
		  then on a new terminal, run the file to get the data from the already running server without closing the first terminal
	*/
	// NonArgsTest(34, exer.DataAPI)
	// exer.GetData("http://localhost:3220/data")

	NonArgsTest(35, exer.WebScraper)
}
