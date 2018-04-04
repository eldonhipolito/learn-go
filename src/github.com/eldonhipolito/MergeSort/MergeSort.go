package main

import "fmt"

type ArrayInt []int

type Sortable interface {
	Sort(asc bool)
}

func MergeSort(arr, temp []int, p, r int, asc bool) {
	if p == r {
		return
	}

	//Divide until base is 1
	q := (r + p) / 2
	MergeSort(arr, temp, p, q, asc)
	MergeSort(arr, temp, q+1, r, asc)

	subArr := temp[:r+1]
	i, i2, sub := 0, 0, 0
	for ; (i+p) <= q && (q+1+i2) <= r; sub++ {
		if arr[i+p] > arr[q+1+i2] == asc {
			subArr[sub] = arr[q+1+i2]
			i2++
		} else {
			subArr[sub] = arr[i+p]
			i++
		}
	}

	for ; (i + p) <= q; i++ {
		subArr[sub] = arr[i+p]
		sub++

	}
	for ; (q + 1 + i2) <= r; i2++ {
		subArr[sub] = arr[q+1+i2]
		sub++

	}

	copy(arr[p:r+1], subArr)

}

func (arr ArrayInt) Sort(asc bool) {
	MergeSort(arr, make([]int, len(arr)), 0, len(arr)-1, asc)

}

func main() {
	var a Sortable

	a = ArrayInt([]int{1, 8, 3, 2, 5})

	//ASC
	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(true)
	fmt.Printf("After sort : %v\n", a)

	a = ArrayInt([]int{6, 4, 3, 2, 1, 5, 6, 3, 11, 711})

	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(false)
	fmt.Printf("After sort : %v\n", a)
}
