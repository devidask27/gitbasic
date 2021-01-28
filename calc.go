package main

import "fmt"

func main()  {
	val1,val2 := 2 , 3
	fmt.Println(Addition(val1,val2))
	fmt.Println(Multiplication(val1,val2))
	
}

func Addition(val1, val2 int) int  {
	return val1 + val2
}

func Multiplication(val1, val2 int) int {
	return val1 * val2
}