SHELL=/usr/bin/bash

all: spec
	@go run main.go

spec:
	@echo 'Installing ./cmd/swagger'
	@cd ./go-swagger && go install ./cmd/swagger
	@echo 'Generating swagger.json ...'
	@~/go/bin/swagger generate spec . > swagger.json
	@echo 'Generated swagger.json'

server: spec
	@go run main.go
