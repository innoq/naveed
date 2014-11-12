mainly an HTTP-to-SMTP bridge, but allowing for user preferenes

> Navid (نوید), also spelled Naveed, is a Persian name meaning "bearer of good
> news" or "best wishes".

— [Wikipedia](http://en.wikipedia.org/wiki/Navid)

see [concept](https://github.com/innoq/naveed/wiki/concept) for details


HTTP API
--------

`/outbox` accepts form `POST`s (i.e. `application/x-www-form-urlencoded`),
expecting an `Authorization` header of the form `Bearer $token` (where `$token`
corresponds to the respective application's token as defined in `tokens.cfg`).
The request body should contain `recipient` (may occur more than once),
`subject` and `body`.

    POST /outbox HTTP/1.1
    Content-Type: application/x-www-form-urlencoded
    Authorization: Bearer abc123

    recipient=foo&recipient=bar&subject=Hello+World&body=lorem+ipsum

Note that recipients are user IDs which are automatically mapped to the
corresponding e-mail addresses.


Getting Started
---------------

* set up dependencies:

        $ go get github.com/gorilla/mux
        $ go get github.com/gorilla/handlers

  optionally, for testing:

        $ go get github.com/stretchr/testify

* ensure `NAVEED_HOST`, `NAVEED_PORT`, `NAVEED_PATH_PREFIX` (if applicable) and
  `NAVEED_ROOT_URL` environment variables are set
* ensure `tokens.cfg` is present in the application's root directory (see below)
* launch server:

        $ make server

NB:

* [`tokens.cfg`](test/fixtures/tokens.cfg) must be created and maintained
  manually - it contains an authorization string for each application
* the `REMOTE_USER` HTTP header is used to automatically determine the current
  user (i.e. the application is expected to be served via a reverse proxy)
