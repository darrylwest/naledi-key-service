package keyservicetest

import (
	"fmt"
	"keyservice/dao"
	"reflect"
	"strings"
	"testing"

	. "github.com/franela/goblin"
)

func TestUserDao(t *testing.T) {
	g := Goblin(t)

	fixtures := new(Fixtures)

	g.Describe("UserDao", func() {

		g.It("should create an instance of user dao", func() {
			ds := dao.NewCachedDataSource(nil)
			userDao := dao.CreateUserDao(ds)

			g.Assert(fmt.Sprintf("%T", userDao)).Equal("dao.UserDao")

			g.Assert(ds.GetCacheLen()).Equal(0)
			val, err := userDao.FindById("mykey")
			g.Assert(err != nil).IsTrue()
			g.Assert(val.GetDOI().GetId()).Equal("")
		})

		g.It("should save a user model and update last updated and version", func() {
			client := dao.GetPrimaryClient()
			client.FlushAll()
			ds := dao.NewCachedDataSource(client)
			dao := dao.CreateUserDao(ds)
			user := fixtures.CreateUserModel()

			doi := user.GetDOI()
			lastUpdated := doi.GetLastUpdated()
			version := doi.GetVersion()

			// fmt.Println( doi )

			u, err := dao.Save(user)

			// fmt.Println( u.GetDOI() )

			g.Assert(err).Equal(nil)
			g.Assert(u.GetDOI().GetId()).Equal(user.GetDOI().GetId())

			g.Assert(u.GetDOI().GetDateCreated().Equal(doi.GetDateCreated())).IsTrue()
			g.Assert(u.GetDOI().GetLastUpdated().After(lastUpdated)).IsTrue()
			g.Assert(u.GetDOI().GetVersion()).Equal(version + 1)
		})

		g.It("should create a user domain key", func() {
			ds := dao.NewCachedDataSource(nil)
			dao := dao.CreateUserDao(ds)
			user := fixtures.CreateUserModel()
			id := user.GetDOI().GetId()

			key := dao.CreateDomainKey(id)

			g.Assert(key != "").IsTrue()
			g.Assert(strings.HasPrefix(key, "User:")).IsTrue()
			g.Assert(strings.HasSuffix(key, id)).IsTrue()

			g.Assert(dao.GetPrefix()).Equal("User:")

			// test to insure prefix added only once
			str := dao.CreateDomainKey(key)
			g.Assert(str).Equal(key)
		})

		g.It("should find a known user by id", func() {
			client := dao.GetPrimaryClient()
			dataSource := dao.NewCachedDataSource(client)

			ref := fixtures.CreateUserModel()
			id := ref.GetDOI().GetId()
			key := "User:" + id
			err := dataSource.Set(key, ref)

			g.Assert(err).Equal(nil)

			dao := dao.CreateUserDao(dataSource)
			user, err := dao.FindById(id)
			g.Assert(err).Equal(nil)

			// g.Assert(user != nil).IsTrue()

			t := reflect.TypeOf(user)
			g.Assert(t.Name()).Equal("User")

			g.Assert(user.GetDOI().GetId()).Equal(id)
			if list, ok := user.Validate(); ok {
				g.Assert(len(list)).Equal(0)
			}

			client.FlushAll()
		})
	})
}
