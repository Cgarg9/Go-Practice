package main

import ("fmt"
	"errors")

func main() {
	numerator := 12
	denominator := 3

	var quotient, remainder, err = division(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("The quotient is: ", quotient, " and the remainder is: ", remainder)
	fmt.Printf("The quotient is %v and the remainder is %v", quotient, remainder)

	fmt.Println("Sum of 1, 2, 3 and 4 is", sum(1,2,3,4))
	fmt.Println("Sum of 1, 2, 3 is", sum(1,2,3))
}

func division (numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var quotient int = numerator/denominator
	var remainder int = numerator%denominator

	return quotient, remainder, err
}

// varidic funtion - variable number of arguments
// only one variable argument allowed and ir must be the last parameter
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}