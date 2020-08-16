.PHONY: installapi,runapi,test,testv

installapi:
	cd /go/src/github.com/howood/moggiecollector/moggiecollector && export GO111MODULE=on && go install

runapi:
	export GO111MODULE=on && go run ./moggiecollector-api/moggiecollector-api.go -v

test:
	export GO111MODULE=on && go test ./...

testv:
	export GO111MODULE=on && go test ./... -v

