package example

// Interface to store data in DB
type ExampleStorage interface {
	DBInit() error
	Create(item ExampleModel) (int, error)
	Get(id int) (ExampleModel, error)
	Update(item ExampleModel) error
	Delete(id int) error
}