package dao

import (
	"keyservice/models"
	"sync"
)

type Cache struct {
	values map[string]models.DataModelType
	sync.RWMutex
}

func NewCache() *Cache {
	cache := Cache{
		values: make(map[string]models.DataModelType),
	}

	return &cache
}

func (c *Cache) GetValues() map[string]models.DataModelType {
	return c.values
}

func (c *Cache) Len() int {
	return len(c.values)
}

func (c *Cache) Get(key string) models.DataModelType {
	c.RLock()
	value := c.values[key]
	c.RUnlock()

	// fmt.Println( value )

	return value
}

func (c *Cache) Set(key string, value models.DataModelType) error {
	c.Lock()
	c.values[key] = value
	c.Unlock()

	// fmt.Println(key, value)

	return nil
}

func (c *Cache) Delete(key string) models.DataModelType {
	if value, ok := c.values[key]; ok {
		c.Lock()
		delete(c.values, key)
		c.Unlock()

		return value
	}

	return nil
}
