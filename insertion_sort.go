package main

/*
	We actually use this algorithm in our daily life while playing card games. We pick a card and compare it
	by going backwards.

	Best Case: O(n)
	Worst Case:  O(n^2)

	Example Steps:

           [5 4 3 2 1]
	(1, 2) [4 5 3 2 1]
	(2, 3) [4 3 5 2 1]
	(1, 2) [3 4 5 2 1]
	(3, 4) [3 4 2 5 1]
	(2, 3) [3 2 4 5 1]
	(1, 2) [2 3 4 5 1]
	(4, 5) [2 3 4 1 5]
	(3, 4) [2 3 1 4 5]
	(2, 3) [2 1 3 4 5]
	(1, 2) [1 2 3 4 5]

*/
func InsertionSort(items []int) []int {

	length := len(items)

	for i := 1; i < length; i++ {
		j := i
		for j >= 1 && items[j] < items[j-1] {
			items[j], items[j-1] = items[j-1], items[j]
			j--
		}
	}
	return items
}
