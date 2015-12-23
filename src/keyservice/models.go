package keyservice

import (
    "time"
    "encoding/hex"
    "strconv"
    "strings"
    "errors"
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
    EncryptedMessage *[]byte // encrypted message including nonce
}

func parseKey(str string) *[KeySize]byte {
    if len(str) != 64 {
        return nil
    }

    v, err := hex.DecodeString( str )
    if err != nil {
        return nil
    }

    var buf [KeySize]byte
    copy(buf[:], v)

    return &buf
}

func parseKey64(str string) *[64]byte {
    if len(str) != 128 {
        return nil
    }

    v, err := hex.DecodeString( str )
    if err != nil {
        return nil
    }

    var buf [64]byte
    copy(buf[:], v)

    return &buf
}

func DecodeMessageFromString(encoded string) (*Message, error) {
    msg := new(Message)

    parts := strings.Split(encoded, ":")

    msg.SignatureKey = parseKey(parts[0])
    msg.Signature = parseKey64(parts[1])
    n, err := strconv.Atoi( parts[2] )
    if err != nil {
        return nil, err
    }

    msg.Number = n

    msg.MyKey = parseKey( parts[3] )
    msg.YourKey = parseKey( parts[4] )

    m, err := hex.DecodeString( parts[5] )
    if err != nil {
        return nil, err
    }

    msg.EncryptedMessage = &m

    if msg.SignatureKey == nil || msg.Signature == nil || msg.MyKey == nil || msg.YourKey == nil {
        return nil, errors.New("could not decode Message from string...")
    }

    return msg, nil
}

func (m *Message) EncodeToString() (string, error) {
    out := make([]string, 6)

    // TODO test all inputs to validate the message...

    out[0] = hex.EncodeToString( m.SignatureKey[:] )
    out[1] = hex.EncodeToString( m.Signature[:] )
    out[2] = strconv.Itoa( m.Number )
    out[3] = hex.EncodeToString( m.MyKey[:] )
    out[4] = hex.EncodeToString( m.YourKey[:] )
    out[5] = hex.EncodeToString( *m.EncryptedMessage )

    return strings.Join(out, ":"), nil
}
