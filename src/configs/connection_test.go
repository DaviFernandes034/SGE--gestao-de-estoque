package configs

import (
	"os"
	"testing"

	
)


func TestConnSucess(t *testing.T){

	os.Setenv("DB_SERVER", "localhost")
	os.Setenv("DB_PORT", "1433")
	os.Setenv("DATABASE", "SGE")
	conn, err:= conn()
	if err != nil {

		t.Fatalf("erro. %v", err)
	}

	if conn.Db == nil {

		t.Fatalf("esperava uma *sql.DB")
	}
}

func TestInitConnection(t *testing.T){

	os.Setenv("DB_SERVER", "localhost")
	os.Setenv("DB_PORT", "1433")
	os.Setenv("DATABASE", "SGE")

	var initCon InitConnection

	_,err:= initCon.Init()
	if err != nil {

		t.Fatalf("erro: %v", err)
	}

	if initCon.Conn == nil || initCon.Conn.Db == nil {

		t.Fatalf("conexão não inicializada")
	}

	

}


