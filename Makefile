export PACKAGES := $(shell env GOPATH=$(GOPATH) go list ./...)
export REVISION := $(shell git describe --exact-match --tags $(git log -n1 --pretty='%h') || git rev-parse --verify --short HEAD || echo ${REVISION})

all: clean dependencies test build

clean:
	rm -vf bin/*

build: build-linux

build-linux: build-linux-amd64 build-linux-i386 build-linux-armv7

build-linux-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.Version=${REVISION}" -o bin/microgpio-linux-amd64 .

build-linux-i386:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags "-X main.Version=${REVISION}" -o bin/microgpio-linux-i386 .

build-linux-armv7:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -ldflags "-X main.Version=${REVISION}" -o bin/microgpio-linux-armv7 .

dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

docker:
	docker build . -t microgpio

docker-test-deps:
	docker run -p3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -d mariadb
	sleep 10
	mysql -uroot -h127.0.0.1 -e 'CREATE DATABASE microgpio;'

	docker run -p5432:5432 -d postgres
	sleep 5
	psql -Upostgres -h127.0.0.1 -c 'CREATE DATABASE microgpio;'

	docker run -p6379:6379 -d redis

predependencies:
	go get -u github.com/golang/dep/cmd/dep

sign:
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-darwin-amd64.sig 				bin/microgpio-darwin-amd64
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-darwin-i386.sig 				bin/microgpio-darwin-i386
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-linux-amd64.sig 				bin/microgpio-linux-amd64
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-linux-i386.sig 					bin/microgpio-linux-i386
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-linux-armv7.sig 					bin/microgpio-linux-armv7
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-windows-amd64.exe.sig 	bin/microgpio-windows-amd64.exe
	gpg --detach-sign --digest-algo SHA512 --no-tty --batch --output bin/microgpio-windows-i386.exe.sig 		bin/microgpio-windows-i386.exe

test:
	GOCACHE=off go test -race -v ./...

benchmark:
	cd ./autocert/cache && go test -bench=. -cpu=1,2,3,4

verify:
	gpg --verify bin/microgpio-darwin-amd64.sig 				bin/microgpio-darwin-amd64
	gpg --verify bin/microgpio-darwin-i386.sig 				bin/microgpio-darwin-i386
	gpg --verify bin/microgpio-linux-amd64.sig 				bin/microgpio-linux-amd64
	gpg --verify bin/microgpio-linux-i386.sig 					bin/microgpio-linux-i386
	gpg --verify bin/microgpio-linux-armv7.sig 					bin/microgpio-linux-armv7
	gpg --verify bin/microgpio-windows-amd64.exe.sig 	bin/microgpio-windows-amd64.exe
	gpg --verify bin/microgpio-windows-i386.exe.sig 		bin/microgpio-windows-i386.exe
