# default - port 8080, env development
run:
	go run ./cmd/api

run-no-limit:
	go run ./cmd/api -limiter-enabled=false

run-cors:
	go run ./cmd/api -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

example-cors:
	go run ./cmd/examples/cors/simple

kill:
	@PID=$$(lsof -ti:8080); \
	if [ ! -z "$$PID" ]; then \
		kill -9 $$PID; \
	else \
		echo "No process is using port 8080"; \
	fi