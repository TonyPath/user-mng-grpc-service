SHELL := /bin/bash

# ==============================================================================
# Docker support

serve:
	docker-compose -f infra/docker-compose.yaml up -d users_mng_svc

rebuild: stop
	docker-compose -f infra/docker-compose.yaml up -d --build --force-recreate users_mng_svc

stop:
	docker-compose -f infra/docker-compose.yaml down

logs:
	docker logs users_mng_svc

# ==============================================================================
# Proto support

gen-proto:
	rm -Rf proto
	mkdir -p proto
	protoc --go_out=proto --go-grpc_out=proto ./proto-schemas/services/user/*
	protoc --go_out=proto --go-grpc_out=proto ./proto-schemas/events/*

# ==============================================================================

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.48.0 golangci-lint run -v

# =============================================================================

test:
	go test ./... -coverprofile=coverage.out -count=1

coverage: test
	go tool cover -func coverage.out

report: coverage
	go tool cover -html=coverage.out -o cover.html

update-deps:
	go get -u -v ./...
	go mod tidy
