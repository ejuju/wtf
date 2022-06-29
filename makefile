test:
	go vet
	go test ./... -race -timeout 30s -cover