/* Exercise Seven
7. Write a function that finds the maximum value in a slice of integers.
*/

package exercise

func GetMaxVal(nums []int) int {
	max := nums[0]
	for val := range nums {
		if val > max {
			max = val
		}
	}
	return max
}
