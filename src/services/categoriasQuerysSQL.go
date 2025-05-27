package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

func InsertCategoria(db *sql.DB, categoria string) (error) {

	//verificando se a categoria ja existe
	var exists bool

	err := db.QueryRow("select count(*) from Categorias where nome = @nome",
	sql.Named("nome", categoria)).Scan(&exists)
	if err != nil {
		return errors.New("erro ao verificar a existencia da categoria")
	}

	if exists {
		return errors.New("categoria ja existente")
	}
	//query para adicionar uma categoria
	querys := "insert into Categorias (nome) values(@nome); SELECT SCOPE_IDENTITY();"
	

	var lastId sql.NullInt64

	err = db.QueryRow(querys, sql.Named("nome", categoria)).Scan(&lastId)
	if err != nil {
		return fmt.Errorf("erro ao inserir a categoria: %v", err)
	}

	//verificando se o id é valido
	if !lastId.Valid {
		return fmt.Errorf("erro ao recuperar id gerado %v", err)
	}

	return nil
}

func GetCategoria(db *sql.DB, categoriaID int64) (int64, string, error) {

	querys := "select ID_Categoria, nome from Categorias where ID_Categoria = @categoriaID"

	var lastId sql.NullInt64
	var nome sql.NullString

	err := db.QueryRow(querys, sql.Named("categoriaID", categoriaID)).Scan(&lastId, &nome)
	if err != nil {

		if err == sql.ErrNoRows {

			return 0, "", fmt.Errorf("categoria não encontrada")
		}

		return 0, "", fmt.Errorf("erro ao mostrar uma categoria: %v", err)
	}

	return lastId.Int64, nome.String, nil
}

func GetAllCategoria(db *sql.DB) ([]models.Categorias, error) {

	query := "select ID_Categoria, nome from Categorias" //query que irá executar

	rows, err := db.Query(query) //executando a query
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar as categorias: %v", err)
	}

	defer rows.Close() //fechando as linhas

	var categorias []models.Categorias // criando o slice para armazenar as categorias

	for rows.Next() { //iterando sobre os resultados da query

		var categoria models.Categorias

		//copia os valores do banco de dados e os passa para a struct
		if err := rows.Scan(&categoria.Id_categoria, &categoria.Nome); err != nil {
			return nil, fmt.Errorf("erro ao escanear categoria: %v", err)
		}

		categorias = append(categorias, categoria) //adicionando uma categoria no slice categorias
	}

	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %v", err)
	}

	return categorias, nil

}

func DeleteCategoria(db *sql.DB, Categoria_id int) error {

	//query
	query := "delete from Categorias where ID_Categoria = @Categoria_id"

	//executando a query

	result, err := db.Exec(query, sql.Named("Categoria_id", Categoria_id))
	if err != nil {
		return fmt.Errorf("erro ao deletar a categoria: %v", err)
	}

	//verificando se alguma linha foi afetada
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar lihas afetadas: %v", err)

	}

	if rows == 0 {
		return fmt.Errorf("nenhuma categoria encontrada com o id: %d.", Categoria_id)
	}

	return nil
}
