/* Exercise Five
5. Write a function that reverses a string.
*/

package exercise

func ReverseString(str string) string {
	// two pointer approach
	runeStr := []rune(str)
	i, j := 0, len(runeStr)-1

	for i <= j {
		temp := runeStr[i]
		runeStr[i] = runeStr[j]
		runeStr[j] = temp
		i++
		j--
	}
	return string(runeStr)
}
