package models

import (
	"time"
)

var ModelStatus *ModelStatusType

type ModelStatusType struct {
	Active string
	Inactive string
	Deleted string
	Banned string
	Valid string
	Expired string
	Canceled string
}

func init() {
	ModelStatus = new(ModelStatusType)

	ModelStatus.Active = "active"
	ModelStatus.Inactive = "inactive"
	ModelStatus.Deleted = "deleted"
	ModelStatus.Banned = "baned"
	ModelStatus.Valid = "valid"
	ModelStatus.Expired = "expired"
	ModelStatus.Canceled = "canceled"
}

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
