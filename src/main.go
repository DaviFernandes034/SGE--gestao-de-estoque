package main

import (
	
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/joho/godotenv"
)

func main(){

	
	//carregando os arquivos .env
	err:= godotenv.Load()
	if err != nil {

		log.Fatalf("erro ao carregar o arquivo .env: %v", err)
	}

	var InitConnection configs.InitConnection //iniciando a inicializaçao do banco de dados

	_,err = InitConnection.Init() //chamando a funcao para iniciar a conexão
	if err != nil {

		log.Fatalf("erro ao iniciar conexão: %v", err)
	}

	

	log.Println("INFO: aplicação iniciada com sucesso!!")
	http.ListenAndServe(":8080", nil)
	
}


