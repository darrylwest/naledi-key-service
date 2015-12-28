package dao

import (
	"errors"
	"keyservice/models"
)

var (
	NotImplementedYet = errors.New("not implemented yet")
	NotFound          = errors.New("not found")
)

type UserDao struct {
	dataSource DataSource
}

// the context contains keys to access primary and secondary data sources
func CreateUserDao(ds DataSource) UserDao {
	dao := UserDao{
		dataSource: ds,
	}

	return dao
}

func (dao *UserDao) CreateDomainKey(key string) string {
	return "User:" + key
}

func (dao *UserDao) Save(user models.User) (models.User, error) {
	user.UpdateVersion()

	key := dao.CreateDomainKey(user.GetDOI().GetId())
	err := dao.dataSource.Set(key, user)

	return user, err
}

func (dao *UserDao) Query() ([]*models.User, error) {
	var list []*models.User
	return list, NotImplementedYet
}

func (dao *UserDao) FindById(id string) (*models.User, error) {
	var user *models.User
	key := dao.CreateDomainKey(id)

	obj, err := dao.dataSource.Get(key)

	if obj != nil {
		user = obj.(*models.User)
	}

	return user, err
}
