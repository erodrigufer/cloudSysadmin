REPO=$(PWD)
EXECUTABLE=$(REPO)/build/vultrClient.bin

all: build

build:
	@# Create directory if it does not exist
	mkdir -p $(REPO)/build
	go build -o $(EXECUTABLE) $(REPO)/cmd/api/*.go

run: build
	@# API_TOKEN is then a parameter for the Makefile as in: 
	@# 'make build API_TOKEN=<...>'
	$(EXECUTABLE) -tokenAPI ${API_TOKEN}

clean:
	rm $(REPO)/build/*.bin

.PHONY : build run clean

