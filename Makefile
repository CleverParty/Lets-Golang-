project:
	echo "Running golang files"

build:
	go build -o bin/main keynesSimulation.go

run:
	go run keynesSimulation.go