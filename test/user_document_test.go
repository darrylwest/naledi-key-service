package keyservicetest

import (
	"keyservice/models"
	"testing"
	"fmt"

	. "github.com/franela/goblin"
)

func TestUserDocumentModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserDocumentModel", func() {
        fixtures := new(Fixtures)

        g.It("should create a new user document model", func() {
            udoc := new(models.UserDocument)
            model := fixtures.CreateUserDocumentModel(fixtures.CreateUserModel())

            fmt.Sprintln( udoc, model )

        })

        g.It("should update it's version and last updated date", func() {
            udoc := fixtures.CreateUserDocumentModel(fixtures.CreateUserModel())

            v := udoc.UpdateVersion()

            doi := udoc.GetDOI()
            g.Assert(doi.GetVersion()).Equal(v)
            g.Assert(v).Equal( int64( 1 ))
        })

        g.It("should create a json string from the populated user document model", func() {
            user := fixtures.CreateUserModel()
			doc := fixtures.CreateUserDocumentModel(user)
			doc.SetStatus(models.ModelStatus.Expired)
			doc.UpdateVersion()

			json, err := models.ModelToJson( doc.ToMap() )

			if err != nil {
				fmt.Println( string(json) )
			}

            // fmt.Println( string(json) )

			g.Assert(err == nil).IsTrue()
			g.Assert(json != nil).IsTrue()
		})

        g.It("should create a user document object from the hash map", func() {
            hash := fixtures.CreateUserDocumentMap()

            doc := new(models.UserDocument)
            doc.FromMap( hash )

            g.Assert(doc.GetStatus()).Equal(models.ModelStatus.Valid)

            doi := doc.GetDOI()
            g.Assert(doi.GetId()).Equal(hash["id"].(string))
        })
    })
}
