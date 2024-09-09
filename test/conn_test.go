package test

import (
	"log"
	"os"
	
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/data"

)

func TestConn(t *testing.T) {


	// Carregar variáveis de ambiente do arquivo .env para o teste
	err := godotenv.Load("../.env")
	assert.NoError(t, err, "erro ao carregar o arquivo .env")

	// Verificar se as variáveis foram carregadas
	dbServer := os.Getenv("DB_SERVER")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	log.Printf("DB_SERVER: %s, DB_PORT: %s, DB_USER: %s, DB_PASSWORD: %s, DB_DATABASE: %s", dbServer, dbPort, dbUser, dbPassword, dbDatabase)

	// Chamar a função Conn e verificar se não há erros
	db, err := data.Conn()
	assert.NoError(t, err, "erro ao se conectar ao banco de dados")
	assert.NotNil(t, db, "a conexão com o banco de dados deve ser não nula")

	// Fechar a conexão ao final do teste
	if db != nil {
		err = db.Close()
		assert.NoError(t, err, "erro ao fechar a conexão com o banco de dados")
	}
}
