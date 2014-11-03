mainly an HTTP-to-SMTP bridge, but allowing for user preferenes

> Navid (نوید), also spelled Naveed, is a Persian name meaning "bearer of good
> news" or "best wishes".

— [Wikipedia](http://en.wikipedia.org/wiki/Navid)

see [concept](https://github.com/innoq/naveed/wiki/concept) for details


Getting Started
---------------

* set up dependencies:

        $ go get github.com/gorilla/mux

  optionally, for testing:

        $ go get github.com/stretchr/testify

* ensure `NAVEED_HOST`, `NAVEED_PORT`, `NAVEED_PATH_PREFIX` (if applicable) and
  `NAVEED_ROOT_URL` environment variables are set
* launch server:

        $ make server
