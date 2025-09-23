package main

import "fmt"


var valorString string = "Hello word" //variavel de escopo global


func HelloWord(variavel string) {  // funcao que imprime uma string
    for i := 0; i < 5; i++ {
		valorString = fmt.Sprintf("%s %d", variavel, i+1) // Sprintf formata a string
		fmt.Println(valorString) // imprime a string formatada
	}
}

func main() {
 
	  HelloWord(valorString)

}