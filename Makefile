install : docker-alpine clean 

docker-alpine:
	GOOS=linux go build -o library_api
	docker build --pull -t gkahen/library-api .

clean:
	rm -f library_api
