# blogging-platform-api

Practice project written in Go based on https://roadmap.sh/projects/blogging-platform-api.

Do not use this as reference to structure a Go project.

## Installation

Clone the project and run it in the root project directory with
```go run .```


## Usage
By default runs on ```localhost:8080``` (hardcoded lol)

```POST /posts``` - Creates a blog with given parameters (seen below).

```PUT /posts/<id>``` - Updates existing post with given parameters (seen below).

```DELETE /posts/<id>``` - Deletes post with given id.

```GET /posts/<id>``` - Returns post with given id.

```GET /posts``` - Returns all posts. Optionally accepts a search term in the form of ```GET /posts?term=<term>``` that returns any blog that contains the term in the title/content/category.

## Blog Format
```json
{
  "title": "Example post",
  "content": "This is the body of the post.",
  "category": "Test",
  "tags": ["Testing", "English"]
}
```

## Additional Info
Used SQlite for database, goose for managing migrations and sqlc for working with sql queries. Everything else is standard stuff.

Feel free to use this for any purpose, but the code quality is pretty bad so I wouldn't recommend it ^^
