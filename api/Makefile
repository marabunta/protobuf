.PHONY: all clean lint dep

all: clean
	prototool format -w && protoc api.proto -I. --go_out=plugins=grpc:.

clean:
	@rm -f *.go

lint:
	prototool lint

dep:
	go get -u github.com/golang/protobuf/protoc-gen-go
