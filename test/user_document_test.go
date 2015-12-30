package keyservicetest

import (
	"fmt"
	"keyservice/models"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestUserDocumentModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserDocumentModel", func() {
		fixtures := new(Fixtures)

		g.It("should create a new user document model", func() {
			udoc := new(models.UserDocument)
			model := fixtures.CreateUserDocumentModel(fixtures.CreateUserModel())

			fmt.Sprintln(udoc, model)

		})

		g.It("should update it's version and last updated date", func() {
			udoc := fixtures.CreateUserDocumentModel(fixtures.CreateUserModel())

			v := udoc.UpdateVersion()

			doi := udoc.GetDOI()
			g.Assert(doi.GetVersion()).Equal(v)
			g.Assert(v).Equal(int64(1))
		})

		g.It("should create a json string from the populated user document model", func() {
			user := fixtures.CreateUserModel()
			doc := fixtures.CreateUserDocumentModel(user)
			doc.SetStatus(models.ModelStatus.Expired)
			doc.UpdateVersion()

			json, err := models.MapToJSON(doc.ToMap())

			if err != nil {
				fmt.Println(string(json))
			}

			// fmt.Println( string(json) )

			g.Assert(err == nil).IsTrue()
			g.Assert(json != nil).IsTrue()
		})

		g.It("should create a user document object from the hash map", func() {
			hash := fixtures.CreateUserDocumentMap()
			// fmt.Printf("%v\n", hash)

			doc := new(models.UserDocument)
			if model, err := doc.FromMap(hash); err == nil {
				v, ok := model.(models.UserDocument)
				g.Assert(ok).IsTrue()
				doc = &v
			}

			// fmt.Printf("%v\n", doc)

			g.Assert(doc.GetStatus()).Equal(models.ModelStatus.Valid)

			doi := doc.GetDOI()
			g.Assert(doi.GetId()).Equal(hash["id"].(string))

			dflt := time.Now()
			dt, err := models.ParseJSONDate(hash, "dateCreated", dflt)
			// fmt.Println(doi.GetDateCreated(), dt)
			g.Assert(err).Equal(nil)
			g.Assert(doi.GetDateCreated()).Equal(dt)

			dt, err = models.ParseJSONDate(hash, "lastUpdated", dflt)
			// fmt.Println(doi.GetDateCreated(), dt)
			g.Assert(err).Equal(nil)
			g.Assert(doi.GetLastUpdated()).Equal(dt)

			// g.Assert(doi.GetLastUpdated()).Equal(hash["lastUpdated"].(string))
			g.Assert(doi.GetVersion()).Equal(int64(hash["version"].(float64)))
		})
	})
}
