# Service Use Cases
<small><em>Naledi Key Service</em></small>

## Register a user

All users must be registered with minimum data including email and back-channel access method (SMS, socket, etc).

## Request a user session

With the exception of registering a new user, all communications are conducted in strict, short-lived sessions.  To begin a session, the user requests a session token by supplying a public-key and session/request route.  The server responds by creating a session and returns the session token and message encryption symmetric key encrypted with the user's public-key.  Subsequent requests from the user must add the session token to the header and encrypt all request data using the message key.  Responses from the server will also encrypt responses with the message key as long as the session is active.

## User Login

Given a valid session, the user must login to a registered account to obtain document information.  A successful login will return the user's settings and a document list with meta data.

## Change a User Passord

Passwords may be changed by the original user using the existing password.  Passwords must be gauged as "strong" to qualify.

## Reset a password

Requires back-channel confirmation.

## Update user settings

User settings such as email, sms address, etc may be updated the logged in user.

## Request an encryption key

To encrypt a document an author requests a symmetry key from the service by supplying document meta data.  The server responds is an ID used to identify the document and meta data (title, author, create date, update date, abstract).

## Request a Shared Document Key

A document owner may request a key for a document set to be shared with a registered user.  It is assumed that this ID has be passed from another user that grants access to a specific document.

## Revoke a Shared Document Key

A document owner may revoke a shared document key at any time.

## Update a Shared Document Key

Shared keys that have an expiration date will expire to restrict access to a document set.  The expiration date may be extended by the original user.

- - -

<small><em>Copyright 2015, darryl west | version 0.90.10 2015-12-14</em></small>