package keyservicetest

import (
	"fmt"
	"strings"
	"testing"

	"keyservice"

	. "github.com/franela/goblin"
)

func TestServer(t *testing.T) {
	g := Goblin(t)

	g.Describe("Server", func() {
		g.It("should do configure routing for known routes", func() {
			mux := keyservice.ConfigureRoutes()

			str := fmt.Sprintf("%v\n", mux)
			// fmt.Println( str )

			g.Assert(strings.Contains(str, "false")).Equal(true)
			g.Assert(strings.Contains(str, "/status")).Equal(true)
			g.Assert(strings.Contains(str, "/ping")).Equal(true)
		})
	})
}
