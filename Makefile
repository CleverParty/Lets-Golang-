project:
	echo "Running golang files"

build:
	# go build -o bin/main keynesSimulation.go
	go get "github.com/timpalpant/go-iex"
	# go build FinData/iex_config.go
	go build FinData/SMA_go.go

run:
	# go run keynesSimulation.go
	# go run FinData/iex_config.go
	echo "Simple Moving Average calculation"
	go run FinData/SMA_go.go