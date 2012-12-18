.PHONY: clean test maketests

checker: checker.go
	go build checker.go

clean:
	rm checker

test: checker
	go test
	time ./checker testing
	time md5sum testing/bigfile

maketests:
	mkdir testing
	dd if=/dev/zero of=testing/bigfile bs=600M count=1
