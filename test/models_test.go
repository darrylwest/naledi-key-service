package keyservicetest

import (
	"keyservice"
	"testing"
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"
	"github.com/darrylwest/cassava-logger/logger"
	"github.com/agl/ed25519"

	. "github.com/franela/goblin"
)

var (
	spub = "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"
	spriv = "1d4f58f4f1e40c72dc695836902119ac553b84693904efac931731ae2ea27b48"
	plainTextMessage = []byte("this is a standard text message with some length to it that will be encrypted.  maybe")
)

func createMessage() *keyservice.Message {
	msg := new(keyservice.Message)

	yourpub, _, _ := keyservice.DecodeKeyPair(spub, spriv)
	mypub, mypriv, _ := box.GenerateKey( rand.Reader )
	sigpub, sigpriv, _ := ed25519.GenerateKey( rand.Reader )

	msg.SignatureKey = sigpub
	msg.Number = 5
	msg.MyKey = mypub
	msg.YourKey = yourpub

	m, _ := keyservice.EncryptBox(yourpub, mypriv, plainTextMessage)
	msg.EncryptedMessage = &m
	msg.Signature = ed25519.Sign(sigpriv, *msg.EncryptedMessage)

	return msg
}

func TestModels(t *testing.T) {
    g := Goblin(t)

	log := func() *logger.Logger {
		ctx := keyservice.NewContextForEnvironment("test")
		return ctx.CreateLogger()
	}()

    g.Describe("Models", func() {
        g.It("should create a message object and populate", func() {
			msg := createMessage()

			log.Info("%v", msg )

			g.Assert(msg.Number).Equal( 5 )
			g.Assert(msg.SignatureKey != nil).IsTrue()
			g.Assert(msg.Signature != nil).IsTrue()
			g.Assert(msg.MyKey != nil).IsTrue()
			g.Assert(msg.YourKey != nil).IsTrue()
			g.Assert(msg.EncryptedMessage != nil).IsTrue()
		})

		g.It("should encode a message to hex string", func() {
			msg := createMessage()

			str, err := msg.EncodeMessageToString()

			g.Assert(err == nil).IsTrue()
			g.Assert(len(str)).Equal( 576 )

			// TODO split string and examine sizes...

		})

		g.It("should decode a known hex string", func() {
			str := "86362ac25e94ff2c07156720d92756b2ed8b227e20c8ddcb5d65b948fe8ee26d:5d7db588042f97934b68f317bde4274e117ff74974484e206a369f76ba76efd21fc067e1b4bfcc7d5376706013f9237ce4beba65d888a75471108fffc453130f:0005:2094e5f036ec09e7fd517832a5866e2cdb72e9c1b03e5c03099941e3528f3d25:587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d:379909be3fb3b812b0315143375cca421d32f320beeaa3fc496a5e6780d12f24ebc0ba33a6528272b14cac69116743fd6fe3366338e9eccfe38ec352f8c1f3c094ab99da3a37e707153876ed097db6bbb1cadb5b1f06a88c7e2ae574cc0cc2f7ca573ec8185ffabd5e69a406d1bd6fd4fca83aa3b8a9426dd70eb3c303"

			msg, err := keyservice.DecodeMessageFromString( str )

			g.Assert(err == nil).IsTrue()
			g.Assert(msg != nil).IsTrue()
		})
    })
}
