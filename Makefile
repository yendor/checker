.PHONY: clean test

checker: checker.go
	go build checker.go

clean:
	rm checker

test:
	go test
	time ./checker testing
