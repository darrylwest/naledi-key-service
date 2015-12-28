package dao

import (
    // "keyservice"
    "keyservice/models"
	// "gopkg.in/redis.v3"
    "errors"
    // "fmt"
)

var (
    NotImplementedYet = errors.New("not implemented yet")
    NotFound = errors.New("not found")
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

func (dao *UserDao) Save(user *models.User) (*models.User, error) {
    user.UpdateVersion()

    json, err := user.ToJSON()
    if err != nil {
        return nil, err
    }

    key := dao.CreateDomainKey( user.GetDOI().GetId() )

    // fmt.Println( key, string(json) )
    err = dao.dataSource.Set( key, string(json) )

    return user, err
}

func (dao *UserDao) Query() ([]*models.User, error) {
    return nil, NotImplementedYet
}

func (dao *UserDao) FindById(id string) (*models.User, error) {
    var user *models.User
    key := dao.CreateDomainKey( id )

    if val, err := dao.dataSource.Get( key ); err != nil {
        return nil, err
    } else if str, ok := val.(string); ok {
        user, err = models.UserFromJSON([]byte(str))
    }

    return user, nil
}
