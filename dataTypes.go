package main

import "fmt"
import "unicode/utf8"

func main() {
	// basic data types - numbers, strings and boolean
	var int1 int8 = 45
	fmt.Println("the first type is int8:", int1)

	var int2 rune = 32516
	fmt.Println("int2 is rune and is basically synonym of int32:", int2)

	var int3 byte = 36
	fmt.Println("int3 is byte and is basically synonym of unit8:", int3)

	var string1 string = "hello world"
	var string2 string = `hello
	world`

	fmt.Println("string 1 is:", string1)
	fmt.Println("string 2 is:", string2)

	// to get length of string do not use len function instead use
	fmt.Println(utf8.RuneCountInString(string1))

	mybool := true
	fmt.Println("mybool is declared using the shorthand notation:", mybool)
}
