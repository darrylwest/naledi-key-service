package keyservicetest

import (
	"fmt"
	"keyservice/models"
	"testing"

	. "github.com/franela/goblin"
)

func TestAccessKeyModel(t *testing.T) {
	g := Goblin(t)

	fixtures := new(Fixtures)

	g.Describe("AccessKeyModel", func() {

		g.It("should create an access key model", func() {
			ak := new(models.AccessKey)
			g.Assert(ak != nil).IsTrue()
		})

		g.It("should create an access key model from map", func() {
			mp := fixtures.CreateAccessKeyMap()

			fmt.Sprintln(mp)

			ak := new(models.AccessKey)
			err := ak.FromMap(mp)
			g.Assert(err).Equal(nil)

			g.Assert(ak.GetId()).Equal(mp["id"].(string))
		})

		g.It("should create a json object from access key map", func() {
			mp := fixtures.CreateAccessKeyMap()

			json, err := models.MapToJSON(mp)
			g.Assert(err).Equal(nil)
			g.Assert(json != nil).IsTrue()
		})

		g.It("should create a map object from populated model", func() {
			ak := models.NewAccessKey(models.NewModelId(), []byte("this is a test"))

			mp := ak.ToMap()
			g.Assert(mp != nil).IsTrue()
			g.Assert(mp["id"].(string)).Equal(ak.GetId())
		})
	})
}
