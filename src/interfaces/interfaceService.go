package interfaces

//interface criada para não ficar digitando as mesmas coisas
//aqui tem todos os metados CRUD(create,read, update, delete)
//usado em todo Service
type ServiceCrud[t any] interface {
/*  esse T any, é um tipo generico, aonde posso passar todos os models como parametro*/
	Save(entity t) 
	SearchID(id int) 
	Update(entity t) 
	Delete(entity t) 
	SearchAll()
}