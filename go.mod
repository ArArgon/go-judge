module github.com/criyle/go-judge

go 1.19

require (
	github.com/creack/pty v1.1.18
	github.com/criyle/go-sandbox v0.9.13
	github.com/elastic/go-seccomp-bpf v1.3.0
	github.com/elastic/go-ucfg v0.8.6
	github.com/gin-contrib/zap v0.1.0
	github.com/gin-gonic/gin v1.8.2
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/prometheus/client_golang v1.14.0
	github.com/zsais/go-gin-prometheus v0.1.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.4.0
	golang.org/x/net v0.4.0
	golang.org/x/sync v0.1.0
	golang.org/x/sys v0.3.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.39.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/term v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
)

retract (
	// File descripter leak when multiple container fork at the same time
	[v0.9.5, v1.1.4]
	// Old version, don't use
	[v0.0.1, v0.9.4]
)
