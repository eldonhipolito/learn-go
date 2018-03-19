package main

import "fmt"



type Sortable interface {
	Sort(asc bool)
}

type ArrayInt []int


func (arrInt ArrayInt) Sort(asc bool) {
	
	if asc {
		
		for i := 0; i < len(arrInt) - 1; i++ {

			//Search smallest starting with i	
			smallest := i
			for j := i + 1; j < len(arrInt); j++ {

				if(arrInt[smallest] > arrInt[j]) {
					smallest = j
				}
			}
			if(smallest != i) {
				arrInt[i] = arrInt[smallest] + arrInt[i]
				arrInt[smallest] = arrInt[i] - arrInt[smallest]
				arrInt[i] = arrInt[i] - arrInt[smallest]
			}
		}		
	
	} else {
			for i := 0; i < len(arrInt) - 1; i++ {

			//Search biggest starting with i
			biggest := i
			for j := i + 1; j < len(arrInt); j++ {
				if(arrInt[biggest] < arrInt[j]) {
					biggest = j
				}
			}
			//Swap values
			if(biggest != i) {
				arrInt[i] = arrInt[biggest] + arrInt[i]
				arrInt[biggest] = arrInt[i] - arrInt[biggest]
				arrInt[i] = arrInt[i] - arrInt[biggest]
			}
			
		}	
	}
	
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