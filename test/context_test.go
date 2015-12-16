package keyservicetest

import (
	"fmt"
	"keyservice"
	"strings"
	"testing"

	"github.com/darrylwest/cassava-logger/logger"

	. "github.com/franela/goblin"
)

func TestContext(t *testing.T) {
	g := Goblin(t)

	g.Describe("Context", func() {
		log := func() *logger.Logger {
			ctx := keyservice.NewContextForEnvironment("test")
			return ctx.CreateLogger()
		}()

		g.It("should create a context struct", func() {
			log.Info("create default context struct")
			ctx := new(keyservice.Context)

			g.Assert(ctx.GetShutdownPort()).Equal(0)
		})

		g.It("should return a three part version string", func() {
			version := keyservice.Version()

			g.Assert(version != "")

			parts := strings.Split(version, ".")
			g.Assert(len(parts)).Equal(3)
		})

		g.It("should return true when env is production", func() {
			g.Assert(keyservice.IsProduction("test")).Equal(false)
			g.Assert(keyservice.IsProduction("development")).Equal(false)
			g.Assert(keyservice.IsProduction("staging")).Equal(false)
			g.Assert(keyservice.IsProduction("production")).Equal(true)
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

			workFolder, _ := hash["workFolder"]
			g.Assert(strings.HasSuffix(fmt.Sprintf("%v", workFolder), ".keyservice")).Equal(true)

			configFile, _ := hash["configFile"]
			g.Assert(strings.HasSuffix(fmt.Sprintf("%v", configFile), "config.json")).Equal(true)
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
