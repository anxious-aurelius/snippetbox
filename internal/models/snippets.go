package models

import "time"

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// TODO: Add a SnippetModel struct with a database connection pool field.

// TODO: Implement Insert method - inserts a new snippet and returns its id.

// TODO: Implement Get method - fetches a specific snippet by id.

// TODO: Implement Latest method - fetches the 10 most recently created snippets.
