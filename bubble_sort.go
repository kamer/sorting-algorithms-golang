package main

/*
	It's the simplest sorting algorithm. It has a basic method:
		1. Iterate all elements.
		2. Compare two adjacent elements in every step. Swap them if needed.

	- Best Case: O(n).
	- Worst Case: O(n^2)

	This algorithm is not generally used because of its bad performance.

	Example Steps:

	       [5 4 3 2 1]
	(0, 1) [4 5 3 2 1]
	(1, 2) [4 3 5 2 1]
	(2, 3) [4 3 2 5 1]
	(3, 4) [4 3 2 1 5]
	(0, 1) [3 4 2 1 5]
	(1, 2) [3 2 4 1 5]
	(2, 3) [3 2 1 4 5]
	(0, 1) [2 3 1 4 5]
	(1, 2) [2 1 3 4 5]
	(0, 1) [1 2 3 4 5]


*/

func BubbleSort(items []int) []int {

	length := len(items)

	for i := 0; i < length; i++ {
		for j := 0; j < (length - i - 1); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}
