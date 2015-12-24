package keyservicetest

import (
	"keyservice"
	"keyservice/models"
	"testing"
	"crypto/rand"
	"strings"
	"encoding/hex"
	// "fmt"

	"golang.org/x/crypto/nacl/box"
	"github.com/darrylwest/cassava-logger/logger"
	"github.com/agl/ed25519"

	. "github.com/franela/goblin"
)

var (
	spub = "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"
	spriv = "1d4f58f4f1e40c72dc695836902119ac553b84693904efac931731ae2ea27b48"
	plainTextMessage = []byte("license:this is a standard text message with some length to it that will be encrypted.  maybe")
)

func createMessage() *models.Message {
	msg := new(models.Message)

	yourpub, _, _ := keyservice.DecodeKeyPair(spub, spriv)
	mypub, mypriv, _ := box.GenerateKey( rand.Reader )
	sigpub, sigpriv, _ := ed25519.GenerateKey( rand.Reader )

	msg.SignatureKey = sigpub
	msg.Number = 1
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

			g.Assert(len(msg.Validate())).Equal( 0 )

			g.Assert(msg.Number).Equal( 1 )
			g.Assert(msg.SignatureKey != nil).IsTrue()
			g.Assert(msg.Signature != nil).IsTrue()
			g.Assert(msg.MyKey != nil).IsTrue()
			g.Assert(msg.YourKey != nil).IsTrue()
			g.Assert(msg.EncryptedMessage != nil).IsTrue()
		})

		g.It("should return error if message is invalid", func() {
			msg := new(models.Message)

			errs := msg.Validate()

			g.Assert(len(errs)).Equal( 6 )
		})

		g.It("should encode a message to hex string", func() {
			msg := createMessage()

			str, err := msg.EncodeToString()

			// fmt.Printf("%s\n\n", str)

			g.Assert(err == nil).IsTrue()
			g.Assert(len(str)).Equal( 592 )

			// TODO split string and examine sizes...
			parts := strings.Split( str, ":")
			g.Assert(len(parts)).Equal(6)

			for i, part := range parts {
				// fmt.Println( part )
				if i != 2 {
					b, err := hex.DecodeString( part )

					g.Assert(err == nil).IsTrue()
					g.Assert(b != nil).IsTrue()
				} else {
					// fmt.Println( part )
					g.Assert(part).Equal("1")
				}
			}

		})

		g.It("should decode a known hex string", func() {
			str := "e388c370815d453c3158178f549c94e3a5eb4efc1d4e6a450e9e0a2c998f801a:dc150713852bbbea8c4c202f59a6386471590c6aac557e4d784200369aeb9936038df223933b74654d3f1170cd6776a1e1899f257a109a25131d9844c19cd809:5:7e3df2eae53f770a0480214982828378e83e4e415e3e311d8fb0145d6cca275b:587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d:35ce0b60e9de37a4dcd7162e2c1fe365782992d0b7debe01c5731972cab8247a5f816a434759ccba4bcaf46860e4875e7f950db06a7657da727c0d928bf516123e13f25d958125e882f35dc4a9f7b40cc67b7237b1eaebeb03c43cd3b79bf4ffd2594ea4fd360bb7cf2ca180a9818f63af62c0883faea7ad96b46d194d"

			msg, err := models.DecodeMessageFromString( str )

			// fmt.Printf("%v\n", msg)
			g.Assert(err == nil).IsTrue()
			g.Assert(msg != nil).IsTrue()
			// g.Assert(len(msg) >= 2).IsTrue()

			g.Assert(msg.Number).Equal( 5 )
			g.Assert(msg.SignatureKey != nil).IsTrue()
			g.Assert(msg.Signature != nil).IsTrue()
			g.Assert(msg.MyKey != nil).IsTrue()
			g.Assert(msg.YourKey != nil).IsTrue()
			g.Assert(msg.EncryptedMessage != nil).IsTrue()

		})

		g.It("should reject a bad message", func() {
			str := "c370815d453c3158178f549c94e3a5eb4efc1d4e6a450e9e0a2c998f801a:dc150713852bbbea8c4c202f59a6386471590c6aac557e4d784200369aeb9936038df223933b74654d3f1170cd6776a1e1899f257a109a25131d9844c19cd809:5:7e3df2eae53f770a0480214982828378e83e4e415e3e311d8fb0145d6cca275b:587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d:35ce0b60e9de37a4dcd7162e2c1fe365782992d0b7debe01c5731972cab8247a5f816a434759ccba4bcaf46860e4875e7f950db06a7657da727c0d928bf516123e13f25d958125e882f35dc4a9f7b40cc67b7237b1eaebeb03c43cd3b79bf4ffd2594ea4fd360bb7cf2ca180a9818f63af62c0883faea7ad96b46d194d"

			msg, err := models.DecodeMessageFromString( str )

			// fmt.Printf("%v\n", msg)
			g.Assert(err != nil).IsTrue()
			g.Assert(msg == nil).IsTrue()

			str = "e388X370815d453c3158178f549c94e3a5eb4efc1d4e6a450e9e0a2c998f801a:dc150713852bbbea8c4c202f59a6386471590c6aac557e4d784200369aeb9936038df223933b74654d3f1170cd6776a1e1899f257a109a25131d9844c19cd809:5:7e3df2eae53f770a0480214982828378e83e4e415e3e311d8fb0145d6cca275b:587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d:35ce0b60e9de37a4dcd7162e2c1fe365782992d0b7debe01c5731972cab8247a5f816a434759ccba4bcaf46860e4875e7f950db06a7657da727c0d928bf516123e13f25d958125e882f35dc4a9f7b40cc67b7237b1eaebeb03c43cd3b79bf4ffd2594ea4fd360bb7cf2ca180a9818f63af62c0883faea7ad96b46d194d"

			msg, err = models.DecodeMessageFromString( str )

			// fmt.Printf("%v\n", msg)
			g.Assert(err != nil).IsTrue()
			g.Assert(msg == nil).IsTrue()
		})
    })
}
