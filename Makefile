build-first-tasks:
	go build -o bin/firstTasks/main cmd/firstTasks/main.go

run-first-tasks:
	go run ./cmd/firstTasks

serve-first-tasks:
	./bin/firstTasks

build-second-tasks:
	go build -o bin/secondTasks/main cmd/secondTasks/main.go

run-second-tasks:
	go run ./cmd/secondTasks

serve-second-tasks:
	./bin/secondTasks
