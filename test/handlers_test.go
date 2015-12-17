package keyservicetest

import (
	"testing"
	"keyservice"

	// "code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
    g := Goblin(t)

    g.Describe("Handlers", func() {
        g.It("should have a status handler that returns a json blob", func() {
            mux := http.NewServeMux()
			mux.HandleFunc("/status", keyservice.StatusHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("GET", "http://test.com/status", nil)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)
			// fmt.Println( recorder.Body.String() )
			// log.Info("status response: %s", recorder.Body.String())

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()

			wrapper := make(map[string]interface{})
			err = json.Unmarshal(recorder.Body.Bytes(), &wrapper)

			g.Assert(err == nil).IsTrue()
			status := wrapper["status"]
			g.Assert(status).Equal("ok")
        })

        g.It("should have a ping handler")
        g.It("should have a shutdown handler")
    })
}
