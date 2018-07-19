# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=mybinary
    BINARY_UNIX=$(BINARY_NAME)_unix
	FLAGS=-extldflags "-static" -H=windowsgui
    
    all: test build
    build: 
            $(GOBUILD) -o $(BINARY_NAME) -v
    test: 
            $(GOTEST) -v ./...
    clean: 
            $(GOCLEAN)
            rm -f $(BINARY_NAME)
            rm -f $(BINARY_UNIX)
    run:
            $(GOBUILD) -o $(BINARY_NAME) -v ./...
            ./$(BINARY_NAME)
    deps:
            $(GOGET) github.com/markbates/goth
            $(GOGET) github.com/markbates/pop
    
	
    # Cross compilation
    build-linux:
        	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	build-win32:
		    GOOS=windows $(GOBUILD) -o $(BINARY_UNIX) -v -a -ldflags '$(flags)'
