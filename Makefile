.PHONY: dependency unit-test cover

go-template-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

go-template-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

run-first-osx:
	make go-template-osx && ./go-template-osx

run-osx: 
	rm go-template-osx && make go-template-osx && ./go-template-osx

run-linux-first: 
	make go-template-linux && ./go-template-linux

run-linux: 
	rm go-template-linux && make go-template-linux && ./go-template-linux

dependency:
	@go get -v ./...

unit-test: dependency
	@go test -v -short ./src/usecase/... ./shared/...

cover :
	@echo "\x1b[32;1m>>> running unit test and calculate coverage \x1b[0m"
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "mode: atomic" > coverage.txt

	@go test ./app/modules/.../service ./pkg/helpers  -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./app/modules/.../service ./pkg/helpers  | grep -v mocks | tr '\n' ',')
	@go tool cover -func=coverage.txt


# Docker Build
docker: Dockerfile
	docker-compose -f "docker-compose.yml" up -d --build 