package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func main() {
  pessoas := []pessoa{
	{nome: "Igor", idade: 26},
	{nome: "Maria", idade: 30},
	{nome: "João", idade: 18},
  }

  if pessoas[0].idade >= 18 {
	fmt.Println("Lista de pessoas:")
  }else if pessoas[1].idade >= 18 {
	fmt.Println("Lista de pessoas:")
  }else{
	fmt.Println("Nenhuma pessoa maior de idade")
  }
  
}