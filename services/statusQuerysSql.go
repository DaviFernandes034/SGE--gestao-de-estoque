package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

func InsertStatus(db *sql.DB, status string)( error){

	//vendo se o status ja existe
	var exists bool

	err:= db.QueryRow("select count(*) from status where nome = @nome",sql.Named("nome", status)).Scan((&exists))
	if err != nil{
		return  errors.New("erro ao verificar a existencia do Status")

	}

	if exists {
		return  errors.New("status ja existente")
	}

	// query para adicionar um status
	query:= "insert into status (nome) values(@nome); SELECT SCOPE_IDENTITY();"
	//insert into Categorias (nome) values(@nome); SELECT SCOPE_IDENTITY();"

	var lastId sql.NullInt64

	err = db.QueryRow(query,sql.Named("nome", status)).Scan((&lastId))
	if err != nil {
		return fmt.Errorf("erro ao inserir um status: %v", err)

	}

	//vendo se o id é valido
	if !lastId.Valid{
		return fmt.Errorf("erro ao recuperar id gerado: %v", err)
	}

	return  nil
}


func GetStatus(db *sql.DB, statusID int64)(int64, string, error){

	query:= " select ID_Status, nome from status where ID_Status = @statusID"

	var lastID sql.NullInt64
	var status sql.NullString

	err:= db.QueryRow(query, sql.Named("statusID",statusID)).Scan(&lastID, &status)
	if err != nil {

		if err == sql.ErrNoRows{
			return 0, " ", fmt.Errorf("status não encontrado: %v", err)
		}

		return 0, " ", fmt.Errorf("erro ao mostrar um status: %v", err)

	}

	return lastID.Int64, status.String, nil
}	

func GetAllstatus(db *sql.DB)([]models.Status, error){

	query:= "select * from status"

	row, err:= db.Query(query)
	if err != nil{

		return nil, fmt.Errorf("erro ao buscar os status: %v", err)

	}

	defer row.Close()

	var statusSlice []models.Status

	for row.Next(){

		var status models.Status

		if err:= row.Scan(&status.Id_status, &status.Nome); err != nil{
			return nil, fmt.Errorf("erro ao escanear o status: %v", err)
		}

		statusSlice = append(statusSlice, status)
	}

	if row.Err(); err != nil {
		return nil, fmt.Errorf("erro ao interar sobre os elementos ")
	}

	return statusSlice, nil
}

