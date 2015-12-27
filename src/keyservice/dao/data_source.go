package dao

import (
    "gopkg.in/redis.v3"
    "sync"
    "fmt"
)

type DataSource struct {
    client *redis.Client
    sync.RWMutex
    cache map[string]interface{}
}

func NewDataSource(client *redis.Client, cache map[string]interface{}) DataSource {
    ds := DataSource{
        client: client,
        cache: cache,
    }

    return ds
}

func (ds *DataSource) Get(key string) (interface{}, error) {
    ds.RLock()
    value := ds.cache[key]
    ds.RUnlock()

    fmt.Println( value )

    return value, nil
}

func (ds *DataSource) Set(key, value string) error {
    ds.Lock()
    ds.cache[key] = value
    ds.Unlock()

    fmt.Println(key, value)

    return nil
}

func (ds *DataSource) Delete(key string) interface{} {
    if value, ok := ds.cache[key]; ok {
        ds.Lock()
        delete(ds.cache, key)
        ds.Unlock()

        return value
    }

    return nil
}

func (ds *DataSource) GetCache() map[string]interface{} {
    return ds.cache
}
