.PHONY: mockgen
mockgen:
	mockgen -source=./saga/step.go -destination=./mock/step_mock.go -package=mock

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: dev
dev:
	go run ./...
