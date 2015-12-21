package keyservice

import (
    "time"
)

type DocumentIdentifier struct {
    id string
    dateCreated time.Time
    lastUpdated time.Time
    version int64
}

type User struct {
    doi  DocumentIdentifier
    username string
    fullname string
    email string
    sms string
    status string // Active, Inactive, Banned
}

type UserDocument struct {
    doi DocumentIdentifier
    owner string // User.doi.id
    name string
    meta string
    share string // User.doi.id
    expires time.Time
    status string // Valid|Expired
}

type ChallengeCode struct {
    doi DocumentIdentifier
    challengeType string // Document, Access
    sendTo string
    sendDate time.Time
    expires time.Time
    status string // Active, Canceled, Expired
}

type AccessKey struct {
    id string
    key []byte
}

type Message struct {
    peerPubKey *[KeySize]byte
    nonce *[NonceSize]byte
    sigPubKey *[KeySize]byte
    sig *[64]byte
    number uint32
    message *[]byte
}
