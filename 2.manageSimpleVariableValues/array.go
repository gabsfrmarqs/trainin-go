package main

import ("fmt")

func main(){
	var arr1 = [5]int{6,9,4,2,0}	
	arr2 := [5]int{0,2,4,9,6}
	arr3 := [...]int{0,2,4,9,6} //inferred lenght


	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("Initialization: ")
	
	arr4 := [5]int{}
	arr5 := [5]int{1,2}
	arr6 := [5]int{1,2,3,4,5}


	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)

	arr7 := [5]int{1:10,2:40} //initialize value in index 1 and 2

	fmt.Println(arr7)
	fmt.Println(len(arr3))

}
