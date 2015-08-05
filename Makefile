NAME=echohead/static
VERSION=latest

all: build

build:
	go build --tags netgo --ldflags '-extldflags "-static"' -installsuffix cgo -o static .

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
