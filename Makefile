NAME=echohead/static
VERSION=latest

all: compile build

compile:
	CGO_ENABLED=0 go build static.go

build:
	docker build -t $(NAME):$(VERSION) --rm .

run:
	docker run --rm -p 8080 echohead/static /static

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
