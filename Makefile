
default:
	CGO_ENABLED=0 go build static.go

setup:
	# get go version of net package, to allow static linking.
	go install -tags netgo -a -v std
