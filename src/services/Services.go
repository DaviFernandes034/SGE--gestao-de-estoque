package services

import (
	"fmt"
	"log"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)

//baseService é um tipo generico, aonde nao preciso ficar digitando a mesma coisa no resto dos arquivos de Service
//"[t any]", Aceita qualquer model que eu passar
type BaseService[t any] struct {
	Repo interfaces.RepositoryCrud[t] //interface generica do repository contendo os metados CRUD

}
/* 
	construtor da struct, passo um dos models no lugar de "t any", e como argumento passo a interface contendo os cruds, 
	em "[t], passo um dos models criados"

	em "interfaces.ServiceCrud[t]", indica que a struct que eu retornar vai implentar a interface ServiceCrud[t]
	(contendo os metados em que todos os servives vao utilizar), novamente em [t],
	passo o model que eu quiser
*/
func NewBaseService[t any](repo interfaces.RepositoryCrud[t]) *BaseService[t] {

	return &BaseService[t]{
		Repo: repo,
	}

}


/*
	metados vindos do serviceCrud, contendo todos os metados que meus services irão usar

	"func (b *BaseService[t])" aqui eu posso trocar o "BaseService[t]" por qualquer service que eu criar, que nao ira ter alteração 
	"b" vai significar qualquer services que eu criar, como categoria, produto, ou status
*/
// Delete implements interfaces.ServiceCrud.
func (b *BaseService[t]) Delete(entity t) {
	log.Println("INFO: funcão Delete chamada")

	err:= b.Repo.Delete(entity)
	if err != nil {

		fmt.Errorf("erro em Delete: %v",err)

	}
}

// Save implements interfaces.ServiceCrud.
func (b *BaseService[t]) Save(entity t) {
	log.Println("INFO: função Save chamada")
	err:=b.Repo.Create(entity)
	if err != nil {
		fmt.Errorf("erro em Save: %v", err)
	}
}

// SearchAll implements interfaces.ServiceCrud.
func (b *BaseService[t]) SearchAll() {
	
	log.Println("INFO: função SearchAll chamada")
	_,err:=b.Repo.FindAll()
	if err != nil {
		fmt.Errorf("erro em SearchAll: %v", err)
	}
}

// SearchID implements interfaces.ServiceCrud.
func (b *BaseService[t]) SearchID(id int) {
	log.Println("INFO: função SearchAll chamada")
	_,err:=b.Repo.FindById(id)
		if err != nil {
		fmt.Errorf("erro em SearchID: %v", err)
	}
}

// Update implements interfaces.ServiceCrud.
func (b *BaseService[t]) Update(entity t) {
	log.Println("INFO: função Update chamada")
	err:=b.Repo.Update(entity)
			if err != nil {
		fmt.Errorf("erro em Update: %v", err)
	}


}


