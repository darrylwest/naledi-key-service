package keyservicetest

import (
    "testing"
    "github.com/darrylwest/cassava-logger/logger"

    "keyservice"

    . "github.com/franela/goblin"
)

func TestCrypto(t *testing.T) {
    g := Goblin(t)

    g.Describe("Crypto", func() {
        log := func() *logger.Logger {
			ctx := keyservice.NewContextForEnvironment("test")
			return ctx.CreateLogger()
		}()

        g.It("should generate a standard symmetric key", func() {
            log.Info("symmetric key generation")

            key, err := keyservice.GenerateSymmetricKey()

            g.Assert(err == nil).IsTrue()
            g.Assert(key != nil).IsTrue()
            g.Assert(len(key)).Equal(32)
        })

        g.It("should generate a standard nonce", func() {
            log.Info("standard nonce generation")

            key, err := keyservice.GenerateNonce()

            g.Assert(err == nil).IsTrue()
            g.Assert(key != nil).IsTrue()
            g.Assert(len(key)).Equal(24)
        })


    })
}
