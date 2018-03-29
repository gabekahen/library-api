# library-api
RESTful API to manage a library of books

## API Endpoints
### /create
The `create` endpoint will attempt to create a new book in the library, based on provided form data.

Every `book` attribute must be present in a `create` request (see below), or an HTTP 400 is returned.

Upon successful creation of a new `book` object, an HTTP 200 is returned to the client.

### /read
The `read` endpoint will fetch a `book` object from storage, and return it as JSON to the client.

The request must include the full, exact title of the book.

### /update
### /delete

## Storage Schema
### book
The `book` object represents a physical piece of literature.

`book` has the following attributes:

Attribute | Value | Description
------------ | ------------- | -------------
title | String | Name or title of the book
author | String | Name(s) of book's author(s)
publisher | String | Name of company the book was published under
publish_date | Date | Date the book was published
rating | Int | Numeric rating of book (min: 1, max: 3)
status | Int | Book's status code within the library (0: CheckedIn, 1: CheckedOut)
