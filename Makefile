build:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o cmd/Go-RedisV1.0.0-alpha.exe .