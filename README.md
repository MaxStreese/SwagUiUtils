# Swagger UI Handler

Swagger UI Handler is just what the name implies: An `http.Handler` serving the 
Swagger UI.

The package is intended to be used in Go services which need to serve their own
OpenApi documentation. You can however also setup your own Swagger UI server
with it as demonstrated by the server command in this repository.

The Swagger UI resources are simply copied from 
[here](https://github.com/swagger-api/swagger-ui/tree/master/dist). The repo
further draws heavy inspiration from 
[here](https://github.com/haxii/go-swagger-ui) (so heavy in fact that you could
argue plagiarism :sweat_smile:).

I intend to implement convenience functions for using the handler with at least
the Echo framework in the future.

In case someone besides me ever uses this and wants to have the ui updated just
drop me a line and I will do the update.
