package keyservicetest

import (
	"keyservice/models"
	"testing"
	"fmt"

	. "github.com/franela/goblin"
)

func TestUserDocumentModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserModel", func() {
        fixtures := new(Fixtures)

        g.It("should create a new user document model", func() {
            udoc := new(models.UserDocument)
            model := fixtures.CreateUserDocumentModel()

            fmt.Sprintln( udoc, model )

        })
    })
}
