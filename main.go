package main

import "fmt"

func main() {

	elements := []int{4, 10, 3, 2, 1, 11, 7, 2, 9, 8}

	fmt.Println("Choose an algorithm to sort hard-coded array.")
	fmt.Printf("1.Bubble Sort\n2.Insertion Sort\n3.Merge Sort\n4.Quick Sort\n")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print(BubbleSort(elements))
	case 2:
		fmt.Print(InsertionSort(elements))
	case 3:
		fmt.Print(MergeSort(elements))
	case 4:
		fmt.Print(QuickSort(elements))
	default:
		fmt.Print("Bad choice!")

	}

}
