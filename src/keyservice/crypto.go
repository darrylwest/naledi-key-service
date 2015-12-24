package keyservice

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"keyservice/models"
)

const (
	KeySize   = models.KeySize
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

// encrypt the message and prepend the output with nonce
func EncryptSymmetric(key *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		log.Error(encryptFailedMessage, err)
		return nil, errors.New(fmt.Sprintf(encryptFailedMessage, err))
	}

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)

	return out, nil
}

// strip the leading nonce and decrypt the message
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

// encrypt with peer pub/priv keys ; prepend the output with nonce
func EncryptBox(peer, priv *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		log.Error(encryptFailedMessage, err)
		return nil, errors.New(fmt.Sprintf(encryptFailedMessage, err))
	}

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = box.Seal(out, message, nonce, peer, priv)

	return out, nil
}

// decrypt with peer pub/priv keys by first stripping the prepended nonce
func DecryptBox(peer, priv *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (box.Overhead + NonceSize) {
		log.Error(decryptFailedMessage, "message too short")
		return nil, errors.New(fmt.Sprintf(decryptFailedMessage, "message too short"))
	}

	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])

	out, ok := box.Open(nil, message[NonceSize:], &nonce, peer, priv)
	if !ok {
		log.Error(decryptFailedMessage, "message could not be decrypted")
		return nil, errors.New(fmt.Sprintf(decryptFailedMessage, "message could not be decrypted"))
	}

	return out, nil
}

func DecodeKeyPair(spub, spriv string) (*[32]byte, *[32]byte, error) {
	pub, err := hex.DecodeString(spub)
	if err != nil {
		return nil, nil, err
	}

	priv, err := hex.DecodeString(spriv)
	if err != nil {
		return nil, nil, err
	}

	var pk *[KeySize]byte = new([KeySize]byte)
	var vk *[KeySize]byte = new([KeySize]byte)

	copy(pk[:], pub)
	copy(vk[:], priv)

	return pk, vk, nil
}

// clear the buffer bytes to zero; should be used to clear encryption keys
func ClearBytes(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}
}
