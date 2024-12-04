package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/gin-gonic/gin"
)


func PostControle(db *sql.DB)gin.HandlerFunc{

	return func(c *gin.Context) {

		var request models.Controle

		err:= c.BindJSON(&request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"erro":"dados invalidos"})
			return
		}

		err = services.InsertControle(db,request.Id_produto, 
			request.Id_status,request.Quantidade,request.Data)
		if err != nil{

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"observação":"controle criado com sucesso"})
	}
}


func GetAllControle(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		controles, err:= services.GetAllControle(db)
		if err != nil{

			c.JSON(http.StatusInternalServerError, gin.H{"erro":err.Error()})
		}

		c.JSON(http.StatusOK, controles)

	}
}


func GetControle(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")
		
		controleId,err:= strconv.ParseInt(id,10,64)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": "id invalido"})
			return
		}

		controle, err:= services.GetControle(db, controleId)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"erro":"controle não achado no servidor"})
			return
		}


		//response

		response:=models.ControleRequest{
			 Id_controle: controle.Id_controle,
			 Produto: models.ProdutosRequest{
				Id_produto: controle.Produto.Id_produto,
				Nome: controle.Produto.Nome,
				Preco: controle.Produto.Preco,
				Lote: controle.Produto.Lote,
				Validade: controle.Produto.Validade,
				Categoria: models.Categorias{
						Id_categoria: controle.Produto.Categoria.Id_categoria,
						Nome: controle.Produto.Categoria.Nome,
				},
			 },
			 	Status: models.Status{
					Id_status: controle.Status.Id_status,
					Nome: controle.Status.Nome,
				},
		}

		c.JSON(http.StatusOK, response)
	}	
}

func DeleteControle(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")
		IdControle, err:= strconv.Atoi(id)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro":"id invalido"})
		}

		err = services.DeleteControle(db, int64(IdControle))
		if err != nil{

			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}