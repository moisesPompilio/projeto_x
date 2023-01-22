package main

import (
	"fmt"
	"net/http"

	"github.com/moisesPompilio/projeto_x/src/adapters/htpp/routes"
	"github.com/moisesPompilio/projeto_x/src/pkg/config/server"
	"github.com/moisesPompilio/projeto_x/src/pkg/migrates"
)

func main() {
	port, err := server.Load_Server()
	migrates.StarMigrate()

	if err != nil {
		panic(err)
	}

	r := routes.LoadRouter()

	fmt.Println("server rodando na porta:", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}
