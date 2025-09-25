package main

import (
	//"CRUD_EM_GO/class"

	roteador "github.com/igorrm19/aprendendo_go/class"
	config "github.com/igorrm19/aprendendo_go/class/config"
)

func main() {
	config.DataBeseConfig()
	roteador.Rota("/root")
	roteador.Servidor()
}
