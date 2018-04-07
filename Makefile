install : setup build docker-alpine clean

setup:
	go get github.com/docker/go-healthcheck
	go get github.com/go-sql-driver/mysql

build:
	GOOS=linux go build -o library_api

docker-alpine:
	docker build --pull -t gkahen/library-api .

clean:
	rm -f library_api
