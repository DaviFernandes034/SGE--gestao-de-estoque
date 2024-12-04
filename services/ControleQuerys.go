package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)


func InsertControle(db *sql.DB, id_Produto, id_Status,
	quantidade int64,data time.Time)(error){

		var existsProdutoId bool
		var existsStatusId bool
		var existsProdutoIdControle bool

		err:= db.QueryRow("select count(*) from Produtos where ID_Produto = @id_produto",
		sql.Named("id_produto",id_Produto)).Scan(&existsProdutoId)
		if err != nil{
			return errors.New("erro ao verificar a existencia do produto")
		}

		err1:= db.QueryRow("select count(*) from status where ID_Status = @id_status",
		sql.Named("id_status", id_Status)).Scan(&existsStatusId)
		if err1 != nil{
			return errors.New("erro ao verificar a existencia do status")
		}

		err2:= db.QueryRow("select count(*) from Controle where ID_Produto = @Id_produto",
		sql.Named("Id_produto", id_Produto)).Scan(&existsProdutoIdControle)
		if err2 != nil {
				return errors.New("erro ao verificar se existe 2 id de produto igual")

		}
		

		if !existsProdutoId{

			return fmt.Errorf("id do produto não encontrado")
		}

		if !existsStatusId{

			return fmt.Errorf("id do status não encontrado")
		}

		if existsProdutoIdControle{

			return fmt.Errorf("id: %v , ja se encontra no Controle", id_Produto)
		}

		if existsProdutoId && existsStatusId{


			query:= `insert into Controle ( ID_Produto, ID_Status, quantidade, data)
				     values ( @id_produto, @id_status, @quantidade, @data);
					 SELECT SCOPE_IDENTITY();`

			var lastId sql.NullInt64

			err := db.QueryRow(query,
			sql.Named("id_produto", id_Produto),
			sql.Named("id_status", id_Status),
			sql.Named("quantidade", quantidade),
			sql.Named("data",data)).Scan(&lastId)
			if err != nil {

				return fmt.Errorf("erro ao adicionar um controle")
			}

			if !lastId.Valid {

				fmt.Errorf("erro ao identificar o id gerado")
			}
		

		}



	return nil
}

func GetAllControle(db *sql.DB)([]models.ControleRequest, error){

		query:= `SELECT c.ID_Controle, 
						p.ID_Produto,
						p.nome,
						p.lote,
						p.preco,
						p.validade,
						cat.ID_Categoria,
						cat.nome,
						s.ID_Status,
						s.nome,
						c.quantidade,
						c.data
					from Controle c 
					INNER join Produtos p
					on c.ID_Produto = p.ID_Produto
					inner join Categorias cat
					on p.ID_Categoria = cat.ID_Categoria
					inner JOIN status s
					on c.ID_Status = s.ID_Status`
        
		rows, err:= db.Query(query)
		if err != nil {

			return nil, fmt.Errorf("erro ao acessar os dados do controle")
		}

		defer rows.Close()

		var Controles []models.ControleRequest

		for rows.Next(){

			var Controle models.ControleRequest

			if err:= rows.Scan(
				&Controle.Id_controle,
				&Controle.Produto.Id_produto,
				&Controle.Produto.Nome,
				&Controle.Produto.Lote,
				&Controle.Produto.Preco,
				&Controle.Produto.Validade,
				&Controle.Produto.Categoria.Id_categoria,
				&Controle.Produto.Categoria.Nome,
				&Controle.Status.Id_status,
				&Controle.Status.Nome,
				&Controle.Quantidade,
				&Controle.Data); err != nil{

					return nil, fmt.Errorf("erro ao escanear o controle")
				}

				Controles = append(Controles, Controle)
		}

		if err:= rows.Err(); err != nil{

			return nil, fmt.Errorf("erro ao interar sobre o controle")
		}

	return Controles, nil
}


func GetControle(db *sql.DB, id_controle int64)(models.ControleRequest, error){

	query:= `SELECT c.ID_Controle, 
			p.ID_Produto,
			p.nome,
			p.lote,
			p.preco,
			p.validade,
			cat.ID_Categoria,
			cat.nome,
			s.ID_Status,
			s.nome,
			c.quantidade,
			c.data
		from Controle c 
		INNER join Produtos p
		on c.ID_Produto = p.ID_Produto
		inner join Categorias cat
		on p.ID_Categoria = cat.ID_Categoria
		inner JOIN status s
		on c.ID_Status = s.ID_Status
		where ID_Controle = @id_controle`

		var Controle models.ControleRequest

		err:= db.QueryRow(query, sql.Named("id_controle", id_controle)).Scan(
			&Controle.Id_controle,
			&Controle.Produto.Id_produto,
			&Controle.Produto.Nome,
			&Controle.Produto.Lote,
			&Controle.Produto.Preco,
			&Controle.Produto.Validade,
			&Controle.Produto.Categoria.Id_categoria,
			&Controle.Produto.Categoria.Nome,
			&Controle.Status.Id_status,
			&Controle.Status.Nome,
			&Controle.Quantidade,
			&Controle.Data)
		if err != nil{

				return models.ControleRequest{},fmt.Errorf("erro ao mostrar o controle")
		}



	return Controle, nil

}

func DeleteControle(db *sql.DB, id int64)(error){

	query:= `delete from Controle  where ID_Controle = @id`

	result, err:= db.Exec(query, sql.Named("id", id))
	if err != nil{
		return fmt.Errorf("erro ao executar a query")
	}

	row, err:= result.RowsAffected()
	if err != nil {
		 return fmt.Errorf("erro ao verificar as linhas afetadas")
	}

	if row == 0{

		return fmt.Errorf("nenhum linha encontrada com o id: %v", id)
	}



	return nil
}
