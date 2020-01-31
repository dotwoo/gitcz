
NAME=$(shell basename $(CURDIR))
VERSION=$(shell git describe --tags --long --dirty || echo "unkown version")
SHORT_VERSION=$(shell git describe --abbrev=0 --tags || echo "unknown version")
COMPILE_NUM=$(shell git describe --abbrev=1 --tags | awk -F- '{if (NF > 1) print $$2; else print 0}')
BUILDTIME=$(shell date '+%Y-%m-%d %H:%M:%S %z')
TARGET=${NAME}-${SHORT_VERSION}
UNAME_S := $(shell uname -s)


GOBUILD=CGO_ENABLED=1 go build -ldflags '-w -s'
		# -X "$(NAME)/g.Version=$(SHORT_VERSION)" \
		# -X "$(NAME)/g.BuildTime=$(BUILDTIME)" \
		# -X "$(NAME)/g.ProgName=$(NAME)" '
SOURCES=\.\/

export GO111MODULE=on
export GOPROXY=https://goproxy.io
# exported to submakes
export

all: gobuild

gobuild:
	$(GOBUILD) -o $(NAME)  $(SOURCES)

test:
	go test ./...

bench:
	go test ./... -bench=.

clean:
	git clean -xdf

lint:
	golangci-lint run

release: 
	cat goreleaser.yml.in| \
	  sed "s/startName/$(NAME)/g" | \
	  sed "s/MainDir/$(SOURCES)/g" > \
	  .goreleaser.yml
	goreleaser --snapshot --skip-publish --rm-dist


.PHONY: all clean