package main

import (
	roteador "github.com/igorrm19/aprendendo_go/class"
	"github.com/igorrm19/aprendendo_go/class/config"
)

func main() {

	roteador.Rota("/root")

	dbconnection := config.DBconfig()
	defer dbconnection.Close()

	roteador.Servidor()

}
