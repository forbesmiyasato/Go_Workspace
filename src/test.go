package main

import (
	"fmt"
	"math"
	"math/rand"
)

func getRandomArray(size int) []int {
	randomIntArray := make([]int, size)
	min := math.MinInt32
	max := math.MaxInt32
	for i := 0; i < size; i++ {
		randomIntArray[i] = min + rand.Intn(max-min)
	}

	return randomIntArray
}

func main() {
	var arraySize int
	fmt.Println("Enter length of array")
	fmt.Scanln(&arraySize)
	unsortedArray := getRandomArray(arraySize)
	for _, value := range unsortedArray {
		fmt.Println(value)
	}
}
