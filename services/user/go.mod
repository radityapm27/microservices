module rpm/microservices/services/user

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require github.com/patrickmn/go-cache v2.1.0+incompatible

require (
	github.com/Azure/azure-amqp-common-go v1.1.4 // indirect
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/azure-storage-blob-go v0.7.0 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.6.0 // indirect
	github.com/DataDog/zstd v1.4.1 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/franela/goblin v0.0.0-20181003173013-ead4ad1d2727 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.5.5
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/mdns v1.0.1 // indirect
	github.com/hashicorp/raft v1.1.1 // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/hashicorp/vault v1.1.4 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/labstack/gommon v0.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/nats-io/nats-streaming-server v0.15.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.2.1 // indirect
	github.com/posener/complete v1.2.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	github.com/uber/jaeger-client-go v2.15.0+incompatible // indirect
	github.com/uber/jaeger-lib v1.5.0 // indirect
	go.mongodb.org/mongo-driver v1.0.1 // indirect
	golang.org/x/mobile v0.0.0-20190806162312-597adff16ade // indirect
	google.golang.org/genproto v0.0.0-20210504143626-3b2ad6ccc450
	gopkg.in/jcmturner/gokrb5.v7 v7.3.0 // indirect
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible // indirect
	rpm/microservices/core v0.0.0
)

replace rpm/microservices/core => ../../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
