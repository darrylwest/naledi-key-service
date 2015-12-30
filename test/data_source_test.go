package keyservicetest

import (
    "fmt"
    "keyservice/dao"
    "testing"

    . "github.com/franela/goblin"
)

func TestDataSource(t *testing.T) {
    g := Goblin(t)



    g.Describe("DataSource", func() {
        fixtures := new(Fixtures)

        g.It("should create a cache only data source", func() {
            dataSource := dao.NewCachedDataSource(nil)

            g.Assert(fmt.Sprintf("%T", dataSource)).Equal("dao.DataSource")

            g.Assert(dataSource.GetCacheLen()).Equal( 0 )
            g.Assert(dataSource.GetCache() != nil).IsTrue()
        })

        g.It("should create a standard data source with primary redis client", func() {
            fixtures.CreateUserModel()
            client := dao.GetPrimaryClient()

            pong, err := client.Ping().Result()
            g.Assert(err).Equal(nil)
            g.Assert(pong).Equal("PONG")
        })

        g.It("should create a standard data source with secondary redis client", func() {
            fixtures.CreateUserModel()
            client := dao.GetSecondaryClient()

            pong, err := client.Ping().Result()
            g.Assert(err).Equal(nil)
            g.Assert(pong).Equal("PONG")
        })
    })
}
