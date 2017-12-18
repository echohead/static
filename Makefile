NAME=echohead/static

all: build

build:
	CGO_ENABLED=0 GOOS=linux go build -gcflags=-trimpath=$(GOPATH) -asmflags=-trimpath=$(GOPATH) -a -ldflags '-extldflags "-static"' -o static .

docker: all
	docker build -t $(NAME) .

release: docker
	docker push $(NAME)
#	$(eval VERSION := $(shell ./static --version))
#	git tag $(VERSION)
#	git push origin --tags
#	hub release create -a static $(VERSION)

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
