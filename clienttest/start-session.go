package main

import (
    "fmt"
    "keyservice"
    "bytes"
    "encoding/hex"
    "net/http"
    "crypto/rand"
    "io/ioutil"
    "golang.org/x/crypto/nacl/box"
    "github.com/agl/ed25519"
)

// TODO move to keyclient package
const KeySize = 32

type Config struct {
    host string
    apikey string
    hostPubKey *[KeySize]byte
    license []byte
    myPubKey *[KeySize]byte
    myPrivKey *[KeySize]byte
}

func (c *Config) CreateKeys() error {
    pub, priv, err := box.GenerateKey( rand.Reader )

    if err == nil {
        c.myPubKey = pub
        c.myPrivKey = priv
    }

    return err
}

var config *Config

func createSession() error {
    url := config.host + "/KeyService/session/create"

    msg := new(keyservice.Message)

    config.CreateKeys()

    msg.YourKey = config.hostPubKey
    msg.MyKey = config.myPubKey
    msg.Number = 1

    m, err := keyservice.EncryptBox( msg.YourKey, config.myPrivKey, config.license )
    if err != nil {
        return err
    }

    msg.EncryptedMessage = &m

    sigpub, sigpriv, err := ed25519.GenerateKey( rand.Reader )
    if err != nil {
        return err
    }

    msg.SignatureKey = sigpub
    msg.Signature = ed25519.Sign(sigpriv, *msg.EncryptedMessage)

    mbody, err := msg.EncodeToString()
    if err != nil {
        return err
    }

    request, err := http.NewRequest("POST", url, bytes.NewBufferString( mbody ))
    if err != nil {
        return err
    }

    request.Header.Set("x-api-key", config.apikey)
    request.Header.Set("Content-Type", "text/plain")
    request.Header.Set("Accept", "text/plain")

    fmt.Println( request )

    client := &http.Client{}
    resp, err := client.Do( request )
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    fmt.Println("status:", resp.Status)
    fmt.Println("headers:", resp.Header)

    body, _ := ioutil.ReadAll( resp.Body )
    fmt.Println("body:", string(body))

    // TODO : decrypt the response
    // check message number == 1

    // verify signature 

    // grab the keys

    // decrypt the response

    // pull the session number and expires data


    return nil
}

func readConfig() error {
    conf := new(Config)

    conf.host = "http://localhost"
    conf.apikey = "2a8c72cb538a47938f874ff3df2fccee" // read from license?

    key, err := hex.DecodeString("587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d")
    if err != nil {
        return err
    }
    conf.hostPubKey = new([KeySize]byte)
    copy(conf.hostPubKey[:], key)

    conf.license = []byte("license:this is my license")

    config = conf

    return nil
}

func main() {
    err := readConfig()

    if err != nil {
        panic( err )
    }

    err = createSession()
}
