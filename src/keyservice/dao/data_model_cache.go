package dao

import (
	"keyservice/models"
	"sync"
)

type DataModelCache struct {
	values map[string]models.DataModelType
	sync.RWMutex
}

func NewDataModelCache() *DataModelCache {
	cache := DataModelCache{
		values: make(map[string]models.DataModelType),
	}

	return &cache
}

func (c *DataModelCache) GetValues() map[string]models.DataModelType {
	return c.values
}

func (c *DataModelCache) Len() int {
	return len(c.values)
}

func (c *DataModelCache) Get(key string) models.DataModelType {
	c.RLock()
	value := c.values[key]
	c.RUnlock()

	// fmt.Println( value )

	return value
}

func (c *DataModelCache) Set(key string, value models.DataModelType) error {
	c.Lock()
	c.values[key] = value
	c.Unlock()

	// fmt.Println(key, value)

	return nil
}

func (c *DataModelCache) Delete(key string) models.DataModelType {
	if value, ok := c.values[key]; ok {
		c.Lock()
		delete(c.values, key)
		c.Unlock()

		return value
	}

	return nil
}
