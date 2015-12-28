package models

import (
	"errors"
	"fmt"
	"time"
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
		ModelStatus.Valid:   true,
		ModelStatus.Expired: true,
	}
}

func NewUserDocument(user User, name string) UserDocument {
	doc := UserDocument{}

	doc.doi = NewDocumentIdentifier()
	doc.owner = user.doi.id
	doc.name = name

	doc.status = ModelStatus.Valid

	return doc
}

func (u UserDocument) GetDOI() DocumentIdentifier {
	return u.doi
}

func (u UserDocument) GetStatus() string {
	return u.status
}

func (u *UserDocument) SetStatus(status string) {
	u.status = status
}

func (u *UserDocument) UpdateVersion() int64 {
	u.doi.UpdateVersion()

	return u.doi.version
}

func (u UserDocument) ToJSON() ([]byte, error) {
	return MapToJSON(u.ToMap())
}

func (u *UserDocument) FromJSON(json []byte) error {
	mp, err := MapFromJSON(json)
	if err != nil {
		return err
	}

	return u.FromMap(mp)
}

func (u UserDocument) ToMap() map[string]interface{} {
	hash := u.doi.ToMap()

	hash["owner"] = u.owner
	hash["name"] = u.name
	hash["meta"] = u.meta
	hash["share"] = u.share

	if dts, err := u.expires.MarshalJSON(); err == nil {
		hash["expires"] = string(dts)
	}

	hash["status"] = u.status

	return hash
}

func (u *UserDocument) FromMap(hash map[string]interface{}) error {
	u.doi.FromMap(hash)

	u.owner = hash["owner"].(string)
	u.name = hash["name"].(string)
	u.meta = hash["meta"].(string)

	if share, ok := hash["share"].(string); ok {
		u.share = share
	}

	if expires, ok := hash["expires"].(time.Time); ok {
		u.expires = expires
	}

	u.status = hash["status"].(string)

	return nil
}

func (u UserDocument) Validate() (list []error, ok bool) {
	list = make([]error, 0, 10)

	if len(u.owner) != 32 {
		list = append(list, errors.New("document must be owned by a valid user"))
	}

	if len(u.name) < 4 {
		list = append(list, errors.New("document name must be at least 4 characters"))
	}

	if _, ok := UserDocumentStatusCodes[u.status]; ok != true {
		list = append(list, errors.New(fmt.Sprintf("status code: %s is not a recognized status for user document", u.status)))
	}

	return list, len(list) == 0
}
