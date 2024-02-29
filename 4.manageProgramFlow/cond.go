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

	numero := 4
	var saida string
	switch numero {
	case 0:
		saida = "leite!"
//		fallthrough // em go, switches tem break por padrão. fallthrough retorna ao comportamento que seria esperado em C, por exemplo
	case 1: 
		saida = "miku miku ooo eee ooo"
	default:
		saida = "essas coisa ai"
	}

	fmt.Println(saida)



}
