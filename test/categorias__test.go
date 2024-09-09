package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/routers/categorias"
)


func MockDb()(*sql.DB, sqlmock.Sqlmock, error){

	db,mock,err:= sqlmock.New()
	if err != nil{

		return nil, nil, err
	}

	return db, mock, nil
}

func TestGetAll(t *testing.T){

	//configura o router
	r := gin.Default()

	routers.ConfigRoutersCategory(r)

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

		gin.SetMode(gin.TestMode) //modo de testes do framework gin
		
		r:= gin.Default() //retorna uma instancia do gin

		db, mock, err:= MockDb() //criaçao do banco de dados de testes
		if err != nil{
			t.Fatalf("erro ao criar mock do banco de dados: %v", err)
		}

		defer db.Close() //fechando o banco de dados

		r.POST("/categorias", routers.PostCategoria(db)) //definindo as rotas

		
		mock.ExpectExec("insert into Categoria").WithArgs("alimento").WillReturnResult(sqlmock.NewResult(1,1))//comando sql mockado

		requestBody:= `{"nome":"alimento"}`//resultado esperado

		//passando o resultado para o metado post
		request, err:= http.NewRequest(http.MethodPost, "/categorias", bytes.NewBufferString(requestBody))
		if err != nil{
			t.Fatalf("erro ao criar requisição: %v ", err)

		} 

		test:= httptest.NewRecorder() //simulando uma requisilção http e inspecionando a resposta que a requisiçao gera
		r.ServeHTTP(test,request)//processa a requisiççao http "test": pega a resposta, "request" a requisiçao simulada 

		assert.Equal(t, http.StatusOK, test.Code) //verificada se o status é ok

		var response models.Categorias

		if err:= json.Unmarshal(test.Body.Bytes(), &response); err !=nil{

			t.Fatalf("erro no unmarshal response: %v", err)
		}

		assert.Equal(t, int(1), response.Id_categoria)
		assert.Equal(t, "alimento", response.Nome)


		//verificando se todas as configuraçoes do mock foram feitas certas
		if err:= mock.ExpectationsWereMet(); err != nil{
			t.Errorf("espectativas nao atendidas: %v", err)
		}
	
		 

}


func TestGet(t *testing.T){

	r:= gin.Default()

	routers.ConfigRoutersCategory(r)

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
