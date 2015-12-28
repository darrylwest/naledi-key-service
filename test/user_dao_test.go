package keyservicetest

import (
	"keyservice/dao"
	"keyservice/models"
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

			g.Assert(ds.GetCacheLen()).Equal(0)
			val, err := userDao.FindById("mykey")
			g.Assert(err).Equal(nil)
			g.Assert(val).Equal((*models.User)(nil))
		})

		g.It("should save a user model and update last updated and version", func() {
			ds := dao.NewCachedDataSource(nil)
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
		})
	})
}
