SHELL := /bin/bash

build: 
	@[ -d bin ] || mkdir bin
	( . ./.setpath ; go build -o bin/keyservice src/main.go )

install-deps:
	go get github.com/codegangsta/negroni
	go get code.google.com/p/go-uuid/uuid
	go get github.com/phyber/negroni-gzip/gzip
	go get gopkg.in/tylerb/graceful.v1
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/franela/goblin
	go get gopkg.in/redis.v3
	go get golang.org/x/crypto/nacl/secretbox
	go get golang.org/x/crypto/nacl/box
	go get golang.org/x/crypto/scrypt
	go get github.com/agl/ed25519

format:
	( gofmt -s -w src/*.go src/*/*.go src/*/*/*.go test/*.go )

qtest:
	@( . ./.setpath ; cd test ; go test )

test:
	@( make qtest ) | tee /tmp/keyservice.test
	@( . ./.setpath ; go vet src/*.go ; go vet src/*/*.go ; go vet src/keyservice/models/*.go ) | tee /tmp/keyservice.vet

watch:
	./watcher.js

client:
	( . ./.setpath ; cd clienttest ; go run start-session.go )

run:
	( . ./.setpath ; go run src/main.go --env=staging --configFile=test/test-config.json )

start:
	nohup /usr/local/bin/keyservice --baseport=9001 --serverCount=2 --shutdownPort=9009 --logname=keyservice --env=production & 

status:
	curl http://localhost:9001/status
	curl http://localhost:9002/status

ping:
	curl http://localhost:9001/ping
	curl http://localhost:9002/ping

stop:
	curl -X POST http://localhost:9009/shutdown

install-webserver:
	@make build
	@( sudo cp ./bin/keyservice /usr/local/bin/ )
	@webserver --version

install:
	@make install-webserver


.PHONY: format
.PHONY: test
.PHONY: qtest
.PHONY: watch
.PHONY: run
