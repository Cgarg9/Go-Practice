package main

import "fmt"

func main(){
	// key is string and value is int
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	myMap["Chirag"] = 21
	fmt.Println(myMap)
	// if key does not exist then default value of the datatype will be returned which in our case would be 0
	// they also return a second value which is true if the key exists else false
	var age, ok = myMap["Rahul"]
	if ok{
		fmt.Println("The age is %v", age)
	}else {
		fmt.Println("Invalid key")
	}

	// to delete a value
	// delete(myMap, "Chirag")

	for name := range myMap {
		fmt.Println("Name is", name)
	}
}