# Service Use Cases
<small><em>Naledi Key Service</em></small>

## Register a user

All users must be registered with minimum data including back-channel access method.

## Update a user

Update a user password, email, back-channel, etc

## Request a user session

Other than registering a user all communications are conducted in a strict session

## User Login

Given a valid session, login to a registered account.

## Change a User Passord

Passwords may be changed by the original user using the existing password.  Passwords must be gauged as "strong" to qualify.

## Reset a password

Requires back-channel confirmation.

## Update user settings

User settings such as email, sms address, etc may be updated the logged in user.

## Submit a key and recieve an ID

An author may submit a key to a specific document and recieve an ID.  The ID may be passed to another user to enable sharing documents.

## Request a Shared Document Key

A document owner may request a key for a document set to be shared with a registered user.  It is assumed that this ID has be passed from another user that grants access to a specific document.

## Revoke a Shared Document Key

A document owner may revoke a shared document key at any time.

## Update a Shared Document Key

Shared keys that have an expiration date will expire to restrict access to a document set.  The expiration date may be extended by the original user.

- - -

<small><em>Copyright 2015, darryl west | version 0.90.10 2015-12-14</em></small>