package keyservicetest

import (
	// "fmt"
	"keyservice"
	// "strings"
	"testing"

	"github.com/darrylwest/cassava-logger/logger"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

    g.Describe("Config", func() {
        log := func() *logger.Logger {
			ctx := keyservice.NewContextForEnvironment("test")
			return ctx.CreateLogger()
		}()

        g.It("should create an instance of config", func() {
            log.Info("create a config struct")

            config := new(keyservice.Config)

            g.Assert(config != nil)
        })
    })
}
