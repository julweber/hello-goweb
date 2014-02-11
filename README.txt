== Go Web example application

This application shows the usage of:
* static files
* routes
* request arguments
	
== Run

$> go run hello-goweb.go

== Install

$> go install

== Requests

* GET -> http://localhost:9999/images/image.jpg -> static file in static/images
* GET -> http://localhost:9999/test/path?arg1=2&arg2=bla -> arguments and routes