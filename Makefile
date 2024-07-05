# default - port 4000, env development
run:
	go run ./cmd/api

kill:
	@PID=$$(lsof -ti:8080); \
	if [ ! -z "$$PID" ]; then \
		kill -9 $$PID; \
	else \
		echo "No process is using port 8080"; \
	fi