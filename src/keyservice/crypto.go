package keyservice

import (
    "io"
    "crypto/rand"
)

const (
    KeySize = 32
    NonceSize = 24
)

func GenerateSymmetricKey() (*[KeySize]byte, error) {
    key := new([KeySize]byte)
    _, err := io.ReadFull(rand.Reader, key[:])

    return key, err
}

func GenerateNonce() (*[NonceSize]byte, error) {
    key := new([NonceSize]byte)
    _, err := io.ReadFull(rand.Reader, key[:])

    return key, err
}
