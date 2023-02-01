# 2023-02-01

- Use go 1.19
- Use https
- Write README.md
- Rewrite Makefile's spec target

# 2021-05-20

- Downloads and serve static files, use these instead of CDN:
    - `swagger-ui.css`
    - `swgger-ui-bundle.js`
    - `swagger-ui-standalone-preset.js`
    - `favicon-16x16.png`
    - `favicon-32x32.png`

# 2021-05-18

- Print message before serving http

# 2021/05/15

```shell
git submodule add https://github.com/go-swagger/go-swagger

cd ./go-swagger && go install ./cmd/swagger # this will create `~/go/bin/swagger` (`~/go/bin/swagger.exe` on Windows) executable file
```

- Create `doc.go` with `swagger:meta` annotation

```shell
swagger generate spec . > swagger.json # Generate spec file from code (with annotations, of course)

swagger serve -F swagger ./swagger.json
```

- Serve `swagger.json` file and Swagger UI
- Edit `swagger:response` by adding `in: body`



# 2021/05/14

```shell
go mod init github.com/duong755/go-boilerplates

go get github.com/gorilla/mux
```

- Create `main.go`
