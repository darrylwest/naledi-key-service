package keyservicetest

import (
    "fmt"
    "sync"
)

type MockRedisClient struct {
    cache map[string]interface{}
    sync.RWMutex
}

func CreateMockRedisClient() *MockRedisClient {
    client := new(MockRedisClient)

    client.cache = make(map[string]interface{})
    return client
}

func (ds *MockRedisClient) GetCache() map[string]interface{} {
    return ds.cache
}

func (ds *MockRedisClient) Get(key string) (interface{}, error) {
    ds.RLock()
    value := ds.cache[key]
    ds.RUnlock()

    fmt.Println( value )

    return value, nil
}

func (ds *MockRedisClient) Set(key, value string) error {
    ds.Lock()
    ds.cache[key] = value
    ds.Unlock()

    fmt.Println(key, value)

    return nil
}

func (ds *MockRedisClient) Delete(key string) interface{} {
    if value, ok := ds.cache[key]; ok {
        ds.Lock()
        delete(ds.cache, key)
        ds.Unlock()

        return value
    }

    return nil
}
