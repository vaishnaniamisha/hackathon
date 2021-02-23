GOSRC=./...
COVERFILE=.cover.out

run:
	@go run main.go

test:
	@go test -coverprofile $(COVERFILE) `go list $(GOSRC)`


