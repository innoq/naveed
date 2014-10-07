mainly an HTTP-to-SMTP bridge, but allowing for user preferenes

> Navid (نوید), also spelled Naveed, is a Persian name meaning "bearer of good
> news" or "best wishes".

— [Wikipedia](http://en.wikipedia.org/wiki/Navid)


Getting Started
---------------

* set up dependencies:

        $ go get github.com/gorilla/mux

* launch server:

        $ make server


Concept
-------

    GET /admin

    GET /apps  -->  [{ ID, URI, name }]
    POST /apps { ID, name }  -->  token # newly generated
    GET /apps/<ID>  -->  { ID, name, token }
    DELETE /apps/<ID>

    GET /preferences  -->  200 # checkbox for each app
    POST /preferences { ID=<0|1> }

    POST /outbox { token, recipients, subject, body }  -->  202

        +-----+                  +--------+               +------+
        | app | ---------------> | Naveed | - - SMTP - -> | user |
        +-----+   POST /outbox   +--------+               +------+
                                     ^
                                     | POST /preferences
                                     |
                                  +------+
                                  | user |
                                  +------+

note that the premise here is that it's each individual application's
responsibility to decide when to send a notification, since that's not
something we can define generically

however, it's up to Naveed whether a notification is forwarded to the
respective users (usually based on their preferences) - that is, applications
do not need to differentiate and should always send notifications for relevant
events
