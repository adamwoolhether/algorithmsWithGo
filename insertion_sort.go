package module02

import "sort"

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.

// I chose to use an anonymous function, instead of breaking the insert job
// into it's own function.
func InsertionSortInt(list []int) {
	var sortedList []int

	for _, listValue := range list {
		sortedList = func(sortedList []int, listValue int) []int {
			for i, sortValue := range sortedList {
				if listValue < sortValue {
					// sortedList[:i] + listValue + sotedList[i:]
					return append(sortedList[:i], append([]int{listValue}, sortedList[i:]...)...)
				}
			}
			return append(sortedList, listValue)
		}(sortedList, listValue)
	}

	for i, v := range sortedList {
		list[i] = v
	}
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
}
