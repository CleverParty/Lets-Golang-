LINTER=golangci-lint

all: test build

test:
    go test keynesSimulation.go

build:
    go build keynesSimulation.go

lint:
	$(LINTER) run