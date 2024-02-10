.PHONY: build version

CURRENT_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')

build:
	go build -ldflags="-X 'main.Name=myapp' -X 'main.Version=v1.0.0' -X 'main.Flag=${CURRENT_DATE}'" command.go
	./command --version
	rm -f command
