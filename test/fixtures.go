package keyservicetest

import (
	"github.com/darrylwest/cassava-logger/logger"
	"keyservice"
	"keyservice/dao"
)

var (
	testContext *keyservice.Context
	testLogger  *logger.Logger
)

func init() {
	println("Initialize Context, Logger, DAO, Fixtures...")

	testContext = keyservice.NewContextForEnvironment("test")
	testLogger = testContext.CreateLogger()
	testContext.ReadConfig()

	dao.InitializeDao(testContext, testLogger)
}

type Fixtures struct {
}

func (f *Fixtures) CreateDOIMap() map[string]interface{} {
	doi := keyservice.NewDocumentIdentifier()

	return doi.ToMap()
}

func (f *Fixtures) CreateKnownUserModel() keyservice.User {
	hash := f.CreateUserMap()

	hash["id"] = "566292010a1c40d08ecdbda39806256d"

	obj, err := keyservice.User{}.FromMap(hash)

	if err != nil {
		panic(err)
	}

	user, _ := obj.(keyservice.User)

	return user
}

func (f *Fixtures) CreateUserModel() keyservice.User {
	user := keyservice.NewUser("john@ymail.com", "john@gmail.com", "7742502211@messaging.sprint.com")

	return user
}

func (f *Fixtures) CreateUserMap() map[string]interface{} {
	doi := keyservice.NewDocumentIdentifier()
	hash := doi.ToMap()

	hash["username"] = "jane@jmail.com"
	hash["fullname"] = "jane doe"
	hash["email"] = "jane@gmail.com"
	hash["sms"] = "4156664321@messaging.sprint.com"

	hash["status"] = keyservice.ModelStatus.Active

	return hash
}

func (f *Fixtures) CreateUserDocumentModel(user keyservice.User) keyservice.UserDocument {
	model := keyservice.NewUserDocument(f.CreateUserModel(), "My Fixture Created User Document")

	return model
}

func (f *Fixtures) CreateUserDocumentMap() map[string]interface{} {
	user := f.CreateUserModel()
	doi := keyservice.NewDocumentIdentifier()
	hash := doi.ToMap()

	hash["name"] = "My Fixture Map Generated Document"

	udoi := user.GetDOI()
	hash["owner"] = udoi.GetId()
	hash["meta"] = "this is my documentmeta data"
	hash["status"] = keyservice.ModelStatus.Valid

	return hash
}

func (f *Fixtures) CreateAccessKeyMap() map[string]interface{} {
	mp := map[string]interface{}{
		"id":  "my-access-key-id",
		"key": "my-access-key-value",
	}

	return mp
}
