package main

import (
	"fmt"
)

func main(){
	colors := []string{"this", "shadow", "world"}
	fmt.Println(colors)

	for i := 0; i<len(colors);i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {			//sexier, and no counters also!
		fmt.Println(colors[i])		//não acho que vou usar anyway
	}

	for _, color := range colors {		//index ignorado
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	sum := 1
	for sum < 1000{
		sum += sum
		fmt.Println("Sum:", sum)
		if sum < 200 {
			goto theEnd //lol
		}
	}

	theEnd: fmt.Println("Fiim") 



}
