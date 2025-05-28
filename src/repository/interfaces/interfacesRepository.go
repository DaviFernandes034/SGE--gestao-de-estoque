package interfaces

//interface criada para não ficar digitando as mesmas coisas
//aqui tem todos os metados CRUD(create,read, update, delete)
//usado em todo repository
type RepositoryCrud[t any] interface {
/*  esse T any, é um tipo generico, aonde posso passar todos os models como parametro*/
	Create(entity t) error
	FindById(id int) (t, error)
	Update(entity t) (error)
	Delete(entity t) (error)
	FindAll()([]t, error)
}