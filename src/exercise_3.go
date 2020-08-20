package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func printArray(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

func getRandomArray(size int) []int {
	randomIntArray := make([]int, size)
	min := math.MinInt32
	max := math.MaxInt32
	// seed := rand.NewSource(time.Now().UnixNano())
	seed := rand.NewSource(0)
	random := rand.New(seed)

	for i := 0; i < size; i++ {
		randomIntArray[i] = min + random.Intn(max-min)
	}

	return randomIntArray
}

func defaultSort(unsortedArray []int) {
	beforeSort := time.Now()
	sort.Ints(unsortedArray[:])
	// for _, value := range unsortedArray {
	// 	fmt.Println(value)
	// }
	// fmt.Println(runtime.NumGoroutine())
	sortDuration := time.Since(beforeSort)
	fmt.Println(sortDuration)
}

func partitionSort(unsortedArray []int) {
	length := len(unsortedArray)
	partition1Start := 0
	partition1End := length / 4
	partition2Start := partition1End
	partition2End := length / 2
	partition3Start := partition2End
	partition3End := int(math.Floor(float64(length) * 0.75))
	partition4Start := partition3End
	partition4End := length
	// fmt.Println(0, partition1End)
	// fmt.Println(partition2Start, partition2End)
	// fmt.Println(partition3Start, partition3End)
	// fmt.Println(partition4Start, partition4End)
	partition1 := unsortedArray[partition1Start:partition1End]
	partition2 := unsortedArray[partition2Start:partition2End]
	partition3 := unsortedArray[partition3Start:partition3End]
	partition4 := unsortedArray[partition4Start:partition4End]

	beforeSort := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(4)
	go sortArray(partition1, &wg)
	go sortArray(partition2, &wg)
	go sortArray(partition3, &wg)
	go sortArray(partition4, &wg)

	// fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	wg.Add(2)
	go mergeSortedArrays(unsortedArray, partition1, partition2, partition1Start, &wg)
	go mergeSortedArrays(unsortedArray, partition3, partition4, partition3Start, &wg)

	wg.Wait()
	newPartition1 := unsortedArray[partition1Start:partition2End]
	newPartition2 := unsortedArray[partition3Start:partition4End]

	wg.Add(1)
	mergeSortedArrays(unsortedArray, newPartition1, newPartition2, partition1Start, &wg)

	// printArray(0, unsortedArray)
	sortDuration := time.Since(beforeSort)
	fmt.Println(sortDuration)
}

func mergeSortedArrays(originalArray []int, array1 []int, array2 []int, originalArrayStart int, wg *sync.WaitGroup) {
	i := 0
	j := 0
	k := 0
	arrayOneLen := len(array1)
	arrayTwoLen := len(array2)
	totalLen := arrayOneLen + arrayTwoLen
	temp := make([]int, totalLen)

	for i < arrayOneLen && j < arrayTwoLen {
		if array1[i] < array2[j] {
			temp[k] = array1[i]
			i++
			k++
		} else {
			temp[k] = array2[j]
			j++
			k++
		}
	}

	for i < arrayOneLen {
		temp[k] = array1[i]
		k++
		i++
	}

	for j < arrayTwoLen {
		temp[k] = array2[j]
		k++
		j++
	}

	copy(originalArray[originalArrayStart:originalArrayStart+totalLen], temp)

	defer wg.Done()
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
			fmt.Println("Invalid sort method (default or partition)")
		}

		if out {
			fmt.Println("Is sorted", sort.IntsAreSorted(unsortedArray))
			// printArray(unsortedArray)
			break
		}
	}
}
