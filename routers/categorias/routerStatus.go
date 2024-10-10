package routers

import (
	"database/sql"
	"net/http"

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
	statusId, err:= services.InsertStatus(db, request.Nome)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	
	}

	//response

	response:= models.Status{
		Id_status: int(statusId),
		Nome: request.Nome,
	}

	c.JSON(http.StatusOK, response)
 }
}