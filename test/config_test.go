package keyservicetest

import (
	// "fmt"
	"keyservice"
	// "strings"
	"testing"

	"github.com/darrylwest/cassava-logger/logger"

	. "github.com/franela/goblin"
)

func createConfigJson() []byte {
	json := []byte(`{
		"name":"KeyServiceTestConfig",
		"appkey":"669a3a9db3f2456f9e1d5ffe9b13b340",
		"baseURI":"KeyService",
		"primaryRedis":{"addr":"localhost:8443","password":"flarb","db":0},
		"secondaryRedis":{"addr":"localhost:8444","password":"flarb","db":0}
	}`)

	return json
}

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

		g.It("should parse a valid config json blob and return a config struct", func() {
			config, err := keyservice.ParseConfig( createConfigJson() )

			g.Assert( err == nil ).Equal( true )
			g.Assert( config != nil ).Equal( true )
		})

		g.It("should read external configuration file", func() {
			file := "test-config.json"

			config, err := keyservice.ReadConfig( file )

			g.Assert( err == nil ).Equal( true )
			g.Assert( config != nil ).Equal( true )

		})

		g.It("should return error if config file is not found", func() {
			file := "bad-file-name.json"

			config, err := keyservice.ReadConfig( file )
			g.Assert( err != nil ).Equal( true )
			g.Assert( config == nil ).Equal( true )
		})
    })
}
