package main

/*
	This algorithm is one of the most used algorithms. It's a divide-and-conquer algorithm
	which means algorithm will be divided into two arrays and and be sorted individually.
	Sort is performed with these simple steps:
	1. A pivot element is selected from array.
	2. Array is divided into two arrays: elements lower than pivot and higher than pivot.
	3. Arrays are sorted with first two steps.

	Best Case: O(nlogn)
	Worst Case: O(n^2)
*/
func QuickSort(items []int) []int {

	length := len(items)

	return sort(items, 0, length-1)

}

func sort(elements []int, lowIndex int, highIndex int) []int {

	i := lowIndex
	j := highIndex
	pivot := elements[lowIndex+(highIndex-lowIndex)/2]

	for i <= j {

		for elements[i] < pivot {
			i++
		}

		for elements[j] > pivot {
			j--
		}

		if i <= j {
			elements[i], elements[j] = elements[j], elements[i]
			j--
			i++
		}
	}

	if lowIndex < j {
		sort(elements, lowIndex, j)
	}
	if highIndex > i {
		sort(elements, i, highIndex)
	}
	return elements
}
