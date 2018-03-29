# library-api
RESTful API to manage a library of books

## API Endpoints
### /create
The `create` endpoint attempts to create a new book in the library.

Every `book` attribute must be present in a `create` request _other than `uid`_ (see below).

If a `create` request does not contain the required attributes, an HTTP 400 is returned.

Upon successful creation of a new `book` object, the `uid` of the new object is returned to the client in JSON format, along with an HTTP 201 response.

### /read
The `read` endpoint fetches a `book` object from storage.

The request must include the `uid` of the desired book.

If the `uid` provided is not in storage, a HTTP 404 is returned.

If the `uid` provided can be associated with a `book` object in storage, the object is returned to the client as JSON, and an HTTP 200 is returned.

### /update
The `update` endpoint allows changes to be made to existing records.

The request must include the `uid` of the `book` object to be changed.

Any additional attributes included will overwrite the existing `book` object's attributes.

If an attribute does not conform to the correct data type (rating is > 3, publish_date is not a valid unix time, etc.), an HTTP 400 is returned.

If the `book` object is successfully updated, or no changes are made, an HTTP 200 response is returned.

### /delete
The `delete` endpoint allows `book` objects to be removed from storage.

The request must include the `uid` of the `book` object to be deleted.

If the `uid` provided does not match any `book` objects in storage, an HTTP 404 is returned.

If the `book` object is successfully deleted, an HTTP 200 response is returned.

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
