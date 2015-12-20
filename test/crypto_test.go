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
        plainTextMessage := []byte("this is a standard text message with some length to it that will be encrypted.  maybe")
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

        g.It("should encrypt a plain text string", func() {
            key, _ := keyservice.GenerateSymmetricKey()

            enc, err := keyservice.Encrypt(key, plainTextMessage)

            g.Assert(err == nil).IsTrue()
            g.Assert(enc != nil).IsTrue()

            log.Info("encrypted: %v", enc)
        })

        g.It("should decrypt an encrypted message", func() {
            key, _ := keyservice.GenerateSymmetricKey()

            enc, err := keyservice.Encrypt(key, plainTextMessage)

            dec, err := keyservice.Decrypt(key, enc)

            g.Assert(err == nil).IsTrue()
            g.Assert(dec).Equal(plainTextMessage)
        })

        g.It("should fail decryption if message is too short", func() {
            key, _ := keyservice.GenerateSymmetricKey()
            enc := make([]byte, 2)

            dec, err := keyservice.Decrypt(key, enc)

            log.Info("error: %s", err)
            g.Assert(err != nil).IsTrue()
            g.Assert(dec == nil).IsTrue()
        })
    })
}
