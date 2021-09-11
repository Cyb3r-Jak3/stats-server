.PHONY: lint test scan

full-test: lint test


lint:
	go vet 
	golint -set_exit_status 

test:
	go test -race -v -coverprofile="c.out" ./...
	go tool cover -func="c.out"

scan:
	gosec -no-fail -fmt json -out security.json ./...
	gosec -no-fail -fmt sarif -out security.sarif ./...

run:
	go build -o stats-server
	stats-server

bench:
	go test -bench=. ./...