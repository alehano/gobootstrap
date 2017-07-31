package example

import (
	"github.com/alehano/gobootstrap/sys/pubsub"
	"github.com/alehano/gobootstrap/config"
	"github.com/alehano/gobootstrap/sys/memcache"
)

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
	pubsub.Publish(config.ExampleCreatedMsg, id)
	return id, nil
}

func (m Manager) Get(id int) (ExampleModel, error) {
	return m.storage.Get(id)
}

func (m Manager) GetCached(id int) (ExampleModel, error) {
	var res ExampleModel
	err := memcached.GetSetObj(config.CacheKeys.ExampleGet(id), &res, func() (interface{}, error) {
		return m.storage.Get(id)
	}, memcached.Expiration1h)
	return res, err
}

func (m Manager) Update(item ExampleModel) error {
	err := m.storage.Update(item)
	if err != nil {
		return err
	}
	pubsub.Publish(config.ExampleUpdateMsg, item.ID)
	return nil
}

func (m Manager) Delete(id int) error {
	err := m.storage.Delete(id)
	if err != nil {
		return err
	}
	pubsub.Publish(config.ExampleDeleteMsg, id)
	return nil
}
