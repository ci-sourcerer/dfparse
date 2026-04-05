build:
	@mkdir -p dist
	go build -o dist/dfparse ./cmd/dfparse

build-image:
    docker build -t dfparse:latest .

clean:
	@rm -rf dist
