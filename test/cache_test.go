package keyservicetest

import (
	"keyservice/dao"
	"testing"

	. "github.com/franela/goblin"
)

func TestCache(t *testing.T) {
	g := Goblin(t)

	g.Describe("Cache", func() {
		g.It("should create a cache instance", func() {
			cache := dao.NewCache()

			g.Assert(cache.Len()).Equal(0)
		})
	})
}
