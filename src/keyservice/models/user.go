package models

import (
	"errors"
)

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

	if u.username == "" {
		list = append(list, errors.New("user name is empty"))
	} else if len(u.username) < 6 {
		list = append(list, errors.New("user name is not valid, must be at least 6 characters"))
	}

	return list, len(list) == 0
}
