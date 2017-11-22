include env/docker.mk
include env/docker-compose.mk

.PHONY: deps
deps:
	dep ensure -v

.PHONY: run
run: BIND = 127.0.0.1
run: PORT = 8080
run:
	( \
	  export BIND=$(BIND) PORT=$(PORT); \
	  go run -ldflags '-s -w -X main.version=dev -X main.commit=unknown -X main.date=unknown' main.go build.go $(COMMAND); \
	)

.PHONY: help
help: COMMAND = help
help: run

.PHONY: migrate
migrate: COMMAND = migrate
migrate: run

.PHONY: server
server: COMMAND = run --with-profiler
server: run

.PHONY: version
version: COMMAND = version
version: run
