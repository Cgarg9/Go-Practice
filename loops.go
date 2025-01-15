package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("value of i is", i)
	}

	// for as an infinite loop
	// for {
	// 	fmt.Println("hi")
	// }

	// for as while loop
	i := 0
	for i < 4 {
		fmt.Printf("The value of i in while is %v \n", i)
		i++
	}

	// for strings
	for i, j := range "Chirag" {
		fmt.Printf("The index is %v and the character is %v \n", i, j)
	}

	//  A for loop can iterate over the sequential values sent on the channel until it closed.
	chnl := make(chan int)
	go func(){
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i:= range chnl {
		fmt.Println(i)
	}
}
