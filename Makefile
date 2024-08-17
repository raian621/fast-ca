.PHONY: watch-test
watch-test:
	go run github.com/mitranim/gow test ./... -v -coverprofile=coverage.out

