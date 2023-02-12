build: 
	go build -o echelon_task cmd/echelon-task/main.go
run: build
	./echelon_task