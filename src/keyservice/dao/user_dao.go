package dao

import (
	"fmt"
	"keyservice/models"
	"strings"
)

type UserDao struct {
	dataSource DataSource
	prefix string
}

// the context contains keys to access primary and secondary data sources
func CreateUserDao(ds DataSource) UserDao {
	dao := UserDao{
		dataSource: ds,
		prefix: "User:",
	}

	return dao
}

func (dao UserDao) GetPrefix() string {
	return dao.prefix
}

func (dao UserDao) CreateDomainKey(key string) string {
	if strings.HasPrefix( key, dao.prefix ) {
		return key
	} else {
		return dao.prefix + key
	}
}

func (dao UserDao) Save(user models.User) (models.User, error) {
	user.UpdateVersion()

	key := dao.CreateDomainKey(user.GetDOI().GetId())
	err := dao.dataSource.Set(key, user)

	return user, err
}

func (dao UserDao) Query() ([]models.User, error) {
	var list []models.User
	return list, fmt.Errorf(NotImplementedYet, "query")
}

// returns user and nil error if found, else returns error
func (dao UserDao) FindById(id string) (models.User, error) {
	var user models.User
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

func (dao UserDao) convertObject(obj interface{}) (models.User, error) {
	var user models.User

	switch v := obj.(type) {
	case string:
		u, err := user.FromJSON([]byte(v))
		if err != nil {
			return user, err
		}

		return dao.convertObject(u)
	case models.User:
		return v, nil
	default:
		return user, fmt.Errorf("could not convert type: %v", v)
	}


}
