# Swag UI Utils

At the core of Swag UI Utils is the *swaguihandler* package which implements a
`http.Handler` that serves the 
[Swagger UI](https://swagger.io/tools/swagger-ui/).

The *swaguihandler* package in turn relies on the *swaguidist* package which
contains the static files from the 
[Swagger UI dist](https://github.com/swagger-api/swagger-ui/tree/master/dist).
The index.html file has been put into a Go constant which is used to create a
html template. All other files have been transformed to byte slices using the
*xxd* utility with the *-i* flag.

The idea for the *swaguidist* and *swaguihandler* packages is taken from the
[go-swagger-ui](https://github.com/haxii/go-swagger-ui) repository so all credit
regarding the byte slice idea goes to that repository.

I created this repository because I wanted to have more flexibility in serving
the UI than what the above repository offers. With this package you are able to
let your Go service serve its own OpenApi documentation via the handler or use 
the *swaguiserver* to have a standalone Swagger UI server or build your own.
Because the `SwagUiHandler` type inside *swaguihandler* implements 
`http.Handler` using the handler with frameworks like 
[Echo](https://github.com/labstack/echo) is straight forward.
