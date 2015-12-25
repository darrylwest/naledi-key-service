package models

import (
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

type ChallengeCode struct {
	doi           DocumentIdentifier
	challengeType string // Document, Access
	sendTo        string
	sendDate      time.Time
	expires       time.Time
	status        string // Active, Canceled, Expired
}

type AccessKey struct {
	id  string
	key []byte
}
