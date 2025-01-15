package main

import "fmt"

func main() {

    // Creating a channel 
    // Here  deadlock arises because 
    // no goroutine is writing 
    // to this channel so, the select  
    // statement has been blocked forever 

    // c := make(chan int) 
    // select { 
    // case <-c: 
    // } 

	// so to avoid this use the defaul statement
	c := make(chan int) 
    select { 
    case <-c: 
    default: 
        fmt.Println("!.. Default case..!") 
    } 
}