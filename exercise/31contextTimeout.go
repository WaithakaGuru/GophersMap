/* Thirty one
31. Use `context.WithTimeout` to cancel a long-running operation.
*/

package exercise

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	count := 0
	fmt.Println(" ----Started long runnig operation :) ---- ")
	for range 10 {
		time.Sleep(time.Second)
		count++
		fmt.Println("--Current progess: COUNT -> ", count, "/9")
	}
	fmt.Println(" ---END: Finished long running operation ---")
}

func TimeOut(ctx context.Context) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	slowOperation(c)
}
