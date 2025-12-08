/*	Exercise Nine
9. Remove duplicate elements from a slice.
*/

package exercise

func RemoveDuplicates[T string | int](items []T) []T {
	cleanElements := make(map[T]int, 0)

	for _, item := range items {
		if _, ok := cleanElements[item]; !ok {
			cleanElements[item] = 1
		}
	}
	var uniqueElements = []T{}
	for key := range cleanElements {
		uniqueElements = append(uniqueElements, key)
	}
	return uniqueElements
}
