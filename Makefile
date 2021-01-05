.PHONY: build
build:
	@go build -ldflags="-X 'main.dsn=myprom:pass@tcp(localhost)/myprom'" -o exponet cmd/main.go
clean:
	rm exponet