.PHONY: update,installapi,runapi,migrate, test,testv

update:
	go mod tidy

upgrade:
	go get -u ./...

installapi:
	cd /go/src/github.com/howood/moggiecollector/moggiecollector && export GO111MODULE=on && go install

runapi:
	export GO111MODULE=on && export GOFLAGS=-mod=mod && go run ./moggiecollector-api/moggiecollector-api.go -v

migrate:
	export GO111MODULE=on && go run ./database/migrate.go -cmd up

test:
	export GO111MODULE=on && go test ./...

testv:
	export GO111MODULE=on && go test ./... -v

testm:
	export GO111MODULE=on && go test  ./...  -run ${method} -v

testparse: install-go-test-tparse
	go test -count=1 --race ./...  -json  | tparse -all

install-go-test-tparse:
	go install github.com/mfridman/tparse@latest

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.62.2 &&  \
	./bin/golangci-lint run ./...

fmt:
	go install golang.org/x/tools/cmd/goimports@v0.28.0
	go install mvdan.cc/gofumpt@v0.7.0
	goimports -w .
	gofumpt -w .
