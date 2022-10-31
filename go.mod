module github.com/kaydxh/sea

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/validator/v10 v10.10.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/kaydxh/golang v0.0.100
	github.com/prometheus/client_golang v1.13.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/objx v0.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	google.golang.org/api v0.100.0
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c // indirect
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/kaydxh/golang v0.0.100 => ../../../github.com/kaydxh/golang
