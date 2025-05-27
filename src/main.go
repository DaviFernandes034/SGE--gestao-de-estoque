package main

import (
	"log"
	"os"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/data" // Importa o pacote data
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/routers"
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

	r.POST("/categoria", routers.PostCategoria(db))
	r.GET("/categoria/:id", routers.GetCategoria(db))
	r.GET("/categoriaAll", routers.GetAllCategoriaRouter(db))
	r.DELETE("/categoriaDelete/:id", routers.DeleteCategoria(db))
	//rotas Status
	r.POST("/status", routers.PostStatus(db))
	r.GET("/status/:id", routers.GetStatus(db)) 
	r.GET("/statusAll", routers.GetStatusAll(db))
	r.DELETE("/statusDelete/:id", routers.DeleteStatus(db))
	//rotas Produtos
	r.POST("/produto", routers.PostProduto(db))
	r.GET("/produtoAll", routers.GetAllProdutos(db))
	r.GET("/produto/:id", routers.GetProduto(db))
	r.DELETE("/produtoDelete/:id", routers.DeleteProduto(db))
	// rotas controle
	r.POST("/controle", routers.PostControle(db))
	r.GET("/controleAll", routers.GetAllControle(db))
	r.GET("controle/:id", routers.GetControle(db))
	r.DELETE("controleDelete/:id", routers.DeleteControle(db))

	

	//iniciar o servidor na porta 8080

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	r.Run(":" + port)

}



