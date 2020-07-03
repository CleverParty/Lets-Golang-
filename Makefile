project:
	printf "Running golang files"

build:
	# go build -o bin/main keynesSimulation.go
	printf "Build stage has begun :\n" # printf seems more functional in printing command line statements
	# go get "github.com/timpalpant/go-iex"
	go build FinData/iex_config.go
	go build keynesSimulation.go
	go build FinData/pkg/SMA_go.go

run:
	printf "Run stage has begin :\n"
	go run FinData/iex_config.go
	printf "^^^Build Completed^^^\n"
	# go run keynesSimulation.go
	# go run FinData/pkg/SMA_go.go