project:
	printf "Runs checks"

build:
	# go build -o bin/main keynesSimulation.go
	printf "Build stage has begun :\n" # printf seems more functional in printing command line statements
	go get "github.com/timpalpant/go-iex"
	go get "github.com/timpalpant/go-iex/iextp/tops"
	go get "gopkg.in/jdkato/prose.v2"
	# go build FinData/iex_config.go
	go build keynesSimulation.go
	# go build poemProcess.go
	# go build FinData/pkg/SMA_go.go
	go test  FinData/pkg/SMA_go_test.go

run:
	printf "Run stage has begin :\n"
	# go run FinData/iex_config.go
	printf "^^^Build Completed^^^\n"
	# go run poemProcess.go
	# go run keynesSimulation.go
	go run FinData/pkg/SMA_go.go