package keyservicetest

import (
	"keyservice/models"
	"testing"

	. "github.com/franela/goblin"
)

func TestUserModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserModel", func() {
		g.It("should create a user model", func() {
			user := new(models.User)

			g.Assert(user != nil).IsTrue()
		})
	})
}
