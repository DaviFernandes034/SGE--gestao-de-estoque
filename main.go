package main

import (
	"log"
	"os"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/data" // Importa o pacote data
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/routers/categorias"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Conectar ao banco de dados
	db, err := data.Conn()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	log.Println("Conexão com o banco de dados estabelecida com sucesso")

	//carregando as rotas

	// roteador gin
	r := gin.Default()

	//configurar as rotas CATEGORIA

	r.POST("/categorias", routers.PostCategoria(db))
	r.GET("/categoria/:id", routers.GetCategoria(db))
	r.GET("/categoriaAll", routers.GetAllCategoriaRouter(db))
	r.DELETE("/categoriaDelete/:id", routers.DeleteCategoria(db))
	


	//iniciar o servidr na porta 8080

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	r.Run(":" + port)

}
