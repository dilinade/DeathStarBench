module github.com/harlow/go-micro-services

go 1.20

replace (
	github.com/armon/go-metrics v0.5.1 => github.com/hashicorp/go-metrics v0.5.1
	github.com/uber-go/atomic v1.11.0 => go.uber.org/atomic v1.11.0

)

require (
	github.com/bradfitz/gomemcache v0.0.0-20230611145640-acc696258285
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20171214222146-0e7658f8ee99
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/hashicorp/consul v1.0.6
	github.com/opentracing-contrib/go-stdlib v0.0.0-20180308002341-f6b9967a3c69
	github.com/opentracing/opentracing-go v1.0.2
	github.com/rs/zerolog v1.29.1
	github.com/uber/jaeger-client-go v2.11.2+incompatible
	golang.org/x/net v0.9.0
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.31.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

require (
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/armon/go-metrics v0.5.1 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/memberlist v0.5.0 // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/pascaldekloe/goe v0.1.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/uber-go/atomic v1.11.0 // indirect
	github.com/uber/jaeger-lib v1.4.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)
