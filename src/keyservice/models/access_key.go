package models

import (
    "errors"
)

type AccessKey struct {
	id  string  // prefixed and hashed, eg BoxKey:209eca2d...
	key []byte  // encrypted with private local key
}

func NewAccessKey(id string, key []byte) *AccessKey {
    ak := new(AccessKey)

    ak.id = id
    ak.key = key

    return ak
}

func (ak *AccessKey) GetId() string {
    return ak.id
}

func (ak *AccessKey) ToMap() map[string]interface{} {
    mp := map[string]interface{} {
        "id": ak.id,
        "key": string(ak.key),
    }

    return mp
}

func (ak *AccessKey) FromMap(v map[string]interface{}) error {

    if id, ok := v["id"].(string); ok {
        ak.id = id
    } else {
        return errors.New("could not find access key id")
    }

    if key, ok := v["key"].(string); ok {
        ak.key = []byte( key )
    } else {
        return errors.New("could not get key")
    }

    return nil
}
