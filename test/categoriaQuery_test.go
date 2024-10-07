package test

import (

	"testing"

	
	"github.com/DATA-DOG/go-sqlmock"
)

	func TestInsertCategoria(t *testing.T) {
		// Cria o mock do banco de dados
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("erro ao criar mock do banco de dados: %v", err)
		}
		defer db.Close()
	
		categoria := "NovaCategoria"
	
		// Configura o mock para a verificação de existência da categoria
		mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM Categorias WHERE nome = \\?").
			WithArgs(categoria).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
	
		// Configura o mock para a inserção da categoria e retorno do SCOPE_IDENTITY()
		mock.ExpectQuery("INSERT INTO Categorias \\(nome\\) VALUES \\(\\?\\); SELECT SCOPE_IDENTITY\\(\\);").
			WithArgs(categoria).
			WillReturnRows(sqlmock.NewRows([]string{"scope_identity"}).AddRow(1))
	
		// Chama a função que será testada
		lastId, err := querys.InsertCategoria(db, categoria)
		if err != nil {
			t.Errorf("erro ao inserir a categoria: %v", err)
		}
	
		if lastId != 1 {
			t.Errorf("esperava id 1, mas recebeu %d", lastId)
		}
	
		// Verifica se todas as expectativas foram atendidas
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("expectativas não atendidas: %v", err)
		}
	}