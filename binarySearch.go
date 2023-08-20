package main

import(
	"fmt"
)

func binSearch(slice []int, val int) (int) {
	min := 0
	max := len(slice)
	var mid int
	mid = (max - min) / 2
	for mid > 1{
		mid = (max + min) / 2
		if slice[mid] > val{
			max = mid
		} else if slice[mid] < val{
			min = mid
		} else{
			return mid
		}
	}
	return -1
	
}

func main(){
	slice := []int{1, 15, 17, 58, 59, 100, 103, 105, 144}
	val := 100
	idx := binSearch(slice, val)
	fmt.Printf("Index of value %v in slice is %v \n", val, idx) 
}
