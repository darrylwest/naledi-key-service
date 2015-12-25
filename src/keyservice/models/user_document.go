package models

import (
	"errors"
    "time"
	"fmt"
)

type UserDocument struct {
	doi     DocumentIdentifier
	owner   string // User.doi.id
	name    string
	meta    string
	share   string // User.doi.id
	expires time.Time
	status  string // Valid|Expired
}

var UserDocumentStatusCodes map[string]bool

func init() {
    UserDocumentStatusCodes = map[string]bool{
        ModelStatus.Valid:true,
        ModelStatus.Expired:true,
    }
}

func NewUserDocument(user *User, name string) *UserDocument {
    doc := new(UserDocument)
    doc.doi = (*NewDocumentIdentifier())
    doc.owner = user.doi.id
    doc.name = name

    doc.status = ModelStatus.Valid

    return doc
}

func (u *UserDocument) Validate() (list []error, ok bool) {
	list = make([]error, 0, 10)

    if len(u.owner) != 32 {
        list = append(list, errors.New("document must be owned by a valid user"))
    }

    if len(u.name) < 4 {
        list = append(list, errors.New("document name must be at least 4 characters"))
    }

    if _, ok := UserDocumentStatusCodes[ u.status ]; ok != true {
        list = append(list, errors.New(fmt.Sprintf("status code: %s is not a recognized status for user document", u.status)))
    }

    return list, len(list) == 0
}
