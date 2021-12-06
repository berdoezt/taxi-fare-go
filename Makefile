APP_NAME = TAXI-FARE-GO

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo 
	@echo 'Usage:'
	@echo '    make build                              Compile the project.'
	@echo '    make run                                Run binary.'
	@echo '    make test                               Run tests on a compiled project.'
	@echo '    make generate                           Find all go:generate command command(s) and execute it.'
	@echo '    make clean                              Clean.'
	@echo

.PHONY: build
build:
	@echo 'Building ${APP_NAME}'
	go build -o ${APP_NAME} main.go

.PHONY: run
run:
	@echo 'Running ${APP_NAME}'
	./${APP_NAME}

.PHONY: test
test:
	@echo 'Running unit test'
	go test ./... -race -cover

.PHONY: generate
generate:
	@echo 'Running generate'
	go generate ./...

.PHONY: clean
clean:
	@echo 'Cleaning binary'
	@test ! -e ${APP_NAME} || rm ${APP_NAME}