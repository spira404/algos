package main

import(
	"fmt"
)

func selSort(slice []int, n int) ([]int){
	
	//for every value of slice
	for i := 0; i <	n - 1; i ++{
		least := i
		// check other values which is least
		for j := i + 1; j < n; j++{
			if slice[j] < slice[least]{
				least = j
			}
		}
		// swap value of current index (i) and least value
		tmp := slice[least]
		slice[least] = slice[i]
		slice[i] = tmp
	}
	return slice
}


func main() {
	// test slice
	slice := []int{18, 19, 49, 221, 2322, 319, 19}
	fmt.Printf("%v - original slice \n", slice)
	slice2 := selSort(slice, len(slice))
	fmt.Printf("%v - sorted slice \n", slice2)
}
