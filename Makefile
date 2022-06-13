
all: build test

build:
		go build -o bin/nutanixclient -buildmode=archive -v pkg/nutanix/v3/*.go

test:
		go test -v pkg/nutanix/fc/*.go
		go test -v pkg/nutanix/v3/*.go
		go test -v pkg/nutanix/*.go  
