package keyservicetest

import (
	"encoding/hex"
	"keyservice"
	"testing"

	"gopkg.in/redis.v3"

	. "github.com/franela/goblin"
)

func createConfigJson() []byte {
	json := []byte(`{
		"name":"KeyServiceTestConfig",
		"appkey":"669a3a9db3f2456f9e1d5ffe9b13b340",
		"baseURI":"KeyService",
		"primaryRedisOptions":{"addr":"localhost:8443","password":"flarb","db":0},
		"secondaryRedisOptions":{"addr":"localhost:8444","password":"blarf","db":1},
		"privateLocalKey":"c2c3f31d02109189c1d22fc5c9e2ecf6bc4384a1117393d23614d1b91bed9271"
	}`)

	return json
}

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		g.It("should create an instance of config", func() {

			config := new(keyservice.Config)

			g.Assert(config != nil).Equal(true)
		})

		g.It("should parse a valid config json blob and return a config struct", func() {
			config, err := keyservice.ParseConfig(createConfigJson())

			g.Assert(err).Equal(nil)
			g.Assert(config != nil).Equal(true)

			hash := config.ToMap()

			g.Assert(hash["name"]).Equal("KeyServiceTestConfig")
			g.Assert(hash["appkey"]).Equal("669a3a9db3f2456f9e1d5ffe9b13b340")
			g.Assert(config.GetAppKey()).Equal("669a3a9db3f2456f9e1d5ffe9b13b340")
			g.Assert(hash["baseURI"]).Equal("KeyService")
			g.Assert(config.GetBaseURI()).Equal("KeyService")

			g.Assert(hash["primaryRedisOptions"] != nil).Equal(true)
			g.Assert(hash["secondaryRedisOptions"] != nil).Equal(true)

			opts, ok := hash["primaryRedisOptions"].(*redis.Options)

			g.Assert(ok).Equal(true)
			g.Assert(opts.Addr).Equal("localhost:8443")
			g.Assert(opts.Password).Equal("flarb")
			g.Assert(opts.DB).Equal(int64(0))

			opts, ok = hash["secondaryRedisOptions"].(*redis.Options)

			g.Assert(ok).Equal(true)
			g.Assert(opts.Addr).Equal("localhost:8444")
			g.Assert(opts.Password).Equal("blarf")
			g.Assert(opts.DB).Equal(int64(1))

			key := config.GetPrivateLocalKey()

			g.Assert(len(key)).Equal(32)
			skey := hex.EncodeToString(key[:])
			g.Assert(skey).Equal("c2c3f31d02109189c1d22fc5c9e2ecf6bc4384a1117393d23614d1b91bed9271")
		})

		g.It("should read external configuration file", func() {
			file := "test-config.json"

			config, err := keyservice.ReadConfig(file)

			g.Assert(err == nil).Equal(true)
			g.Assert(config != nil).Equal(true)

			hash := config.ToMap()

			// fmt.Printf("%s\n", hash)

			g.Assert(hash["name"]).Equal("KeyServiceTestConfig")
			g.Assert(hash["appkey"]).Equal("c2b4d9bf-652e-4915-ab23-7a0e0e32e362")
			g.Assert(config.GetAppKey()).Equal("c2b4d9bf-652e-4915-ab23-7a0e0e32e362")
			g.Assert(hash["baseURI"]).Equal("KeyService")
			g.Assert(config.GetBaseURI()).Equal("KeyService")

			g.Assert(hash["primaryRedisOptions"] != nil).Equal(true)
			g.Assert(hash["secondaryRedisOptions"] != nil).Equal(true)

			opts, ok := hash["primaryRedisOptions"].(*redis.Options)

			g.Assert(ok).Equal(true)
			g.Assert(opts.Addr).Equal("localhost:15101")
			g.Assert(opts.Password).Equal("")
			g.Assert(opts.DB).Equal(int64(0))

			opts = config.GetPrimaryRedisOptions()
			g.Assert(opts.Addr).Equal("localhost:15101")
			g.Assert(opts.Password).Equal("")
			g.Assert(opts.DB).Equal(int64(0))

			opts, ok = hash["secondaryRedisOptions"].(*redis.Options)

			g.Assert(ok).Equal(true)
			g.Assert(opts.Addr).Equal("localhost:15102")
			g.Assert(opts.Password).Equal("")
			g.Assert(opts.DB).Equal(int64(0))

			opts = config.GetSecondaryRedisOptions()
			g.Assert(opts.Addr).Equal("localhost:15102")
			g.Assert(opts.Password).Equal("")
			g.Assert(opts.DB).Equal(int64(0))

			key := config.GetPrivateLocalKey()

			g.Assert(len(key)).Equal(32)
			skey := hex.EncodeToString(key[:])
			g.Assert(skey).Equal("c2c3f31d02109189c1d22fc5c9e2ecf6bc4384a1117393d23614d1b91bed9271")
		})

		g.It("should return error if config file is not found", func() {
			file := "bad-file-name.json"

			config, err := keyservice.ReadConfig(file)
			g.Assert(err != nil).Equal(true)
			g.Assert(config == nil).Equal(true)
		})
	})
}
