package keyservicetest

import (
    "keyservice/models"
)

type Fixtures struct {

}

func (f *Fixtures) CreateDOIMap() map[string]interface{} {
    doi := models.NewDocumentIdentifier()

    return doi.ToMap()
}

func (f *Fixtures) CreateUserModel() *models.User {
    user := models.NewUser("john@ymail.com", "john@gmail.com", "7742502211@messaging.sprint.com")

    return user
}

func (f *Fixtures) CreateUserMap() map[string]interface{} {
    doi := models.NewDocumentIdentifier()
    hash := doi.ToMap()

    hash["username"] = "jane@jmail.com"
	hash["fullname"] = "jane doe"
	hash["email"] = "jane@gmail.com"
	hash["sms"] = "4156664321@messaging.sprint.com"

	hash["status"] = models.ModelStatus.Active

    return hash
}

func (f *Fixtures) CreateUserDocumentModel(user *models.User) *models.UserDocument {
    model := models.NewUserDocument(f.CreateUserModel(), "My Fixture Created User Document")

    return model
}

func (f *Fixtures) CreateUserDocumentMap() map[string]interface{} {
    user := f.CreateUserModel()
    doi := models.NewDocumentIdentifier()
    hash := doi.ToMap()

    hash["name"] = "My Fixture Map Generated Document"

    udoi := user.GetDOI()
	hash["owner"] = udoi.GetId()
	hash["meta"] = "this is my documentmeta data"
	hash["status"] = models.ModelStatus.Valid

    return hash
}
