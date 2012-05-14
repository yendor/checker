.PHONY: clean

checker: checker.go
	go build checker.go

clean:
	rm checker
