# jsontosql
json to sql tool

# Build

### 1. Mac

```
go build -ldflags -w -o jsontosql-osx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags -w -o jsontosql-linux
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags -w -o jsontosql-win.exe
```

### 2. Linux

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 
```

### 3. Win

```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build 
fi  
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build 
```

Optimize packet size:

`go build -ldflags -w`

# Download

[MacOsx](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-osx)

[Linux](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-linux)

[Windows](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-win.exe)