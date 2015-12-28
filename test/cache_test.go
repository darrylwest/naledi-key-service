package keyservicetest

import (
	"keyservice/dao"
	"testing"

	. "github.com/franela/goblin"
)

func TestCache(t *testing.T) {
	g := Goblin(t)

	fixtures := new(Fixtures)

	g.Describe("Cache", func() {
		g.It("should create a cache instance", func() {
			cache := dao.NewCache()

			g.Assert(cache.Len()).Equal(0)
		})

		g.It("should get/set a model", func() {
			cache := dao.NewCache()

			ref := fixtures.CreateUserModel()
			key := ref.GetDOI().GetId()

			cache.Set(key, ref)
			g.Assert(cache.Len()).Equal(1)

			user := cache.Get(key)

			g.Assert(cache.Len()).Equal(1)
			g.Assert(user != nil).IsTrue()
		})
	})
}
