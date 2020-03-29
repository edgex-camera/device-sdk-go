module github.com/edgex-camera/device-sdk-go

go 1.13

require (
	github.com/edgexfoundry/go-mod-core-contracts v0.1.36
	github.com/edgexfoundry/go-mod-registry v0.1.9
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/google/uuid v1.1.0
	github.com/gorilla/mux v1.7.4
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go v1.1.4
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/edgexfoundry/go-mod-core-contracts => github.com/edgexfoundry/go-mod-core-contracts v0.1.14

replace github.com/edgexfoundry/go-mod-registry => github.com/edgexfoundry/go-mod-registry v0.1.9

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
