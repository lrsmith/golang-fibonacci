TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

build: fmtcheck
	go build -o ./build/fibonacci_rest_api

# Docker 
docker-container: certs build
	$(eval IMAGE_ID=$(shell sh -c "docker build -q . | awk -F':' '{print $2}'"))

docker-run: docker-container
	docker run -p 8443:8443 -dit $(IMAGE_ID)

docker-commit: docker-run
	$(eval CONTAINER_ID=$(shell sh -c "docker ps -lq"))
	docker commit $(CONTAINER_ID) golang-fibonacci
	docker stop $(CONTAINER_ID)
	docker rm $(CONTAINER_ID)

# Test
test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

# Initialization
certs:
	if [ ! -d "config" ]; then  mkdir config; fi
	if [ ! -f "config/server.key" ]; then  \
		openssl genrsa -out config/server.key 2048; \
	fi
	if [ ! -f "config/server.crt" ]; then \
		openssl req -new -x509 -sha256 -key config/server.key -out config/server.crt -days 3650; \
	fi

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

clean:
	rm ./build/fibonacci_rest_api

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile website website-test

