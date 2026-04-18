target = main
.PHONY: run build

run:
	@go run cmd/main.go data/dummy-env.env

build: 
	@mkdir -p build/
	go build -o build/$(target) cmd/main.go

build/$(target):


clean:
	rm -rf build/