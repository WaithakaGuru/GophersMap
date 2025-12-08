/* Exercise Forteen
14. Print all command-line arguments passed to your program.
*/

package exercise

import (
	"fmt"
	"os"
)

func PrintCLIArguments() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
