package keyservicetest

import (
	"keyservice"
	"testing"
	"time"
	// "fmt"

	. "github.com/franela/goblin"
)

func TestSessions(t *testing.T) {
	g := Goblin(t)

	g.Describe("Sessions", func() {
		expires := time.Now().Unix() - 1

		g.It("should have a valid session map", func() {
			sessions := keyservice.GetSessions()

			g.Assert(sessions != nil).IsTrue()
			g.Assert(sessions.Len() >= 0).IsTrue()
		})

		g.It("should create a sessions object", func() {
			sessions := keyservice.GetSessions()
			count := sessions.Len()

			// user := keyservice.CreateUser( "de306daed86b4bd5bd25dbaf59122897", "d6a0c822-8311-4b67-8024-eb7b5e27dd4e", "6bcfd43ddab34b3a88c38a8aeac7ed82" )
			session := keyservice.CreateSession(expires) // , &user)

			g.Assert(sessions.Len()).Equal(count + 1)

			g.Assert(session != nil).IsTrue()
			// sessionUser := session.GetUser()
			// ssid := sessionUser.GetSession()

			// g.Assert(len(sessionUser.GetSession())).Equal(36)
			// g.Assert(sessionUser.GetId()).Equal("de306daed86b4bd5bd25dbaf59122897")
			// g.Assert(len(ssid)).Equal(36)
		})

		g.It("should purge expired sessions", func() {
			keyservice.PurgeAllSessions()
			sessions := keyservice.GetSessions()

			g.Assert(sessions.Len()).Equal(0)

			for sessions.Len() < 10 {
				keyservice.CreateSession(expires)
			}

			g.Assert(sessions.Len()).Equal(10)
			keyservice.PurgeExpiredSessions()
			g.Assert(sessions.Len()).Equal(0)

		})

		g.It("should validate a new session", func() {
			session := keyservice.CreateSession(0) // , nil)
			// user := session.GetUser()
			ssid := session.GetSSID()

			// fmt.Println("ssid: ", ssid)

			// g.Assert( user == nil ).IsTrue()
			g.Assert(keyservice.ValidateSession(ssid)).IsTrue()

			keyservice.RemoveSession(ssid)

			g.Assert(keyservice.ValidateSession(ssid)).Equal(false)
		})

		g.It("should return a session object from it's ssid", func() {
			// user := keyservice.CreateUser("de306daed86b4bd5bd25dbaf59122897", "", "orgid")
			session := keyservice.CreateSession(0) // , &user)

			// u := session.GetUser()
			ssid := session.GetSSID()
			exp := session.GetExpires()

			ss, ok := keyservice.FindSession(ssid)

			g.Assert(ok).IsTrue("should be ok")
			g.Assert(ss != nil).IsTrue("should not be nil")

			g.Assert(exp >= time.Now().Unix()).IsTrue("expire time should be in future")
			// g.Assert(u.GetId()).Equal("de306daed86b4bd5bd25dbaf59122897")

			keyservice.RemoveSession(ssid)
			ss, ok = keyservice.FindSession(ssid)

			g.Assert(ss == nil).IsTrue()
			g.Assert(ok).Equal(false)
		})
	})
}
