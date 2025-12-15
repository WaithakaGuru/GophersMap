/* Exercise Twenty
20. Demonstrate the use of `time.Timer` and `time.Ticker`.
*/

package exercise

import "time"


func TimerAndTicker() {

	ticker1 := time.NewTicker(time.Second)

	timer1 := time.NewTimer(30 * time.Second)

	ticker1.C
}