#directorio donde se descargarán las dependencias de go
GOPATH=$(shell pwd)/vendor:$(shell pwd) 
#directorio donde se dejará el binario a compilar
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=APImapCovid19CL

build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GONAME) $(GOFILES)
