NAME=echohead/static
VERSION=latest

all: build image

build:
	CGO_ENABLED=0 go build static.go

image:
	docker build -t $(NAME):$(VERSION) --rm .

run:
	docker run --rm -p 3000:8080 echohead/static --bind :8080 --root /

push:
	docker push $(NAME):$(VERSION)

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
