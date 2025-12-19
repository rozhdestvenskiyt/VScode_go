package main

import "fmt"


func add(num_1 int , num_2 int) int {
	result := num_1 * num_2
	return result
}

func main() {
	fmt.Printf("The result is %d" , add(5 ,5))
}