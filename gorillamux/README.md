# 2021/05/14

```shell
go mod init openlms

go get github.com/gorilla/mux
```

- Create `main.go`

# 2021/05/15

```shell
git submodule add https://github.com/go-swagger/go-swagger

cd ./go-swagger && go install ./cmd/swagger # this will create `~/go/bin/swagger` (`~/go/bin/swagger.exe` on Windows) executable file
```

- Create `doc.go` with `swagger:meta` annotation

```shell
swagger generate spec . > swagger.json # Generate spec file from code (with annotations, of course)

swagger serve --flavor=swagger ./swagger.json
```
