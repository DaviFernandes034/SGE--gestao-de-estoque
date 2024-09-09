package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/data/querys"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/gin-gonic/gin"
)


func ConfigRoutersCategory(r *gin.Engine){

	r.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"OK",
		})
	})

	
	



}

func GetCategoria(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context){

		//pegando o parametro id da URL
		id:= c.Param("id")

		//convertando id(string) para int64
		categoriaID, err:= strconv.ParseInt(id,10,64)
		if err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"erro": "id invalido"})
			return
		}


		//chamando a funçao do banco de dados
		categoriaID,categoriaNome,err:= querys.GetCategoria(db, categoriaID)
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "categoria nao encontrada"})
			return
		}

		//corpo da resposta
		response:= models.Categorias{
			Id_categoria: categoriaID,	
			Nome: categoriaNome,
		}

		c.JSON(http.StatusOK, response) //status ok

		
	}
}

func PostCategoria(db *sql.DB) gin.HandlerFunc{


	return func(c *gin.Context) {

		var request models.CategoriasRequest
		
		err:= c.BindJSON(&request)
		if  err != nil{
		   c.JSON(http.StatusBadRequest, gin.H{"error": "dados invalidos"})
		   return
	   }

	     
	   //inserindo a categoria ao banco de dados

	   categoriaId,err := querys.InsertCategoria(db, request.Nome)
	   if err != nil{
		   c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		   return
	   }

	   //response

	   response:= models.Categorias{
		   
		   Id_categoria: categoriaId,
		   Nome: request.Nome,
	   }

	   c.JSON(http.StatusOK, response)


	}
	
}



