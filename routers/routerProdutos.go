package routers

import (
	"database/sql"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/gin-gonic/gin"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)


func PostProduto(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		var request models.Produtos

		err:= c.BindJSON(&request)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados invalidos"})
			return
		}

		err = services.InsertProduto(db,request.CategoriaId, request.Nome, request.Preco, request.Lote, request.Validade)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"observação": "Produto criado com sucesso"})
	}
}
