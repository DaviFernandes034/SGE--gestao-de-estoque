package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	_ "github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
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



func GetAllProdutos(db *sql.DB)([]models.ProdutosRequest, error){

	query:= `SELECT p.ID_Produto, p.nome, p.preco, p.lote, p.validade,p.ID_Categoria, 
			c.ID_Categoria,c.nome
			from Produtos p
			INNER JOIN Categorias c
			ON
			p.ID_Categoria = c.ID_Categoria`


	rows, err:= db.Query(query)
	if err != nil{

		return nil, fmt.Errorf("erro ao buscar os produtos: %v ", err)
	}

	defer rows.Close()

	var produtos []models.ProdutosRequest


	for rows.Next(){

		var produto models.ProdutosRequest
		var categoria models.Categorias

		if err:= rows.Scan(&produto.Id_produto, 
			&produto.Nome, 
			&produto.Preco, 
			&produto.Lote,
			&produto.Validade, 
			&produto.CategoriaId,
			&categoria.Id_categoria,
			&categoria.Nome); err != nil{

			return nil, fmt.Errorf("erro ao escanear os produtos")
		}

		produto.Categoria = categoria
		produtos = append(produtos, produto)

	}

	if err:= rows.Err(); err != nil{

		return nil, fmt.Errorf("erro ao interar sobre o produto")
	}

	return produtos, nil
}


func GetProduto(db *sql.DB, id_produto int64)(models.ProdutosRequest,models.Categorias, error){

	query:= `SELECT p.ID_Produto, p.nome, p.preco, p.lote, p.validade,p.ID_Categoria, 
			c.ID_Categoria,c.nome
			from Produtos p
			INNER JOIN Categorias c
			ON
			p.ID_Categoria = c.ID_Categoria
			where p.ID_Produto = @id_produto`

		
		var produto models.ProdutosRequest
		var categoria models.Categorias

		err:= db.QueryRow(query, sql.Named("id_produto", id_produto)).Scan(&produto.Id_produto, 
			&produto.Nome, 
			&produto.Preco, 
			&produto.Lote,
			&produto.Validade, 
			&produto.CategoriaId,
			&categoria.Id_categoria,
			&categoria.Nome)
		if err != nil{

			return models.ProdutosRequest{},models.Categorias{},
			fmt.Errorf("erro ao mostrar o produto: %v", err)
		}

		produto.Categoria = categoria

		return produto, categoria, nil

}


func DeleteProduto(db *sql.DB, id int64)(error){


	query:= "delete from Produtos where ID_Produto = @id"

	result, err:= db.Exec(query, sql.Named("id", id))
	if err != nil{

		return fmt.Errorf("erro ao executar a query: %v", err)
	}

	row, err:= result.RowsAffected()
	if err != nil{
		return fmt.Errorf("erro ao verificar as linhas afetadas: %v ", err)
	}

	if row == 0{
		return fmt.Errorf("nenhuma categoria encontrada com o id: %d", id)
	}

	return nil
}