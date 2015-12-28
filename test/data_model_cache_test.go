package keyservicetest

import (
	"keyservice/dao"
	"keyservice/models"
	"testing"

	. "github.com/franela/goblin"
)

func TestCache(t *testing.T) {
	g := Goblin(t)

	fixtures := new(Fixtures)

	g.Describe("Cache", func() {
		g.It("should create a cache instance", func() {
			cache := dao.NewDataModelCache()

			g.Assert(cache.Len()).Equal(0)
		})

		g.It("should get/set a model", func() {
			cache := dao.NewDataModelCache()

			ref := fixtures.CreateUserModel()
			key := ref.GetDOI().GetId()

			cache.Set(key, ref)
			g.Assert(cache.Len()).Equal(1)

			model := cache.Get(key)

			g.Assert(cache.Len()).Equal(1)
			g.Assert(model != nil).IsTrue()
			user := model.(models.User)
			g.Assert(user.GetDOI()).Equal(ref.GetDOI())
			g.Assert(user.GetUsername()).Equal(ref.GetUsername())

			ref1 := fixtures.CreateUserDocumentModel(user)
			key = ref1.GetDOI().GetId()
			cache.Set(key, ref1)
			g.Assert(cache.Len()).Equal(2)

			udoc := cache.Get(key)
			g.Assert(cache.Len()).Equal(2)
			g.Assert(udoc != nil).IsTrue()
		})
	})
}
