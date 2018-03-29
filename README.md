# library-api
RESTful API to manage a library of books

## API Endpoints
### /create
The `create` endpoint will attempt to create a new book in the library, based on provided form data.

Every `book` attribute must be present in a `create` request _other than `uid`_ (see below).

If a `create` request does not contain the required attributes, an HTTP 400 is returned.

Upon successful creation of a new `book` object, the `uid` of the new object is returned to the client in JSON format, along with an HTTP 200 response.

### /read
The `read` endpoint will fetch a `book` object from storage.

The request must include the `uid` of the desired book.

If the `uid` provided is not in storage, a HTTP 404 is returned.

If the `uid` provided can be associated with a `book` object in storage, the object is returned to the client as JSON, and an HTTP 200 is returned.

### /update
### /delete

## Storage Schema
### book
The `book` object represents a physical piece of literature.

`book` has the following attributes:

Attribute | Value | Description
------------ | ------------- | -------------
uid | Int | Unique identifier 
title | String | Name or title of the book
author | String | Name(s) of book's author(s)
publisher | String | Name of company the book was published under
publish_date | Date | Date the book was published
rating | Int | Numeric rating of book (min: 1, max: 3)
status | Int | Book's status code within the library (0: CheckedIn, 1: CheckedOut)
