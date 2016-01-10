package keyservicetest

import (
	"fmt"
	"keyservice"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestUserDocumentModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserDocumentModel", func() {
		fixtures := new(Fixtures)

		g.It("should create a new user document model", func() {
			udoc := new(keyservice.UserDocument)
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
			doc.SetStatus(keyservice.ModelStatus.Expired)
			doc.UpdateVersion()

			json, err := keyservice.MapToJSON(doc.ToMap())

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

			doc := new(keyservice.UserDocument)
			if model, err := doc.FromMap(hash); err == nil {
				v, ok := model.(keyservice.UserDocument)
				g.Assert(ok).IsTrue()
				doc = &v
			}

			// fmt.Printf("%v\n", doc)

			g.Assert(doc.GetStatus()).Equal(keyservice.ModelStatus.Valid)

			doi := doc.GetDOI()
			g.Assert(doi.GetId()).Equal(hash["id"].(string))

			dflt := time.Now()
			dt, err := keyservice.ParseJSONDate(hash, "dateCreated", dflt)
			// fmt.Println(doi.GetDateCreated(), dt)
			g.Assert(err).Equal(nil)
			g.Assert(doi.GetDateCreated()).Equal(dt)

			dt, err = keyservice.ParseJSONDate(hash, "lastUpdated", dflt)
			// fmt.Println(doi.GetDateCreated(), dt)
			g.Assert(err).Equal(nil)
			g.Assert(doi.GetLastUpdated()).Equal(dt)

			// g.Assert(doi.GetLastUpdated()).Equal(hash["lastUpdated"].(string))
			g.Assert(doi.GetVersion()).Equal(int64(hash["version"].(float64)))
		})
	})
}
