package main


import "fmt"

func fib(arr []uint, count uint) uint {
	if(count == 1) {
		arr[count - 1] = 0
	 return 0
	}
	if(count == 2) {
		arr[count - 1] = 1
	return 1
	}
	if(arr[count - 1] != 0) {
		return arr[count - 1]
	}

	n := fib(arr, count - 2) + fib(arr,count - 1)
	arr[count - 1] = n

	return n
}

func main() {

	arr := make([]uint, 100)
	fib(arr,100)

	fmt.Println(arr)
}