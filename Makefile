build-app:
	@go build -o ./.bin/app ./cmd/mainApp/main.go

run:build-app
	@./.bin/app
