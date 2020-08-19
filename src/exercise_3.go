package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func printArray(num int, array []int) {
	for _, value := range array {
		fmt.Println(num, value)
	}
}

func getRandomArray(size int) []int {
	randomIntArray := make([]int, size)
	min := math.MinInt32
	max := math.MaxInt32
	for i := 0; i < size; i++ {
		randomIntArray[i] = min + rand.Intn(max-min)
	}

	return randomIntArray
}

func defaultSort(unsortedArray []int) {
	beforeSort := time.Now()
	sort.Ints(unsortedArray[:])
	for _, value := range unsortedArray {
		fmt.Println(value)
	}
	sortDuration := time.Since(beforeSort)
	fmt.Println(sortDuration)
}

func partitionSort(unsortedArray []int) {
	length := len(unsortedArray)
	// fmt.Println(0, length/4)
	// fmt.Println(length/4, length/2)
	// fmt.Println(length/2, int(math.Floor(float64(length)*0.75)))
	// fmt.Println(int(math.Floor(float64(length)*0.75)), length-1)
	partition1 := unsortedArray[0 : length/4]
	partition2 := unsortedArray[length/4 : length/2]
	partition3 := unsortedArray[length/2 : int(math.Floor(float64(length)*0.75))]
	partition4 := unsortedArray[int(math.Floor(float64(length)*0.75)):length]

	beforeSort := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(4)
	go sortArray(partition1, &wg)
	go sortArray(partition2, &wg)
	go sortArray(partition3, &wg)
	go sortArray(partition4, &wg)

	wg.Wait()
	printArray(0, unsortedArray)

	sortDuration := time.Since(beforeSort)
	fmt.Println(sortDuration)
}

func sortArray(array []int, wg *sync.WaitGroup) {
	sort.Ints(array)
	defer wg.Done()
}
func main() {
	var arrayLength int
	var sortMethod string

	for {
		fmt.Println("Enter length of array")
		fmt.Scanln(&arrayLength)
		unsortedArray := getRandomArray(arrayLength)

		out := false
		fmt.Println("Enter Sort Method")
		fmt.Scanln(&sortMethod)
		switch sortMethod {
		case "default":
			defaultSort(unsortedArray)
			out = true
		case "partition":
			partitionSort(unsortedArray)
			out = true
		default:
			fmt.Println("Invalid sort method")
		}

		if out {
			break
		}
	}
}
