package keyservicetest

import (
	"keyservice/models"
	"testing"
    "time"
    "fmt"

	. "github.com/franela/goblin"
)

func CreateDocumentIdentifierMap() map[string]interface{} {
    doi := models.NewDocumentIdentifier()

    return doi.ToMap()
}

func TestDocumentIdentifierModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("DocumentIdentifierModel", func() {
		g.It("should create a user model", func() {
			doi := new(models.DocumentIdentifier)

			g.Assert(doi != nil).IsTrue()
            g.Assert(doi.GetId()).Equal("")
		})

        g.It("should create a new model id", func() {
            id := models.NewModelId()

            g.Assert(len(id)).Equal( 32 )
            // println( id )

        })

        g.It("should create a new valid instance", func() {
            now := time.Now().UTC()
            doi := models.NewDocumentIdentifier()

            g.Assert(len(doi.GetId())).Equal( 32 )
            // fmt.Println( doi.GetDateCreated().Sub(now).Seconds() )
            g.Assert(doi.GetDateCreated().Sub(now).Seconds() < 0.001 ).IsTrue()
            g.Assert(doi.GetLastUpdated().Sub(now).Seconds() < 0.001 ).IsTrue()
            g.Assert(doi.GetVersion()).Equal( int64(0) )
        })

        g.It("should create a map of doi values", func() {
            doi := models.NewDocumentIdentifier()

            hash := doi.ToMap()

            fmt.Sprintln( hash )

            g.Assert(hash["id"].(string)).Equal(doi.GetId())

			vers := int64( hash["version"].(float64) )
            g.Assert(vers).Equal(doi.GetVersion())
        })

        g.It("should create a doi from a compatible map", func() {
            hash := models.NewDocumentIdentifier().ToMap()
			dflt := (*new(time.Time))

            dateCreated, _ := models.ParseJSONDate( hash, "dateCreated", dflt )
            lastUpdated, _ := models.ParseJSONDate( hash, "lastUpdated", dflt )
            version := int64( hash["version"].(float64) )

            doi := new(models.DocumentIdentifier)
            doi.FromMap( hash )

            g.Assert(hash["id"].(string)).Equal(doi.GetId())

            json, err := doi.GetDateCreated().MarshalJSON()
            g.Assert(err == nil).IsTrue()
            fmt.Sprintf("%s\n", json)

            g.Assert(doi.GetDateCreated()).Equal(dateCreated)
            g.Assert(doi.GetLastUpdated()).Equal(lastUpdated)
            g.Assert(doi.GetVersion()).Equal(version)
        })
	})
}
