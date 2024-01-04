APP_NAME=debt-snowball
APP_VERSION=v2.5.0


.PHONEY: run dev prod clean
build:
	@wails build

run: build
	@./build/bin/${APP_NAME}.exe

dev:
	@wails dev

test:
	@go test ./...

scan:
	@golangci-lint run
	@gosec ./...

prod:
	@wails build -o ${APP_NAME}-${APP_VERSION}.exe -upx -race -nsis

clean:
	rm -f ./build/bin/${APP_NAME}-${APP_VERSION}*
