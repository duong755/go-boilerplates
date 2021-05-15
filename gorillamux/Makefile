all: spec
	@go run main.go

spec:
	@cd ./go-swagger && go install ./cmd/swagger
	@~/go/bin/swagger generate spec . > swagger.json

server: spec
	@go run main.go
