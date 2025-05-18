package interfaces

type CrudService interface {
	Create(name string)
	Delete(id uint)
	List()
	Update(id uint, name string)
}
