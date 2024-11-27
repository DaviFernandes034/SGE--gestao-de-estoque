package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/gin-gonic/gin"
)


func PostProduto(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		var request models.Produtos

		err:= c.BindJSON(&request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados invalidos"})
			return
		}

		err = services.InsertProduto(db,request.CategoriaId, request.Nome, request.Preco, 
		request.Lote, request.Validade)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"observação": "Produto criado com sucesso"})
	}
}



func GetAllProdutos(db * sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		produtos, err:= services.GetAllProdutos(db)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, produtos)
	}

}


func GetProduto(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")

		produtoId, err:= strconv.ParseInt(id,10,54)
		if err != nil{

			c.JSON(http.StatusBadRequest, gin.H{"erro":"id invalido"})
			return
		}

		produto,categoria, err:= services.GetProduto(db, produtoId)
		if err != nil{
			c.JSON(http.StatusInternalServerError, 
			gin.H{"erro":"Produto não encontrado no servidor"})
		}

		//response

		response:= models.ProdutosRequest{
			Id_produto: produto.Id_produto,
			Nome: produto.Nome,
			Lote: produto.Lote,
			Validade: produto.Validade,
			CategoriaId: produto.CategoriaId,
			Categoria: models.Categorias{
				Id_categoria: categoria.Id_categoria,
				Nome: categoria.Nome,
			},
		}

		c.JSON(http.StatusOK, response)
			
	}
}


func DeleteProduto(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")
		produtoId, err:= strconv.ParseInt(id, 10, 64)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"erro": "id invalid"})
		}

		err = services.DeleteProduto(db, produtoId)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		}

		c.Status(http.StatusNoContent)
	}
}