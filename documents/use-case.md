# Naledi Key Service API Use Cases

## Request a user session

All communications are conducted in strict, short-lived sessions with sequencially numbered messages sent with asymmetric keys.  To begin a session, the client application requests a session token by accessing the session/create route and supplying it's public-key.  The server responds by creating a session and returns the session token encrypted with the user's public-key (peer key).  

Subsequent requests from the client/user must add the session token to the header and encrypt all request data using the server's public key.  Responses from the server are encrypted with the client's public key.  If any errors are detected, the session is terminated.

## Register a user

All users must be registered with minimum data including email and back-channel access method (SMS, socket, etc).

## User Login

Given a valid session, the user must login to a registered account to obtain document information.  A successful login will return the user's settings.

## Change a User Password

Passwords may be changed by the original user using the existing password.  Passwords must be gauged as "strong" to qualify.

## Reset a password

Requires back-channel confirmation that supplies a one time use code (8 to 12 digits).

## Update user settings

User settings such as email, sms address, etc may be updated the logged in user.

## Request a list of documents

The user may request a list of all documents that the user has access to.  The list is fetched from the document database, encrypted using the peer key, and returned.

## Request an encryption key

To encrypt a document an author requests a symmetry key from the service by supplying document meta data.  The server responds with an ID used to identify the document and meta data (title, author, create date, update date, abstract, status) plus the newly generated key.  The key is then encrypted using the servers private symmetric key and stored in a remote database (id:key).

## Request Creation of a Shared Key

A document owner may request the creation of a shared key for a document with a registered user as designated by the author/owner.  The key is generated and stored for the registered user and the specified document.  Both the document owner and the registered user now share and have access this key.  The registered owner has the ability to update or revoke access to the shared key.

## Revoke Access from a Shared Document Key

A document owner may revoke a shared document key at any time.  The key and it's relations to author and registered user are deleted.

## Update Access to a Shared Document Key

Shared keys that have an expiration date will expire to restrict access to a document set.  The expiration date may be extended by the original user.

## Request Access to a Shared Key

A registered user that has been granted access to a document may request the key to enable decryption.  If the key is valid, the service responds by creating a single-use challenge code that is sent to the user via back-channel (SMS, email, etc).   The registered user then sumbits the challenge code and if verified, the service returns the shared key or an error message if the code or shared key are invalid.

- - -

<small><em>Copyright 2015, darryl west | version 0.90.10 2015-12-19</em></small>