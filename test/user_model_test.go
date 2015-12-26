package keyservicetest

import (
	"keyservice/models"
	"testing"
	"fmt"

	. "github.com/franela/goblin"
)

func TestUserModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("UserModel", func() {
		fixtures := new(Fixtures)

		g.It("should create a user model", func() {
			user := new(models.User)

			// g.Assert(user.GetDOI()).IsTrue()
			doi := user.GetDOI()
			g.Assert(doi.GetId() == "").IsTrue()
			g.Assert(user.GetUsername()).Equal("")
			g.Assert(user.GetFullname()).Equal("")
			g.Assert(user.GetEmail()).Equal("")
			g.Assert(user.GetSMS()).Equal("")
			g.Assert(user.GetStatus()).Equal("")
		})

		g.It("should validate a user struct", func() {
			user := fixtures.CreateUserModel()

			errs, ok := user.Validate()

			g.Assert(ok).Equal(true)
			g.Assert(len(errs)).Equal(0)

			user.SetStatus("flarb")
			errs, ok = user.Validate()

			fmt.Sprintf("%v\n", user )

			g.Assert(ok).Equal(false)
			g.Assert(len(errs)).Equal(1)
		})

		g.It("should create a map of a user struct", func() {
			user := fixtures.CreateUserModel()

			hash := user.ToMap()

			fmt.Sprintln( hash )

			doi := user.GetDOI()
			g.Assert(hash["id"].(string)).Equal(doi.GetId())

			g.Assert(hash["username"].(string)).Equal(user.GetUsername())
			g.Assert(hash["fullname"].(string)).Equal(user.GetFullname())
			g.Assert(hash["email"].(string)).Equal(user.GetEmail())
			g.Assert(hash["sms"].(string)).Equal(user.GetSMS())
			g.Assert(hash["status"].(string)).Equal(user.GetStatus())

		})

		g.It("should populate a user object from map", func() {
			hash := fixtures.CreateUserMap()

			user := new(models.User)
			if err := user.FromMap( hash ); err != nil {
				fmt.Println(err)
			}

			doi := user.GetDOI()
			g.Assert(doi.GetId()).Equal(hash["id"].(string))

			g.Assert(user.GetUsername()).Equal(hash["username"].(string))
			g.Assert(user.GetFullname()).Equal(hash["fullname"].(string))
			g.Assert(user.GetEmail()).Equal(hash["email"].(string))
			g.Assert(user.GetSMS()).Equal(hash["sms"].(string))
			g.Assert(user.GetStatus()).Equal(hash["status"].(string))
		})

		g.It("should create a json string from the populated user model", func() {
			user := fixtures.CreateUserModel()
			user.SetStatus(models.ModelStatus.Banned)
			user.UpdateVersion()

			user.SetFullName("John Doe")

			json, err := models.MapToJSON( user.ToMap() )

			if err != nil {
				fmt.Println( string(json) )
			}

			g.Assert(err == nil).IsTrue()
			g.Assert(json != nil).IsTrue()
		})

		g.It("should create a user map from the json object", func() {
			model := fixtures.CreateUserModel()
			model.SetStatus(models.ModelStatus.Banned)
			model.UpdateVersion()

			model.SetFullName("John Doe")

			json, err := models.MapToJSON( model.ToMap() )
			g.Assert(err == nil).IsTrue()

			hash, err := models.MapFromJSON( json )
			g.Assert(err == nil).IsTrue()

			user := new(models.User)
			err = user.FromMap( hash )
			
			g.Assert(err).Equal(nil)

		})
	})
}
