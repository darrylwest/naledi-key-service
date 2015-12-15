package keyservicetest

import (
	"keyservice"
	"testing"

	. "github.com/franela/goblin"
)

func TestContext(t *testing.T) {
	g := Goblin(t)

	g.Describe("Context", func() {
		keyservice.CreateLogger(keyservice.NewContextForEnvironment("test"))

		g.It("should create a context struct", func() {
			ctx := new(keyservice.Context)

			g.Assert(ctx.GetShutdownPort()).Equal(0)
		})
	})
}
