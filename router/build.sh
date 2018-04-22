(cd common && go build) || exit 10
(cd standard && go build) || exit 10
(cd appengine && go build) || exit 10
(cd fasthttp && go build) || exit 10
