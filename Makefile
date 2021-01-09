build:
	go build -o bin/golang-github-app cmd/golang-github-app/main.go
run:
	go run cmd/golang-github-app/main.go
.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo