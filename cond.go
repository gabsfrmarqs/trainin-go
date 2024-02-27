package main

import (
	"fmt"
)

func main(){
	name := "Paramore"
	
	if name == "Paramore" {
		fmt.Println("Vou me matar hoje às 23:24")
	} else{
		fmt.Println("BOM DEMAIS")
	}

	// two thingies
	
	if x := 4; x < 4 {
		fmt.Println("eba menor de três")
	} else if x == 4 {
		fmt.Println("no exato momento em que estou escrevendo isso, um casal está se pegando loucamente na minha frente na praça de alimentação da faculdade, alguém elimine minha vida por favor")
	} else {
		fmt.Println("huuum lúcia no céu com diamantes")
	}
}
