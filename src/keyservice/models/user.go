package models

import (
	"errors"
	"net/mail"
	"fmt"
)

type User struct {
	doi      DocumentIdentifier
	username string
	fullname string
	email    string
	sms      string
	status   string // Active, Inactive, Banned
}

var UserStatusCodes map[string]bool

func init() {
	UserStatusCodes = map[string]bool{
		ModelStatus.Active:true,
		ModelStatus.Inactive:true,
		ModelStatus.Banned:true,
		ModelStatus.Deleted:true,
	}
}

func NewUser(username, email, sms string) *User {
	user := new(User)
	user.doi = (*NewDocumentIdentifier())

	user.username = username
	user.email = email
	user.sms = sms

	user.status = ModelStatus.Active

	return user
}

func (u *User) GetDOI() DocumentIdentifier {
	return u.doi
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetFullname() string {
	return u.fullname
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetSMS() string {
	return u.sms
}

func (u *User) GetStatus() string {
	return u.status
}

func (u *User) SetStatus(status string) {
	u.status = status
}

func (u *User) ToMap() map[string]interface{} {
	hash := u.doi.ToMap()

	hash["username"] = u.username
	hash["fullname"] = u.fullname
	hash["email"] = u.email
	hash["sms"] = u.sms

	hash["status"] = u.status

	return hash
}

func (u *User) FromMap(hash map[string]interface{}) error {
	u.doi.FromMap( hash )

	u.username = hash["username"].(string)
	u.fullname = hash["fullname"].(string)

	u.email = hash["email"].(string)
	u.sms = hash["sms"].(string)
	u.status = hash["status"].(string)

	return nil
}

func (u *User) Validate() (list []error, ok bool) {
	list = make([]error, 0, 10)

	if !validateEmail(u.username) {
		list = append(list, errors.New("user name is not valid, it should look like an email"))
	}

	if !validateEmail(u.email) {
		list = append(list, errors.New("user must have a vaild email address"))
	}

	if _, ok := UserStatusCodes[ u.status ]; ok != true {
		list = append(list, errors.New(fmt.Sprintf("status code: %s is not a recognized status", u.status)))
	}

	return list, len(list) == 0
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress( email )

	return err == nil
}
