package main


import "fmt"

type Sortable interface {
	Sort(asc bool)
}

type ArrayInt []int {
}


func (arrInt *ArrayInt) Sort(asc bool) {
	if asc {
		for i := 1; i < len(arrInt); i++ {
			cElem  := arrInt[i]
			var p int
			for p = i - 1; p > -1 && cElem < arrInt[p]; p--{
				arrInt[p + 1] = arrInt[p]
			}

			arrInt[p + 1] = cElem 
		}
	} else {
		for i := 1; i < len(arrInt); i++ {
			cElem  := arrInt[i]
			var p int
			for p = i - 1; p > -1 && cElem > arrInt[p]; p--{
				arrInt[p + 1] = arrInt[p]
			}

			arrInt[p + 1] = cElem 
		}
	}
}

func main() {
	var a Sortable

	a = &ArrayInt{[]int{1,8,3,2,5}}

	//ASC
	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(true)
	fmt.Printf("After sort : %v\n", a)

	a = &ArrayInt{[]int{6,4,3,2,1,5,6,3,11}}

	fmt.Printf("Prior to sort : %v\n", a)
	a.Sort(false)
	fmt.Printf("After sort : %v\n", a)




}