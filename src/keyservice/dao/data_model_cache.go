package dao

import (
	"keyservice/models"
	"sync"
	"time"
)

// CacheItemStats to speed update operations ?
type CacheItemStats struct {
	timeCached int64 // unix seconds when this item was last set
	lastAccessed int64 // unix seconds when this item was last get
	lastSavedToDb int64  // unix seconds when this item was saved to the database
}

type CacheItem struct {
	model models.DataModelType
	timeCached int64 // unix seconds when this item was last set
	lastAccessed int64 // unix seconds when this item was last get
}

func (item CacheItem) Values() (models.DataModelType, int64, int64) {
	return item.model, item.timeCached, item.lastAccessed
}

type DataModelCache struct {
	values map[string]CacheItem
	sync.RWMutex
}

func NewDataModelCache() *DataModelCache {
	cache := DataModelCache{
		values: make(map[string]CacheItem),
	}

	return &cache
}

func (c *DataModelCache) GetValues() []models.DataModelType {
	list := make([]models.DataModelType, len(c.values))

	for _,v := range c.values {
		list = append(list, v.model)
	}

	return list
}

func (c *DataModelCache) Len() int {
	return len(c.values)
}

func (c *DataModelCache) Get(key string) models.DataModelType {
	item := c.GetItem(key)
	if item != nil {
		return item.model
	} else {
		return nil
	}
}

func (c *DataModelCache) GetItem(key string) *CacheItem {
	c.RLock()
	item, ok := c.values[key]
	c.RUnlock()

	if ok {
		c.Lock()
		item.lastAccessed = time.Now().Unix()
		c.Unlock()

		return &item
	} else {
		return nil
	}
}

func (c *DataModelCache) Set(key string, value models.DataModelType) error {
	now := time.Now().Unix()
	item := CacheItem{
		model: value,
		timeCached: now,
		lastAccessed: now,
	}

	c.Lock()
	c.values[key] = item
	c.Unlock()

	// fmt.Println(key, value)

	return nil
}

func (c *DataModelCache) Delete(key string) models.DataModelType {
	if item, ok := c.values[key]; ok {
		c.Lock()
		delete(c.values, key)
		c.Unlock()

		return item.model
	}

	return nil
}

// TODO : implement this
func (c *DataModelCache) PurgeOldItems(age int64) int {
	// this method activates a crawler that continuously sweeps the cache to remove old items

	return 0
}
