/*  Exercise Eighteen
18. Use a channel to send and receive data between goroutines.
*/

// Span a go routine that will send a done message when it is done
// another listening routines awaits the done message to print a message

package exercise

import (
	"fmt"
	"time"
)

func someHeavyCalculation() {
	sum := 0

	for i := range 1000 {
		sum += i
		time.Sleep(2 * time.Millisecond)
	}
}

func printCalculationStatus() {
	fmt.Println("Done with calculation")
}

func HandleChannels() {
	calcDoneChannel := make(chan bool)
	now := time.Now()
	go func() {
		someHeavyCalculation()
		printCalculationStatus()
		calcDoneChannel <- true
	}()

	<-calcDoneChannel

	defer close(calcDoneChannel)
	fmt.Println("Calculation done and print func called in", time.Since(now))
}
