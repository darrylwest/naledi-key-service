package dao

import (
	"gopkg.in/redis.v3"
	"keyservice/models"
	// "github.com/darrylwest/cassava-logger/logger"
)

type DataSource struct {
	client *redis.Client
	cache  *DataModelCache
}

// create a new datasource with optional recis client; cache is always created...
func NewDataSource(client *redis.Client) DataSource {
	ds := DataSource{}

	ds.client = client

	// make this a no-op cache
	ds.cache = NewDataModelCache()

	return ds
}

func NewCachedDataSource(client *redis.Client) DataSource {
	ds := DataSource{}

	ds.client = client
	ds.cache = NewDataModelCache()

	return ds
}

func (ds *DataSource) Get(key string) (models.DataModelType, error) {
	value := ds.cache.Get(key)

	log.Info("get: %s=%v", key, value)

	return value, nil
}

func (ds *DataSource) GetByType(key string, model models.DataModelType) (interface{}, error) {
	if ds.client == nil {
		return nil, nil
	}

	str, err := ds.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	value, err := model.FromJSON([]byte(str))
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (ds *DataSource) Set(key string, value models.DataModelType) error {
	ds.cache.Set(key, value)

	log.Info("set: %s=%v", key, value)
	if ds.client != nil {
		err := ds.client.Set(key, value, 0).Err()

		return err
	}

	return nil
}

func (ds *DataSource) Delete(key string) interface{} {
	value := ds.cache.Delete(key)

	// TODO : remove from redis
	if ds.client != nil {
		ds.client.Del(key)
	}

	return value
}

func (ds *DataSource) GetCache() *DataModelCache {
	return ds.cache
}

func (ds *DataSource) GetCacheLen() int {
	return ds.cache.Len()
}
