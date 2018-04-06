# library-api
RESTful API to manage a library of books

## API Endpoints
### /create
The `/create` endpoint attempts to create a new book in the library.

Arguments can be passed through the URL to set book attributes

EXAMPLE:

```
http://localhost:8080/create?Title=My%20Golang%20Adventure&Publisher=Nobody
```

The `/create` endpoint performs _some_ data validation, and will throw an error on invalid field data and / or an inability to create the book object in storage.

### /read
*TODO*
The `read` endpoint fetches a `book` object from storage.

### /update
*TODO*
The `update` endpoint allows changes to be made to existing records.

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
