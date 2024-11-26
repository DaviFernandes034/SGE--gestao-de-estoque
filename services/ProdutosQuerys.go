package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"


	_"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	
)


func InsertProduto(db *sql.DB, categoriaId int64,nome string, preco float64, lote string, validade time.Time)(error){


	var exists bool
	var existsId bool //verificando se o id passado, existe na tabela categoria

	err:= db.QueryRow("select count(*) from Produtos where lote = @lote", 
	sql.Named("lote", lote)).Scan(&exists)
	if err != nil{
		return errors.New("erro ao verificar a existencia dos lotes")
	}

	if exists{
		return errors.New("lote do produto ja existente no banco de dados")
	}

	err1 := db.QueryRow("select count(*) from Categorias where ID_Categoria = @categoriaId",
	sql.Named("categoriaId",categoriaId )).Scan(&existsId)
	if err1 != nil{
		return errors.New("erro ao verificar a existencia do ID da categoria")
	}

	if !existsId{

		return fmt.Errorf("categoriaId não existe na tabela categoria")

	}

	if existsId{

		query:= `insert into Produtos (ID_categoria, nome, preco, lote, validade)
		values ( @categoriaID, @nome, @preco, @lote, @validade); 
		SELECT SCOPE_IDENTITY();`
	
		var lastId sql.NullInt64
		
		err = db.QueryRow(query,
		sql.Named("categoriaID",categoriaId),
		sql.Named("nome", nome), 
		sql.Named("preco", preco),
		sql.Named("lote", lote), 
		sql.Named("validade",validade)).Scan(&lastId)
		if err != nil{

				return fmt.Errorf("erro ao adicionar um produto: %v", err)
			
		}
	
		if !lastId.Valid{
			fmt.Errorf("erro ao identificar o id gerado")
		}
	
	}

	return nil
}