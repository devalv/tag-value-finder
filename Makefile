setup: ## Install all the build and lint dependencies
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.48.0
	go install golang.org/x/tools/cmd/goimports@latest

lint:
	golangci-lint run -c .golangci.yml