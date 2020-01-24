package main

/*
	It's a divide-and-conquer algorithm which means input array is divided into two arrays
	until there is nothing left to divide. Then all divided parts are merged.

	Best Case: O(nlogn)
	Worst Case: O(nlogn)

	Example Divide:

	[5 4]		[3 2 1]
	[5] [4]		[3] [2 1]
					[2] [1]

*/
func MergeSort(items []int) []int {

	length := len(items)

	if length <= 1 {
		return items
	}

	middle := int(length / 2)
	leftPart := items[:middle]
	rightPart := items[middle:]

	return merge(MergeSort(leftPart), MergeSort(rightPart))
}

func merge(left []int, right []int) []int {

	resultArray := make([]int, len(left)+len(right))

	indexOfResult := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			resultArray[indexOfResult] = left[0]
			left = left[1:]
		} else {
			resultArray[indexOfResult] = right[0]
			right = right[1:]
		}

		indexOfResult++
	}

	for j := 0; j < len(left); j++ {
		resultArray[indexOfResult] = left[j]
		indexOfResult++
	}

	for j := 0; j < len(right); j++ {
		resultArray[indexOfResult] = right[j]
		indexOfResult++
	}
	return resultArray
}
