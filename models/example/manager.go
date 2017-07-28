package example

func NewExampleMan(storage ExampleStorage) Manager {
	return Manager{storage: storage}
}

/*
Use manager to get|set data
 */
type Manager struct {
	storage ExampleStorage
}

func (m Manager) Create(item ExampleModel) (int, error) {
	if err := item.Validate(); err != nil {
		return 0, err
	}
	id, err := m.storage.Create(item)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m Manager) Get(id int) (ExampleModel, error) {
	data, err := m.storage.Get(id)
	if err != nil {
	}
	return data, nil
}

func (m Manager) Update(item ExampleModel) error {
	return m.storage.Update(item)
}

func (m Manager) Delete(id int) error {
	return m.storage.Delete(id)
}
