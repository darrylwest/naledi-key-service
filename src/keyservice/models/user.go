package models

type User struct {
	doi      DocumentIdentifier
	username string
	fullname string
	email    string
	sms      string
	status   string // Active, Inactive, Banned
}

func (u *User) Validate() (list []error, ok bool) {
	list = make([]error, 0, 10)

	return list, len(list) == 0
}
