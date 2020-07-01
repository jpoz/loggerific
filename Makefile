.PHONY: test_run

test_run:
	go build -o logmaker cmd/logmarker/main.go
	go build -o loggerific cmd/loggerific/main.go
	./logmaker | ./loggerific

test_chart:
	go run cmd/chart/main.go
