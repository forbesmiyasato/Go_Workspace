package main

import (
	"sync"
)

//Merge ....
func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

// MergeSortMulti sorts the slice s using Merge Sort Algorithm
func MergeSortMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	go func() {
		l = MergeSortMulti(s[:n])
		wg.Done()
	}()

	go func() {
		r = MergeSortMulti(s[n:])
		wg.Done()
	}()

	wg.Wait()
	return Merge(l, r)
}
