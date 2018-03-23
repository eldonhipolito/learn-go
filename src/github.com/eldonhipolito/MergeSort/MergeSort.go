package main


import "fmt"

type ArrayInt []int 


type Sortable interface {
	Sort(asc bool)
}

func MergeSort(arr []int, p,q,r int, asc bool) {
	if((r - p) == 0) {
	//fmt.Printf("Ret : P : %v - Q: %v - R: %v\n", p, q, r)
		return
	}

	//Divide until base is 1
//fmt.Printf("P : %v - Q: %v - R: %v\n", p, q, r)
	MergeSort(arr, p, (p + q) / 2 , q, asc)
	MergeSort(arr, q + 1, (r + q + 1)/ 2, r, asc)

	subArr := make([]int, (r + 1) - p)
	i,i2,sub := 0,0,0
	for ; (i + p) <= q && (q + 1 + i2) <= r; sub++ {
		//fmt.Printf("arr[i + p] : %v - arr[q + 1 + i2] : %v\n",arr[i + p], arr[q + 1 + i2])
		if(arr[i + p] > arr[q + 1 + i2] == asc) {
			subArr[sub] = arr[q + 1 + i2]
			i2++
		} else {
			subArr[sub] = arr[i + p]
			i++
		}
	}

	for ; (i + p) <= q; i++ {
		subArr[sub] = arr[i + p]
		sub++

	}
	for ; (q + 1 + i2) <= r; i2++ {
		subArr[sub] = arr[ q + 1 + i2]
		sub++

	}

	//fmt.Printf("Subarr[] : %v\n", subArr)	
	copy(arr[p:r + 1], subArr)
	

	
}

func (arr ArrayInt) Sort(asc bool) {
		MergeSort(arr, 0, len(arr) / 2 , len(arr) - 1, asc)

	
}


func main() {
	var a Sortable

	a = ArrayInt([]int{1,8,3,2,5})

	//ASC
	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(true)
	fmt.Printf("After sort : %v\n", a)

	a = ArrayInt([]int{6,4,3,2,1,5,6,3,11})

	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(false)
	fmt.Printf("After sort : %v\n", a)
}