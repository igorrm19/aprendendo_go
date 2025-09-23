package main

import (
	  "primeiro_projeto_go/class"
	  "fmt"
)

func main(){
  valor := "valor da variavel"
  Objeto := classe.Desconhecida{
	ValorUndefaidString: "string indefinida",
	ValorAserUndfaidInt: 0, 
	ValorUndefindedBool: false}

  fmt.Println(classe.Publica(valor))
  fmt.Println(Objeto.ValorUndefaidString)
}

// go run objetos_em_go.go