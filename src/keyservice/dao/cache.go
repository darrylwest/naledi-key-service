package dao

import (
	"sync"
)

type Cache struct {
	values map[string]interface{}
	sync.RWMutex
}

func NewCache() *Cache {
	cache := Cache{
		values: make(map[string]interface{}),
	}

	return &cache
}

func (c *Cache) GetValues() map[string]interface{} {
	return c.values
}

func (c *Cache) Len() int {
	return len(c.values)
}

func (c *Cache) Get(key string) interface{} {
	c.RLock()
	value := c.values[key]
	c.RUnlock()

	// fmt.Println( value )

	return value
}

func (c *Cache) Set(key, value string) error {
	c.Lock()
	c.values[key] = value
	c.Unlock()

	// fmt.Println(key, value)

	return nil
}

func (c *Cache) Delete(key string) interface{} {
	if value, ok := c.values[key]; ok {
		c.Lock()
		delete(c.values, key)
		c.Unlock()

		return value
	}

	return nil
}
