package main

import "fmt"

func main() {
	// fixed size, all elements are of same data type
	var intArr [3]int32
	intArr[0] = 242
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// memory address of the elements
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	// as you can see these are coniguously allocated  

	// another way to initialize
	intArr2 := [3]int32{1,2,3}
	fmt.Println(intArr2[0:3])

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v and the capacity is %v", len(intSlice), cap(intSlice))
	fmt.Println()
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v and the capacity is %v", len(intSlice), cap(intSlice))
	// even though the capacity has doubled we cannot access index 4 and 5
}