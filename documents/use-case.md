# Naledi Key Service API Use Cases

## Request a user session

All communications are conducted in strict, short-lived sessions.  To begin a session, the client application requests a session token by supplying it's public-key and session/request route.  The server responds by creating a session and returns the session token and message encryption symmetric key encrypted with the user's public-key.  Subsequent requests from the user must add the session token to the header and encrypt all request data using the message key.  Responses from the server are encrypted with the message key.

## Register a user

All users must be registered with minimum data including email and back-channel access method (SMS, socket, etc).

## User Login

Given a valid session, the user must login to a registered account to obtain document information.  A successful login will return the user's settings and a document list with meta data.

## Change a User Passord

Passwords may be changed by the original user using the existing password.  Passwords must be gauged as "strong" to qualify.

## Reset a password

Requires back-channel confirmation that supplies a one time use code (8 to 12 digits).

## Update user settings

User settings such as email, sms address, etc may be updated the logged in user.

## Request an encryption key

To encrypt a document an author requests a symmetry key from the service by supplying document meta data.  The server responds with an ID used to identify the document and meta data (title, author, create date, update date, abstract, status).

## Request a Shared Document Key

A document owner may request a key for a document to be shared with a registered user as designated by the author/owner.  The key is generated and stored for the registered user.  Both the document owner and the registered user now share this key.

## Revoke Access from a Shared Document Key

A document owner may revoke a shared document key at any time.  The key and it's relations to author and registered user are deleted.

## Update Access to a Shared Document Key

Shared keys that have an expiration date will expire to restrict access to a document set.  The expiration date may be extended by the original user.

- - -

<small><em>Copyright 2015, darryl west | version 0.90.10 2015-12-14</em></small>