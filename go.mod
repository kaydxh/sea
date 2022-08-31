module github.com/kaydxh/sea

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/kaydxh/golang v0.0.74
	github.com/prometheus/client_golang v1.12.1
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/objx v0.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9 // indirect
	google.golang.org/api v0.81.0
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/kaydxh/golang v0.0.74 => ../../../github.com/kaydxh/golang
