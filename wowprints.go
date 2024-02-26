package main

import ("fmt")

func main(){
	var str1 string = "roku bun"
	var str2 string = "no ichi"
	
	fmt.Print("This is a print with newline using 'backslash n'\n")
	fmt.Println("This is a print using newline from Println")
	fmt.Printf("This is a print using prinft: %v", str1)
	fmt.Println(" " ,str2)

	// %v prints value
	// %T prints type
}

/*
Also, var can be used inside/outside functions.
Inferring works only inside a function
*/
