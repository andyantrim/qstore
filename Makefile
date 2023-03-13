.PHONY: default install build 

default: build

install: install-docker install-buildx install-yq

build:
	docker build \
	  --build-arg BUILD_DATE=$(shell date --iso-8601=minutes) \
	  --build-arg BUILD_VCS_REF=$(shell git rev-parse --short HEAD) \
	  -t andyantrim/qstore \
	  .
