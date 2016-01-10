package keyservice

import (
	"fmt"
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

func (dao UserDao) Save(user User) (User, error) {
	user.UpdateVersion()

	key := dao.CreateDomainKey(user.GetDOI().GetId())
	err := dao.dataSource.Set(key, user)

	return user, err
}

func (dao UserDao) Query() ([]User, error) {
	var list []User
	return list, fmt.Errorf(NotImplementedYet, "query")
}

// returns user and nil error if found, else returns error
func (dao UserDao) FindById(id string) (User, error) {
	var user User
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

func (dao UserDao) convertObject(obj interface{}) (User, error) {
	var user User

	switch v := obj.(type) {
	case string:
		u, err := user.FromJSON([]byte(v))
		if err != nil {
			return user, err
		}

		return dao.convertObject(u)
	case User:
		return v, nil
	default:
		return user, fmt.Errorf("could not convert type: %v", v)
	}

}
