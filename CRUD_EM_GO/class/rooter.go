package roteador

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		data := map[string]interface{}{
			"email": user.email,
			"senha": user.senha,
		}
		//ctx.JSON(200, gin.H{"mensagem": "rota acessada"})
		ctx.AsciiJSON(http.StatusOK, data)
	})
}

func Servidor() {
	r.Run()
}
