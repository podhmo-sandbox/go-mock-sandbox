test:
	go test -v ./...

gen:
	go generate ./...

setup:
	go get -v github.com/travisjeffery/mocker/cmd/mocker
	go get -v github.com/vektra/mockery/.../
