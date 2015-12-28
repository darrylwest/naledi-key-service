package keyservicetest

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/box"
	"testing"
	// "fmt"

	"keyservice"

	. "github.com/franela/goblin"
)

func TestCrypto(t *testing.T) {
	g := Goblin(t)

	g.Describe("Crypto", func() {
		spub := "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"
		spriv := "1d4f58f4f1e40c72dc695836902119ac553b84693904efac931731ae2ea27b48"
		plainTextMessage := []byte("this is a standard text message with some length to it that will be encrypted.  maybe")

		g.It("should generate a standard symmetric key", func() {
			key, err := keyservice.GenerateSymmetricKey()

			g.Assert(err == nil).IsTrue()
			g.Assert(key != nil).IsTrue()
			g.Assert(len(key)).Equal(32)
		})

		g.It("should generate a standard nonce", func() {
			key, err := keyservice.GenerateNonce()

			g.Assert(err == nil).IsTrue()
			g.Assert(key != nil).IsTrue()
			g.Assert(len(key)).Equal(24)
		})

		g.It("should symmetrically encrypt a plain text string", func() {
			key, _ := keyservice.GenerateSymmetricKey()

			enc, err := keyservice.EncryptSymmetric(key, plainTextMessage)

			g.Assert(err == nil).IsTrue()
			g.Assert(enc != nil).IsTrue()
		})

		g.It("should decrypt a symmetrically encrypted message", func() {
			key, _ := keyservice.GenerateSymmetricKey()

			enc, err := keyservice.EncryptSymmetric(key, plainTextMessage)

			dec, err := keyservice.DecryptSymmetric(key, enc)

			g.Assert(err == nil).IsTrue()
			g.Assert(dec).Equal(plainTextMessage)
		})

		g.It("should fail decryption if message is too short", func() {
			key, _ := keyservice.GenerateSymmetricKey()
			enc := make([]byte, 2)

			dec, err := keyservice.DecryptSymmetric(key, enc)

			g.Assert(err != nil).IsTrue()
			g.Assert(dec == nil).IsTrue()
		})

		g.It("should encrypt a message using pub/priv keys", func() {
			pub, priv, _ := box.GenerateKey(rand.Reader)

			enc, err := keyservice.EncryptBox(pub, priv, plainTextMessage)

			g.Assert(err == nil).IsTrue()
			g.Assert(len(enc) > (24 + len(plainTextMessage))).IsTrue()

		})

		g.It("should decrypt a box encrypted message", func() {
			// created by the server, and peer public key is passed to client
			srvpub, srvpriv, _ := keyservice.DecodeKeyPair(spub, spriv)
			clientpub, clientpriv, _ := box.GenerateKey(rand.Reader)

			enc, err := keyservice.EncryptBox(srvpub, clientpriv, plainTextMessage)
			g.Assert(err == nil).IsTrue()

			dec, err := keyservice.DecryptBox(clientpub, srvpriv, enc)

			g.Assert(err == nil).IsTrue()
			g.Assert(dec).Equal(plainTextMessage)
		})

		g.It("should clear a buffer's bytes to zero", func() {
			clear := new([keyservice.KeySize]byte)
			key, _ := keyservice.GenerateSymmetricKey()

			keyservice.ClearBytes(key[:])
			g.Assert(*key).Equal(*clear)
		})
	})
}
