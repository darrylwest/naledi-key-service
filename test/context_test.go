package keyservicetest

import (
	"keyservice"
	"testing"

	. "github.com/franela/goblin"
)

func TestContext(t *testing.T) {
	g := Goblin(t)

	g.Describe("Context", func() {
		log := keyservice.CreateLogger(keyservice.NewContextForEnvironment("test"))

		g.It("should create a context struct", func() {
            log.Info("create default context struct")
			ctx := new(keyservice.Context)

			g.Assert(ctx.GetShutdownPort()).Equal(0)
		})

        g.It("should create a context struct with defaults set", func() {
			ctx := keyservice.NewDefaultContext()

			g.Assert(ctx.GetShutdownPort()).Equal(9009)

			hash := ctx.ToMap()

			g.Assert(hash != nil)

			if value, ok := hash["webroot"]; ok {
				g.Assert(value).Equal("public")
			}

			g.Assert(hash["baseport"]).Equal(9001)
			g.Assert(hash["shutdownPort"]).Equal(9009)
			g.Assert(hash["serverCount"]).Equal(2)
		})

        g.It("should create context from args", func() {
			ctx := keyservice.ParseArgs()

			g.Assert(ctx.GetShutdownPort()).Equal(9009)

			hash := ctx.ToMap()

			g.Assert(hash["baseport"]).Equal(9001)
			g.Assert(hash["shutdownPort"]).Equal(9009)
			g.Assert(hash["serverCount"]).Equal(2)

		})
	})
}
