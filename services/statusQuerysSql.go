package services

import (
	"database/sql"
	"errors"
	"fmt"
)

func InsertStatus(db *sql.DB, status string)(int64, error){

	//vendo se o status ja existe
	var exists bool

	err:= db.QueryRow("select count(*) from status where nome = @nome",sql.Named("nome", status)).Scan((&exists))
	if err != nil{
		return 0, errors.New("erro ao verificar a existencia do Status")

	}

	if exists {
		return 0, errors.New("status ja existente")
	}

	// query para adicionar um status
	query:= "insert into status (nome) values(@nome); SELECT SCOPE_IDENTITY();"

	var lastId sql.NullInt64

	err = db.QueryRow(query,sql.Named("nome", status)).Scan((&lastId))
	if err != nil {
		return 0, fmt.Errorf("erro ao inserir um status: %v", err)

	}

	//vendo se o id é valido
	if !lastId.Valid{
		return 0, fmt.Errorf("erro ao recuperar id gerado: %v", err)
	}

	return lastId.Int64, nil
}