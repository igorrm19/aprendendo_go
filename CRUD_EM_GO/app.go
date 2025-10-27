// pg_isready --> verifica se o servidor esta rodando
// sudo service postgresql start --> inicia servidor
// ss -nlt | grep 5432 --> verifica se esta ouvindo na porta 5432

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	roteador "github.com/igorrm19/aprendendo_go/class"
	"github.com/igorrm19/aprendendo_go/class/config"
	"github.com/igorrm19/aprendendo_go/handleres"
	"github.com/igorrm19/aprendendo_go/models"
)

func main() {

	roteador.Rota("/root")

	dbconnection := config.DBconfig()

	_, err := dbconnection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal("Erro no dbconnetion.exel em app.go")
	}

	defer dbconnection.Close()

	taskHendle := handleres.NewTaskHandle(dbconnection)

	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskHendle.ReadTask).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

/*

 func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello word"))
	}

*/
