package keyservice

import (
    "io"
    "crypto/rand"
    "errors"
    "fmt"
    "golang.org/x/crypto/nacl/secretbox"
)

const (
    KeySize = 32
    NonceSize = 24
)

var (
    encryptFailedMessage = "encryption failed: %s"
    decryptFailedMessage = "decryption failed: %s"
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

func EncryptSymmetric(key *[KeySize]byte, message []byte) ([]byte, error) {
    nonce, err := GenerateNonce()
    if err != nil {
        log.Error(encryptFailedMessage, err)
        return nil, errors.New(fmt.Sprintf( encryptFailedMessage, err ))
    }

    out := make([]byte, len(nonce))
    copy(out, nonce[:])
    out = secretbox.Seal(out, message, nonce, key)

    return out, nil
}

func DecryptSymmetric(key *[KeySize]byte, message []byte) ([]byte, error) {
    if len(message) < (NonceSize + secretbox.Overhead) {
        log.Error(decryptFailedMessage, "message too short")
        return nil, errors.New(fmt.Sprintf(decryptFailedMessage, "message too short"))
    }

    var nonce [NonceSize]byte
    copy(nonce[:], message[:NonceSize])
    out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
    if !ok {
        log.Error(decryptFailedMessage, "unknown reason")
        return nil, errors.New(fmt.Sprintf(decryptFailedMessage, "unknown reason"))
    }

    return out, nil
}