NAME=echohead/static

all: build

build:
	go build --tags netgo --ldflags '-extldflags "-static"' -installsuffix cgo -o static .

release: all
	$(eval VERSION := $(shell ./static --version))
	git tag $(VERSION)
	git push origin --tags
	hub release create -a static $(VERSION)

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
