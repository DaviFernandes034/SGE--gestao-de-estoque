package dtos




type StatusDTO struct{

	Nome string `json:"nome"`
}


func NewStatusDto(nome string)*StatusDTO{

	return&StatusDTO{
		Nome: nome,
	}
}