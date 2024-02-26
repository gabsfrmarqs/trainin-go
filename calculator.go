package main

import ("fmt")

func main(){
	var num1 int = 3
	var num2 int = 7
	num3 := 5 // type is being inferred
	
	
	fmt.Println("num1: ", num1)
	fmt.Println("num2: ", num2)
	fmt.Println("num3: ", num3)

	fmt.Println(num1 + num2)
	fmt.Println(num1 + num3)
	fmt.Println(num2 + num3)
}
