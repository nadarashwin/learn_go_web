package models

import (
	"errors"
	"time"
)

// ErrNoRecord to be used when no model is present
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet struct that corresponds to the fields in our snippet table(DB)
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}
