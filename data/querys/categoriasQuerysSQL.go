package querys

import (
	"database/sql"
	"errors"
	"fmt"

	
)


func InsertCategoria(db *sql.DB, categoria string )(int64, error){

		//verificando se a categoria ja existe
		var exists bool

		err:= db.QueryRow("select count(*) from Categorias where nome = @nome", sql.Named("nome", categoria)).Scan(&exists)
		if err != nil{
			return 0, errors.New("erro ao verificar a existencia da categoria")
		}

		if exists{
			return 0, errors.New("categoria ja existente")
		}
	//query para adicionar uma categoria
	querys:= "insert into Categorias (nome) values(@nome); SELECT SCOPE_IDENTITY();"

	var lastId sql.NullInt64

	err = db.QueryRow(querys,sql.Named("nome",categoria)).Scan(&lastId)
	if err != nil{
		return 0, fmt.Errorf("erro ao inserir a categoria: %v", err)
	}


	//verificando se o id é valido
	if !lastId.Valid{
		return 0, fmt.Errorf("erro ao recuperar id gerado %v", err)
	}

	return lastId.Int64, nil 
}
 
func GetCategoria(db *sql.DB, categoriaID int64)(int64, string, error){

	 querys:="select ID_Categoria, nome from Categorias where ID_Categoria = @categoriaID"

	 var lastId sql.NullInt64
	 var nome sql.NullString

	 err:= db.QueryRow(querys, sql.Named("categoriaID", categoriaID)).Scan(&lastId, &nome)
	 if err != nil{

		if err == sql.ErrNoRows{

			return 0, "", fmt.Errorf("categoria não encontrada")
		}

		return 0, "", fmt.Errorf("erro ao mostrar uma categoria: %v", err)
	 }

	return lastId.Int64, nome.String, nil
}
