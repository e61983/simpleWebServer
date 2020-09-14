TARGET = simpleWebServer
SRC = main.go

GO ?= go
GOFMT ?= gofmt "-s"

all: $(TARGET)

$(TARGET): $(SRC)
	$(GO) build -v -tags '$(TAGS)' -o release/$@

.PHONY: clean
clean:
	@$(RM) -r release

.PHONY: fmt
fmt:
	@echo Fotmatting... $(SRC)
	@$(GOFMT) -w $(SRC)
