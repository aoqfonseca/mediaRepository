SRC = $(GOPATH)/src
GITHUBCOM = $(GOPATH)/src/github.com

setup:
	@go build .

run: build
	@shell server --config config.yml 

test:
	@go test -i ./...
	@go test ./...

check:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo " details, please check INSTALL.md file and/or"
	@echo " http://golang.org/cmd/go/#GOPATH_environment_variable"
	@exit 1
endif

setup: check
	@mkdir -p $(GITHUBCOM)
	@ln -f -s `pwd`/  $(GITHUBCOM)/mediaRepository
	@go get github.com/tools/godep
	@$(GOPATH)/bin/godep restore

build:
	@go build webserver/server.go
