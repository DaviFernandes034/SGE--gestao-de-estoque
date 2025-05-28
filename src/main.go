package main

import (

	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/joho/godotenv"
)

func main(){

	log.Println("SGE-SISTEMA DE GESTAO DE ESTOQUE!")
	log.Print("desenvolvedor: Davi fernandes")
	//carregando os arquivos .env
	err:= godotenv.Load()
	if err != nil {

		log.Fatalf("erro ao carregar o arquivo .env: %v", err)
	}

	log.Println("ARQUIVOS .ENV CARREGADOS")

	var InitConnection configs.InitConnection //iniciando a inicializaçao do banco de dados

	_,err = InitConnection.Init() //chamando a funcao para iniciar a conexão
	if err != nil {

		log.Fatalf("erro ao iniciar conexão: %v", err)
	}

	log.Println("INFO: CONEXÃO COM O BANCO DE DADOS, COMPLETA!")

	log.Println("INFO: aplicação iniciada com sucesso!!")
	log.Println("----------------------------------------------------------------")

	http.ListenAndServe(":8080", nil)
	
}


