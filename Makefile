tidy:
	go mod tidy
run: tidy
	go run .

mysql: misc/setup_test_db.sh

test: tidy mysql
	go test . -v
