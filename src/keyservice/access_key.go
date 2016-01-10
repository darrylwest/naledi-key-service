package keyservice

import (
	"errors"
)

type AccessKey struct {
	id  string // plain id; encryption happens at DAO
	key []byte // plain key; encrypted with private local key at DAO
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
	mp := map[string]interface{}{
		"id":  ak.id,
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
		ak.key = []byte(key)
	} else {
		return errors.New("could not get key")
	}

	return nil
}
