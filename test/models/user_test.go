package modelstest

import (
    "keyservice/models"
    "testing"

    . "github.com/franela/goblin"
)

func TestUserModel(t *testing.T) {
    g := Goblin(t)

    g.Describe("User", func() {
        g.It("should create a user model")
    })
}
