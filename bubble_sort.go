package module02

import (
	"sort"
	"strings"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
// Bubble sort is most optimized for already, or nearly sorted lists.
func BubbleSortInt(list []int) {
	// Unoptimized:
	/*	for i := 0; i < len(list); i++ {
		for p := 0; p < len(list)-1; p++ {
			if list[p] > list[p+1] {
				list[p], list[p+1] = list[p+1], list[p]
			}
		}
	}*/

	// Optimized (improves it a bit):
	for i := 0; i < len(list); i++ {
		swapped := false

		for p := 0; p < len(list)-1-i; p++ {

			if list[p] > list[p+1] {
				list[p], list[p+1] = list[p+1], list[p]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	/*	for i := 0; i < len(list); i++ {
		for p := 0; p< len(list)-1; p++ {
			if strings.ToLower(list[p]) > strings.ToLower(list[p+1]) {
				list[p], list[p+1] = list [p+1], list[p]
			}
		}
	}*/

	// Optimized:
	for i := 0; i < len(list); i++ {
		swapped := false

		for p := 0; p < len(list)-1-i; p++ {

			if strings.ToLower(list[p]) > strings.ToLower(list[p+1]) {
				list[p], list[p+1] = list[p+1], list[p]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	for i := 0; i < len(people); i++ {
		for p := 0; p < len(people)-1; p++ {
			if people[p].Age > people[p+1].Age {
				people[p], people[p+1] = people[p+1], people[p]
			}
			if people[p].Age == people[p+1].Age {
				if people[p].LastName > people[p+1].LastName {
					people[p], people[p+1] = people[p+1], people[p]
				}
				if people[p].LastName == people[p+1].LastName {
					if people[p].FirstName > people[p+1].FirstName {
						people[p], people[p+1] = people[p+1], people[p]
					}
				}

			}

		}

	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		swapped := false

		for p := 0; p < list.Len()-1-i; p++ {
			if list.Less(p+1, p) {
				list.Swap(p, p+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// Some scratch code i used for clarity
// func main() {
// 	people := []Person{
// 		{FirstName: "Adam", LastName: "Woo", Age: 55},
// 		{FirstName: "Ron", LastName: "Woo", Age: 28},
// 		{FirstName: "Joe", LastName: "Smith", Age: 25},
// 		{FirstName: "Joe", LastName: "Michaels", Age: 45},
// 		{FirstName: "Ron", LastName: "Smith", Age: 28},
// 	}
//
// 	BubbleSortPerson(people)
// 	fmt.Println(people)
// }
