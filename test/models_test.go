package keyservicetest

import (

	"keyservice/models"
	"testing"
	// "fmt"


	. "github.com/franela/goblin"
)

func TestModels(t *testing.T) {
	g := Goblin(t)

	g.Describe("Models", func() {
		g.It("should create a document identifier", func() {
			doi := new(models.DocumentIdentifier)

			g.Assert(doi != nil).IsTrue()
		})
	})
}
