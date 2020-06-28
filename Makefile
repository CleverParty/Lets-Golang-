project:
	echo "Running golang files"

build:
	# go build -o bin/main keynesSimulation.go
	go get "github.com/timpalpant/go-iex"
	go build FinData/iex_config.go

run:
	# go run keynesSimulation.go
	go run FinData/iex_config.go