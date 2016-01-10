package keyservicetest

import (
	"fmt"
	"keyservice"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestCache(t *testing.T) {
	g := Goblin(t)

	fixtures := new(Fixtures)

	g.Describe("Cache", func() {
		g.It("should create a cache instance", func() {
			cache := keyservice.NewDataModelCache()

			g.Assert(cache.Len()).Equal(0)
		})

		g.It("should get/set a model", func() {
			cache := keyservice.NewDataModelCache()

			ref := fixtures.CreateUserModel()
			key := ref.GetDOI().GetId()

			g.Assert(cache.Len()).Equal(0)

			cache.Set(key, ref)
			g.Assert(cache.Len()).Equal(1)

			model := cache.Get(key)

			g.Assert(cache.Len()).Equal(1)
			g.Assert(model != nil).IsTrue()
			user, ok := model.(keyservice.User)
			g.Assert(ok).IsTrue()
			g.Assert(user.GetDOI()).Equal(ref.GetDOI())
			g.Assert(user.GetUsername()).Equal(ref.GetUsername())

			ref1 := fixtures.CreateUserDocumentModel(user)
			key = ref1.GetDOI().GetId()
			cache.Set(key, ref1)
			g.Assert(cache.Len()).Equal(2)

			model = cache.Get(key)
			g.Assert(cache.Len()).Equal(2)
			udoc, ok := model.(keyservice.UserDocument)
			g.Assert(ok).IsTrue()
			g.Assert(udoc.GetDOI().GetId()).Equal(ref1.GetDOI().GetId())

			// insure that a bad cast won't panic if second parameter (ok) is used
			_, ok = model.(keyservice.User)
			g.Assert(ok).IsFalse()
		})

		g.It("should update access and cached time stamps for get/set", func() {
			cache := keyservice.NewDataModelCache()

			ref := fixtures.CreateUserModel()
			key := ref.GetDOI().GetId()

			g.Assert(cache.Len()).Equal(0)

			now := time.Now().Unix()
			cache.Set(key, ref)
			g.Assert(cache.Len()).Equal(1)

			item := cache.GetItem(key)
			g.Assert(item != nil).IsTrue()

			model, cached, accessed := item.Values()

			g.Assert(model.GetDOI().GetId()).Equal(ref.GetDOI().GetId())
			g.Assert(cached >= now).IsTrue()
			g.Assert(accessed >= now).IsTrue()

			fmt.Sprintf("%v %d %d", model, cached, accessed)
		})

		g.It("should delete a model", func() {
			cache := keyservice.NewDataModelCache()

			ref := fixtures.CreateUserModel()
			key := ref.GetDOI().GetId()

			g.Assert(cache.Len()).Equal(0)

			cache.Set(key, ref)
			g.Assert(cache.Len()).Equal(1)

			model := cache.Delete(key)
			g.Assert(cache.Len()).Equal(0)
			g.Assert(model != nil).IsTrue()

			model = cache.Delete(key)
			g.Assert(cache.Len()).Equal(0)
			g.Assert(model).Equal(nil)
		})

		g.It("should flush/clear all cached items")

	})
}
