package main

import (
	"fmt"
	"math"
)


func main(){
	f1, f2, f3 := 23.4, 2.4, 5.0
	floatSum := f1 + f2 + f3
	
	// ugly way
	fmt.Println("Float sum (normal):  ", floatSum)
	
	// shitass way
	floatSum = math.Round(floatSum*100)/100
	fmt.Println("Float sum (shitass): ", floatSum)

	//the python way
	fmt.Println("Float sum (toptop): %.2f",floatSum) 

}
