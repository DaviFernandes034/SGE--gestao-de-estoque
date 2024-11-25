package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/gin-gonic/gin"
)


func PostStatus(db *sql.DB) gin.HandlerFunc{

 return func(c *gin.Context) {

	var request models.StatusRequest

	err:= c.Bind(&request)
	if err!= nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "dados invalidos"})
		return
	}

	//inserindo um novo status
	err = services.InsertStatus(db, request.Nome)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	
	}

	

	c.JSON(http.StatusOK, gin.H{"Observação": "Status adicionado com sucesso"})
 }
}


func GetStatus(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		id:= c.Param("id")
		statusId, err:= strconv.ParseInt(id,10,64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "id invalido"})
		}


		statusId, statusNome, err:= services.GetStatus(db, statusId)
		if err != nil{

			c.JSON(http.StatusInternalServerError, gin.H{"erro": "status não encontrado no servidor"})
			return
		}

		response:= models.Status{

			Id_status: int(statusId),
			Nome: statusNome,
		}

		c.JSON(http.StatusOK, response)
	}
}


func GetStatusAll(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {

		status, err:= services.GetAllstatus(db)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, status)
	}
}