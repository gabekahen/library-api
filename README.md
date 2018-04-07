# library-api
RESTful API to manage a library of books

# Setup / Install
Make sure the [Golang binaries](https://golang.org/doc/install) are present on your system.
[Docker](https://docs.docker.com/install/) is required to build the Docker image.
[Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube) is not required, but can be used to test Kubernetes deployments.

The golang dependancies ([go-sql-driver/mysql](https://github.com/docker/go-healthcheck), [docker/go-healthcheck](https://github.com/docker/go-healthcheck)) can be installed using `make setup`

The `library-api` binary can be built in the working directory using `make build`.

The `gkahen/library-api` container is based on Alpine. The image can be pulled from [Dockerhub](https://hub.docker.com/r/gkahen/library-api/), or built using `make install`.

A script `mysql-testserv.sh` has been provided to stand up a local MySQL server in Docker for testing.

# Kubernetes
This repository contains a Kubernetes manifest file `library-api.yaml`. This will create the MySQL and library-api deployments, pods, and services.

## API Endpoints
### /create
The `/create` endpoint creates a new book in the library.

Arguments can be passed through URL parameters:

Title=<string>
Author=<string>
Publisher=<string>
PublishDate=<int> - REQUIRED, unix-date in seconds (1523142945)
Rating=<int> - numbers below -127 or above 127 will truncated.
Status=<int> - numbers below -127 or above 127 will truncated.

Title, Author, and Publisher will be set to "" if left blank. Rating, Status will be set to 0.

EXAMPLE:

```
curl -v -G http://localhost:8080/create \
  --data-urlencode "Title=My Golang Adventure" \
  --data-urlencode "Publisher=Nobody" \
  --data-urlencode "PublishDate=1523142945"
```

If the book was created successfully, an HTTP 201 status code is returned, and the UID of the created book is returned in JSON format

RESPONSE:

```
{UID: 1}
```

### /read/
The `/read/` endpoint fetches a book object from storage, given a UID.

EXAMPLE:

```
curl -v -G http://localhost:8080/read/1
```

If the book UID was found in storage, an HTTP 200 status code is returned, and the book object is returned in JSON format.

RESPONSE:

```
{"UID":1,"Title":"My Golang Adventure","Author":"","Publisher":"Nobody","PublishDate":"2018-04-07T23:15:45Z","Rating":0,"Status":0}
```

### /update
The `/update` endpoint allows changes to a book's rating or status.

Arguments can be passed through URL parameters:

|Key=Value|Details|
|---------|-------|
`Rating=<int>` | numbers below -127 or above 127 will truncated.
`Status=<int>` | numbers below -127 or above 127 will truncated.

EXAMPLE:

```
curl -v -G http://localhost:8080/update \
  --data-urlencode "UID=1" \
  --data-urlencode "Status=1"
```

If the update is successfull, an HTTP 200 status is returned, along with the updated object in JSON format.

RESPONSE:

```
{"UID":1,"Title":"My Golang Adventure","Author":"","Publisher":"Nobody","PublishDate":"2018-04-07T23:15:45Z","Rating":0,"Status":1}
```

### /delete
The `/delete/` endpoint allows `book` objects to be removed from storage.

The request URL should include the `UID` of the book to be deleted.

EXAMPLE:

```
http://localhost:8080/delete/LLJvMjsQPn0t79mIhOvLjPkbjTw=
```

The `/delete/` endpoint will throw an error if it's unable to delete the object from storage.

## Storage Schema
### book
The `book` object represents a physical piece of literature.

`book` has the following attributes:

Attribute | Value | Description
------------ | ------------- | -------------
uid | String | auto-generated unique identifier
title | String | Name or title of the book
author | String | Name(s) of book's author(s)
publisher | String | Name of company the book was published under
publish_date | time.Time | Date the book was published (Golang time library)
rating | Int | Numeric rating of book (min: 1, max: 3)
status | Int | Book's status code within the library (0: CheckedIn, 1: CheckedOut)

TODO: Bulk object CRUD?
