package keyservicetest

import (
    "testing"
    "keyservice"
    "keyservice/dao"

    "fmt"

    "github.com/darrylwest/cassava-logger/logger"

    . "github.com/franela/goblin"
)

func TestUserDao(t *testing.T) {
    g := Goblin(t)

    // fixtures := new(Fixtures)

    g.Describe("UserDao", func() {
        ctx := keyservice.NewContextForEnvironment("test")
        ctx.ReadConfig()
        log := func() *logger.Logger {
			return ctx.CreateLogger()
		}()

        g.It("should create an instance of user dao", func() {
            log.Info("create the user dao")
            userDao := dao.CreateUserDao( ctx )

            g.Assert(userDao != nil).IsTrue()

            fmt.Println( ctx )
        })
    })
}
