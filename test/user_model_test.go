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

		g.It("should validate a user struct", func() {
			user := new(models.User)

			errs, ok := user.Validate()

			g.Assert(ok == false).IsTrue()
			g.Assert(len(errs) > 0).IsTrue()
		})
	})
}
