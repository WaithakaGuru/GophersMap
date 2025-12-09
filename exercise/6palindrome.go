/* Exercise Six
6. Check if a given string is a palindrome.
*/

package exercise

func IsPalindrome(str string) bool {
	// use two pointer method
	i, j := 0, len(str)-1

	for i <= j {
		if !(str[i] == str[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
