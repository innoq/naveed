HTTP-to-SMTP bridge for arbitrary providers to send e-mail notifications based
on user preferences

> Navid (نوید), also spelled Naveed, is a Persian name meaning "bearer of good
> news" or "best wishes".

— [Wikipedia](http://en.wikipedia.org/wiki/Navid)


HTTP API
--------

`/outbox` accepts form `POST`s (i.e. `application/x-www-form-urlencoded`),
expecting an `Authorization` header of the form `Bearer $token` (where `$token`
corresponds to the respective application's token as defined in `tokens.cfg`).
The request body should contain `recipient` (may occur more than once),
`subject` and `body`. It may optionally contain a custom `sender`.

    POST /outbox HTTP/1.1
    Content-Type: application/x-www-form-urlencoded
    Authorization: Bearer abc123

    sender=foo&recipient=bar&recipient=baz&subject=Hello+World&body=lorem+ipsum

Note that sender and recipients are user IDs which are automatically mapped to
the corresponding e-mail addresses.


Getting Started
---------------

* set up dependencies:

        $ go get github.com/gorilla/mux
        $ go get github.com/gorilla/handlers
        $ go get github.com/BurntSushi/toml

  optionally, for testing:

        $ go get github.com/stretchr/testify

* configure your application by customizing `naveed.ini`
* ensure the following environment variables are set:
    * `NAVEED_USERS_URL`, `NAVEED_USERS_USERNAME` and `NAVEED_USERS_PASSWORD`
      for synchronizing the user index - this can be avoided by manually placing
      `users.json` in the application's root directory
* ensure `tokens.cfg` is present in the application's root directory (see below)
* launch server:

        $ make server

NB:

* [`tokens.cfg`](test/fixtures/tokens.cfg) must be created and maintained
  manually - it contains an authorization string for each application
* the `REMOTE_USER` HTTP header is used to automatically determine the current
  user (i.e. the application is expected to be served via a reverse proxy)


To use with docker

* build with `docker build --rm -t naveed .`
* run with   `docker run --name runningnaveed -p80:8465 -it -e NAVEED_USERS_URL=XXX -e NAVEED_USERS_USERNAME=YYY -e NAVEED_USERS_PASSWORD=ZZZ --rm naveed`

Architectural Overview
----------------------

Naveed accepts messages via HTTP (cf. [HTTP API](#http-api)), turns the
respective user IDs of sender and recipients into e-mail addresses via the user
index, and sends those messages via e-mail while remaining oblivious to the
contents.

Incoming messages are required to provide an authorization token, which is
verified against a list of known tokens - ideally each "message provider"
application is assigned a separate token.

Users can mute individual message providers. Those preferences are stored as
plain text files.

Message delivery is delegated to Unix `sendmail` to avoid the need for SMTP
authentication - i.e. the host system is expected to be configured accordingly.

The user index is a simple JSON file which maps user IDs to full names and
e-mail addresses. If the corresponding settings are present (cf.
[Getting Started](#getting-started)), that index will periodically be
synchronized.
