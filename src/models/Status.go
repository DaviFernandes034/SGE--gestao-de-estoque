package models


type Status struct{

	Id_status int `json:"status_id"`
	Nome string `json:"Nome"`
}


type StatusRequest struct{

	Nome string `json:"Nome"`
}