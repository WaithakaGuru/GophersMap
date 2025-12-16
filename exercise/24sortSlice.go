/* Exercise Twenty Four
24. Sort a slice of structs by a field (e.g., age).
*/

package exercise

type person struct {
	age  int
	name string
}

var People []person = []person{
	{age: 34, name: "Ann"},
	{age: 23, name: "Amos"},
	{age: 78, name: "mwangi"},
	{age: 45, name: "Brave"},
	{age: 12, name: "Kim"},
}

func SortSlice(items []person) []person {
	// strategy :
	// comapare current to next , swap if need be and increment current
	// repeat until no swap is done
	var swapCount = 1
	for swapCount > 0 {
		swapCount = 0
		for i := range items {
			if i < len(items)-1 && items[i+1].age < items[i].age {
				items[i+1], items[i] = items[i], items[i+1]
				swapCount++
			}
		}
	}
	return items
}
