package dao

import (
	"fmt"
	"keyservice"
	"strings"
)

type UserDao struct {
	dataSource DataSource
	prefix     string
}

// the context contains keys to access primary and secondary data sources
func CreateUserDao(ds DataSource) UserDao {
	dao := UserDao{
		dataSource: ds,
		prefix:     "User:",
	}

	return dao
}

func (dao UserDao) GetPrefix() string {
	return dao.prefix
}

func (dao UserDao) CreateDomainKey(key string) string {
	if strings.HasPrefix(key, dao.prefix) {
		return key
	} else {
		return dao.prefix + key
	}
}

func (dao UserDao) Save(user keyservice.User) (keyservice.User, error) {
	user.UpdateVersion()

	key := dao.CreateDomainKey(user.GetDOI().GetId())
	err := dao.dataSource.Set(key, user)

	return user, err
}

func (dao UserDao) Query() ([]keyservice.User, error) {
	var list []keyservice.User
	return list, fmt.Errorf(NotImplementedYet, "query")
}

// returns user and nil error if found, else returns error
func (dao UserDao) FindById(id string) (keyservice.User, error) {
	var user keyservice.User
	key := dao.CreateDomainKey(id)

	obj, err := dao.dataSource.Get(key)

	if err != nil {
		return user, err
	}

	if obj != nil {
		return dao.convertObject(obj)
	}

	return user, fmt.Errorf(NotFound, "user", id)
}

func (dao UserDao) convertObject(obj interface{}) (keyservice.User, error) {
	var user keyservice.User

	switch v := obj.(type) {
	case string:
		u, err := user.FromJSON([]byte(v))
		if err != nil {
			return user, err
		}

		return dao.convertObject(u)
	case keyservice.User:
		return v, nil
	default:
		return user, fmt.Errorf("could not convert type: %v", v)
	}

}
