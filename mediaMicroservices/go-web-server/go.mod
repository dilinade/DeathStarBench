module go-web-server

replace github.com/apache/thrift/tutorial/go/gen-go/media => ../gen-go/media

go 1.21

require (
	github.com/apache/thrift v0.19.0
	github.com/apache/thrift/tutorial/go/gen-go/media v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
)
