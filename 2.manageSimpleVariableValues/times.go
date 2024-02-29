package main

import (
	"fmt"
	"time"
)

func main(){
	n := time.Now()
	fmt.Println(n)

	t := time.Date(2024, time.February, 27, 17, 19, 0, 0, time.UTC)
	fmt.Println("Look!: ", t)
	fmt.Println(t.Format(time.ANSIC))

}

