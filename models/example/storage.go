package example

/*
All methods to save data to DB
 */
type ExampleStorage interface {
	DBInit() error
	Create(item ExampleModel) (int, error)
	Get(id int) (ExampleModel, error)
	Update(item ExampleModel) error
	Delete(id int) error
}