package keyservicetest

import (
	"fmt"
	"keyservice"
	"testing"

	"reflect"

	. "github.com/franela/goblin"
)

func TestDataSource(t *testing.T) {
	g := Goblin(t)

	g.Describe("DataSource", func() {
		fixtures := new(Fixtures)

		g.It("should create a cache only data source", func() {
			dataSource := keyservice.NewCachedDataSource(nil)

			g.Assert(fmt.Sprintf("%T", dataSource)).Equal("keyservice.DataSource")

			g.Assert(dataSource.GetCacheLen()).Equal(0)
			g.Assert(dataSource.GetCache() != nil).IsTrue()
		})

		g.It("should create a standard data source with primary redis client", func() {
			fixtures.CreateUserModel()
			client := keyservice.GetPrimaryClient()

			pong, err := client.Ping().Result()
			g.Assert(err).Equal(nil)
			g.Assert(pong).Equal("PONG")
		})

		g.It("should create a standard data source with secondary redis client", func() {
			fixtures.CreateUserModel()
			client := keyservice.GetSecondaryClient()

			pong, err := client.Ping().Result()
			g.Assert(err).Equal(nil)
			g.Assert(pong).Equal("PONG")
		})

		g.It("should set a known data model with a datasource and cache", func() {
			client := keyservice.GetPrimaryClient()
			dataSource := keyservice.NewCachedDataSource(client)

			ref := fixtures.CreateUserModel()
			key := "User:" + ref.GetDOI().GetId()
			err := dataSource.Set(key, ref)

			g.Assert(err).Equal(nil)

			// read it back directly from the database
			if str, err := client.Get(key).Result(); err == nil {
				g.Assert(err).Equal(nil)
				g.Assert(str != "").IsTrue()
				// should be an object

				t := reflect.TypeOf(str)

				g.Assert(t.Name()).Equal("string")

				v := reflect.ValueOf(str)
				g.Assert(v.Kind()).Equal(reflect.String)

				obj, err := ref.FromJSON([]byte(str))
				g.Assert(err).Equal(nil)
				t = reflect.TypeOf(obj)

				g.Assert(t.Name()).Equal("User")

				user, ok := obj.(keyservice.User)
				g.Assert(ok).IsTrue()
				g.Assert(user.GetDOI().GetId()).Equal(ref.GetDOI().GetId())
				g.Assert(user.GetUsername()).Equal(ref.GetUsername())

			} else {
				g.Assert(false).IsTrue()
			}

			// read back from cache
			if obj, err := dataSource.Get(key); err == nil {
				g.Assert(err).Equal(nil)
				t := reflect.TypeOf(obj)
				g.Assert(t.Name()).Equal("User")

				user, ok := obj.(keyservice.User)
				g.Assert(ok).IsTrue()
				g.Assert(user.GetDOI().GetId()).Equal(ref.GetDOI().GetId())
				g.Assert(user.GetUsername()).Equal(ref.GetUsername())
			} else {
				fmt.Println(err)
				g.Assert(false).IsTrue()
			}

		})

		g.It("should remove a known entity from database", func() {
			client := keyservice.GetPrimaryClient()
			dataSource := keyservice.NewCachedDataSource(client)

			ref := fixtures.CreateUserModel()
			key := "User:" + ref.GetDOI().GetId()
			err := dataSource.Set(key, ref)
			g.Assert(err).Equal(nil)

			obj := dataSource.Delete(key)
			g.Assert(obj != nil).IsTrue()

			t := reflect.TypeOf(obj)
			g.Assert(t.Name()).Equal("User")

			if user, err := dataSource.Get(key); err == nil {
				g.Assert(err).Equal(nil)
				g.Assert(user).Equal(nil)
			}
		})
	})
}
