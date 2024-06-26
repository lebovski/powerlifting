SHELL := /bin/bash
SCHEMA_LOCAL_PATH=${SCHEMA_PATH}
ifeq (${SCHEMA_PATH},)
	SCHEMA_LOCAL_PATH=${GOPATH}/src/gitlab.stageoffice.ru/UCS-COMMON/schema
endif

all: clean fix lint test build calculate

fix:
	gofmt -w .
	go mod tidy

clean:
	rm calculator || true
	rm train.pdf || true

build:
	go build -mod=readonly -o calculator main.go

power:
	./calculator -c ./internal/config/powerlifting_an.json
	./calculator -c ./internal/config/powerlifting_dol.json

power_light:
	./calculator -c ./internal/config/powerlifting_light.json

body:
	./calculator -c ./internal/config/bodybuilding.json

street:
	./calculator -c ./internal/config/streetlifting_an.json
	./calculator -c ./internal/config/streetlifting_ad.json

calculate: power

test:
	go test -mod=readonly -coverprofile overalls.coverprofile ./...
	go tool cover -func=./overalls.coverprofile

lint:
	golangci-lint run ./...