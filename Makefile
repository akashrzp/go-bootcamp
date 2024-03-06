build-first-tasks:
	go build -o bin/firstTasks/main cmd/firstTasks/main.go

run-first-tasks:
	go run ./cmd/firstTasks

serve-first-tasks:
	./bin/firstTasks