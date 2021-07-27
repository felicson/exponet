LD_FLAGS=-X 'main.dsn=user:pass@tcp(localhost)/db'
.PHONY: build
build:
	@go build -ldflags="$(LD_FLAGS)" -o exponet cmd/main.go
clean:
	rm exponet

run:
	@go run -ldflags="$(LD_FLAGS)" cmd/main.go