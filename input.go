package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escreve algo: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Eba tomale: ", input)


	// Input de string convertendo para número
	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Valor: ", aFloat)
	}




}
