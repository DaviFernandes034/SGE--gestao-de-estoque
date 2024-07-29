package routers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	
)


func TestGetAll(t *testing.T){

	//configura o router
	r := gin.Default()

	ConfigRoutersCategory(r)

	//criando uma requisição GET
	request,_:= http.NewRequest("GET", "/all", nil)
	response:= httptest.NewRecorder()
	
	//executando
	r.ServeHTTP(response, request)

	//verificando o statusCode

	if response.Code != http.StatusOK{

		t.Errorf("status code esperado %d , obtido %d ", http.StatusOK, response.Code)
	}

	//verificando o corpo da resposta
	expected:= `{"message":"OK"}`

	if response.Body.String() != expected{

		t.Errorf("resposta esperada %s, obtida %s", expected, response.Body.String())
	}


}


func TestPost(t *testing.T){


	r:= gin.Default()

	ConfigRoutersCategory(r)

	//criando uma requisição post

	load:= []byte("name=rada")

	request, _:= http.NewRequest("POST", "/categorias", bytes.NewBuffer(load))
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	response:= httptest.NewRecorder()

	//executa
	r.ServeHTTP(response, request)


	if  response.Code != http.StatusOK{

		t.Errorf("code esperado %d , obtido %d", http.StatusOK, response.Code)
	}

	//verificando o corpo da resposta

	expected:= `{"message":"bao rada"}`
	if response.Body.String() != expected{

		t.Errorf("corpo da resosta esperado %s , obtido %s", expected, response.Body.String())
	}

}


func TestGet(t *testing.T){

	r:= gin.Default()

	ConfigRoutersCategory(r)

	resquest, _:= http.NewRequest("GET", "/categoria/rada", nil)
	response:= httptest.NewRecorder()

	//executando

	r.ServeHTTP(response, resquest)

	if response.Code != http.StatusOK{

		t.Errorf("code esperado %d, obtido %d", http.StatusOK, response.Code)

	}

	//corpo da resquest
	expected:= `{"message":"bao rada"}`

	if response.Body.String() != expected{

		t.Errorf("corpo esperado %s, corpo obtido %s ", expected, response.Body.String())
	}




}
