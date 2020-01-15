ifeq (, $(shell which go))
$(info "No 'golang' in PATH, please install golang to build")
$(info "  linux      : sudo apt-get install golang")
$(info "  macosx     : brew install golang")
$(info "  download   : https://golang.org/download/")
$(error "unable to complete build")
endif


.DEFAULT_GOAL := standard

BINARY_NAME ?= response-caching-demo

standard:
	$(info Building binary with name: $(BINARY_NAME))
	@go build -o $(BINARY_NAME) .

clean:
	@rm -f $(BINARY_NAME)
