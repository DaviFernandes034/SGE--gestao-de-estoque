package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategoriaRouter(db *sql.DB) gin.HandlerFunc{

	return func (c *gin.Context){

		categorias, err:= services.GetAllCategoria(db)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error() })
			return
		}
			c.JSON(http.StatusOK, categorias)

	}
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
		categoriaID,categoriaNome,err:= services.GetCategoria(db, categoriaID)
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

	   categoriaId,err := services.InsertCategoria(db, request.Nome)
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

func DeleteCategoria(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")
		categoriaId, err:= strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalido"})
			return
		}

		err = services.DeleteCategoria(db, categoriaId)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}



