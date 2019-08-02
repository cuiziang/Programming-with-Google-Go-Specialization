package main

import (
	"fmt"
	"strconv"
)

func swap(a *[]int, i int, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func QuickSort(a *[]int, left int, right int) {
	if right <= left {
		return
	}

	idx := left

	pivotIndex := left

	swap(a, pivotIndex, right)

	for i := left; i < right; i++ {
		if (*a)[i] < (*a)[right] {
			swap(a, i, idx)
			idx++
		}
	}

	swap(a, idx, right)
	QuickSort(a, left, idx)
	QuickSort(a, idx+1, right)
	return
}

func ConcurrentQuickSort(a *[]int, left int, right int, chanSend chan int) {
	if (right - left) < 1000 {
		QuickSort(a, left, right)
		chanSend <- 0
		return
	}

	if right <= left {
		chanSend <- 0
		return
	}

	idx := left

	pivotIndex := left

	swap(a, pivotIndex, right)

	for i := left; i < right; i++ {
		if (*a)[i] < (*a)[right] {
			swap(a, i, idx)
			idx++
		}
	}

	swap(a, idx, right)

	chanReceive := make(chan int, 4)

	go ConcurrentQuickSort(a, left, idx, chanReceive)
	go ConcurrentQuickSort(a, idx+1, right, chanReceive)
	<-chanReceive
	<-chanReceive
	chanSend <- 0
	return
}

func main() {

	var testArr1 []int
	for {
		fmt.Printf("Enter a number a[%d], x to stop:\n")
		var input string
		_, _ = fmt.Scanln(&input)
		if input == "x" {
			break
		}
		num, _ := strconv.Atoi(input)
		testArr1 = append(testArr1, num)
	}

	chanWait := make(chan int)
	go ConcurrentQuickSort(&testArr1, 0, len(testArr1)-1, chanWait)
	<-chanWait
	fmt.Print(testArr1)
}
