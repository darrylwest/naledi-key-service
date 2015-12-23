package keyservice

import (
    "time"
    "encoding/hex"
    "fmt"
    "strings"
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
    SignatureKey *[KeySize]byte
    Signature *[64]byte
    Number int
    MyKey *[KeySize]byte // my public box key
    YourKey *[KeySize]byte // the peer's public box key
    Message *[]byte // message including nonce
}

func DecodeMessageFromString(encoded string) (*Message, error) {
    msg := new(Message)

    return msg, nil
}

func (m *Message) EncodeMessageToString() (string, error) {
    out := make([]string, 6)

    // TODO test all inputs to validate the message...

    out[0] = hex.EncodeToString( m.SignatureKey[:] )
    out[1] = hex.EncodeToString( m.Signature[:] )
    out[2] = fmt.Sprintf("%d", m.Number )
    out[3] = hex.EncodeToString( m.MyKey[:] )
    out[4] = hex.EncodeToString( m.YourKey[:] )
    out[5] = hex.EncodeToString( *m.Message )

    return strings.Join(out, ":"), nil
}
