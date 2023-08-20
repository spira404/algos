package main

import(
	"fmt"
)

func qSort(slice []int) ([]int){
	
	if len(slice) < 2{
		return slice
	}

	var pivot int
	pivot = len(slice) / 2
	less := []int{}
	greater := []int{}

	for i := 0; i < len(slice); i++{
		if i != pivot{
			if slice[i] <= slice[pivot]{
				less = append(less, slice[i])
			} else{
				greater = append(greater, slice[i])
			}
		}
	}
	less = qSort(less)
	greater = qSort(greater)
	one := append(less, slice[pivot])
	two := append(one, greater...)
	return two
}


func main() {
	// test slice
	slice := []int{18, 19, 49, 221, 2322, 319, 19}
	fmt.Printf("%v - original slice \n", slice)
	slice2 := qSort(slice)
	fmt.Printf("%v - sorted slice \n", slice2)
}
