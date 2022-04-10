clean:
	go clean

test: clean
	go test ./test

run: test
	docker-compose -f docker/docker-compose.yml rm -f transactions_routine
	docker-compose -f docker/docker-compose.yml up --build transactions_routine
