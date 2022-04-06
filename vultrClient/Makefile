REPO=$(PWD)
EXECUTABLE=$(REPO)/build/vultrClient.bin

TOKEN=$(cat $(REPO)/secrets/vultrKeyAPI.secrets)

all: build

build:
	@# Create directory if it does not exist
	mkdir -p $(REPO)/build
	go build -o $(EXECUTABLE) $(REPO)/cmd/api/*.go

run: build
	$(EXECUTABLE) -tokenAPI ${API_TOKEN}

clean:
	rm $(REPO)/build/*.bin

.PHONY : build run clean

