.PHONY: test_run

test_run:
	go build -o logmaker cmd/logmarker/main.go
	go build -o loggerific cmd/loggerific/main.go
	./logmaker | ./loggerific
