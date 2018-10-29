.PHONY: all clean lint

all: clean
	prototool format -w && protoc api.proto -I. --go_out=plugins=grpc:.

clean:
	@rm -f *.go

lint:
	prototool lint
