.PHONY: installapi,runapi,migrate, test,testv

installapi:
	cd /go/src/github.com/howood/moggiecollector/moggiecollector && export GO111MODULE=on && go install

runapi:
	export GO111MODULE=on && go run ./moggiecollector-api/moggiecollector-api.go -v

migrate:
	export GO111MODULE=on && go run ./database/migrate.go -cmd up

test:
	export GO111MODULE=on && go test ./...

testv:
	export GO111MODULE=on && go test ./... -v

