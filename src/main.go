package main

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
)

func main(){

	var InitConnection configs.InitConnection

	err:= InitConnection.Init()
	if err != nil {

		log.Fatalf("erro ao iniciar conexão: %w", err)
	}

	

	log.Println("INFO: aplicação iniciada com sucesso!!")
	http.ListenAndServe(":8080", nil)
	
}


