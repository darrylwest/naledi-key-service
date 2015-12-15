# Naledi Key Service

## Overview

Naledi Key Service is used to manage a set of encryption keys for a user's documents.  A "document" can be a single file or multiple files of any type combined into a single tar-ball.  The service generates and stores document keys for a user to enable encryption and decryption.  The author may share these keys with other registered users. The actual document(s) never pass throught the service.

## Document Key Sharing

A document owner may grant access to one or more registered users by requesting a shared key for that user.  The user obtains shared keys by logging in to query document access.  

When a document key is requested, it is granted only after a back-channel code has been verified.  With the key, the user is able to decrypt the owner's shared document(s).

It is important to note that the actual documents never pass through the service, only the keys.

## Communication Protocol

All communications are through a single endpoint, e.g. http://naledi.com/KeyService with routing and data parameters specified using public/private key encryption.  Keys and document specs are stored on separate private machines.

### Public/Private Key Communications

### Establishing a Session

* requestor sends public key

### Symetric Key Exchange

## License

_See the [LICENSE](LICENSE) file._

- - -
<small><em>Copyright 2015, darryl west | version 0.90.10 2015-12-14</em></small>
