package main

import (
	"fmt"
	"time"
)

func insertionSort(A []int) {
	for j:=1; j<len(A); j++ {
		key := A[j]
		i := j - 1
		for ((i >= 0) && (A[i] > key)) {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	fmt.Print("\nSorted Array Is: ", A)
}

func main() {

	fmt.Print("\nHow many numbers are to be sorted? ")
	var n int
	fmt.Scanf("%d", &n)

	//A stores a slice of length n and type int
	A := make([]int, n)

	fmt.Print("Enter keys: ")

	for i:=0; i<n; i++ {
		_, err := fmt.Scan(&A[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	start := time.Now()
	insertionSort(A)
	elapsed := time.Since(start)

	fmt.Printf("\nTotal number of statements executed are => 7\n")
	fmt.Println("Total time taken in nanoseconds is => ", (int64(elapsed/time.Nanosecond)))
}

