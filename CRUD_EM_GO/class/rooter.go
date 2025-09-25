package roteador

import "github.com/gin-gonic/gin"

type Usuario struct {
	email string
	senha string
}

var user = Usuario{
	email: "",
	senha: "",
}

var r = gin.Default()

func Rota(rota string) {
	r.GET(rota, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"mensagem": "rota acessada"})
	})
}

func Servidor() {
	r.Run()
}
