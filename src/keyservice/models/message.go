package models

import (
    "encoding/hex"
    "strings"
    "strconv"
    "errors"
)

const (
    KeySize = 32
    DoubleKeySize = 64
)

type Message struct {
    SignatureKey *[KeySize]byte
    Signature *[DoubleKeySize]byte
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

func parseKey64(str string) *[DoubleKeySize]byte {
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

    errs := m.Validate()
    if len(errs) > 0 {
        return "", errors.New("message structure is not valid")
    }

    out[0] = hex.EncodeToString( m.SignatureKey[:] )
    out[1] = hex.EncodeToString( m.Signature[:] )
    out[2] = strconv.Itoa( m.Number )
    out[3] = hex.EncodeToString( m.MyKey[:] )
    out[4] = hex.EncodeToString( m.YourKey[:] )
    out[5] = hex.EncodeToString( *m.EncryptedMessage )

    return strings.Join(out, ":"), nil
}

func (m *Message) Validate() []error {
    list := make([]error, 0, 6)

    if m.SignatureKey == nil {
        list = append(list, errors.New("signature key is nil"))
    }

    if m.Signature == nil {
        list = append(list, errors.New("signature is nil"))
    }

    if m.Number < 1 {
        list = append(list, errors.New("number cannot be zero"))
    }

    if m.MyKey == nil {
        list = append(list, errors.New("my box key is nil"))
    }

    if m.YourKey == nil {
        list = append(list, errors.New("you box key is nil"))
    }

    if m.EncryptedMessage == nil || len(*m.EncryptedMessage) == 0 {
        list = append(list, errors.New("encrypted message is nil or zero length"))
    }

    return list
}
