all: test app-linux app-mac

export GO111MODULE=on

SRC := cmd/app/*.go


test: $(SRC)
	go test ./cmd/app/... -v 

app-linux: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o parameters-util-linux github.com/alext234/parameters-util/cmd/app 


app-mac: $(SRC)
	GOOS=darwin GOARCH=amd64 go build -o parameters-util-mac github.com/alext234/parameters-util/cmd/app     


