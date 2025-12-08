/*	Exercise Eight
8. Count the occurrences of each word in a string and store them in a map.
*/

package exercise

func CountOccurences(str string) map[string]int {
	var wordCountMap map[string]int = make(map[string]int, 0)

	for _, word := range str {
		if _, ok := wordCountMap[string(word)]; ok {
			wordCountMap[string(word)] += 1
		} else {
			wordCountMap[string(word)] = 1
		}
	}
	return wordCountMap
}
