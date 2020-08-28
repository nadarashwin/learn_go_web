package mysql

import (
	"database/sql"

	"github.com/nadarashwin/learn_go_web/pkg/models"
)

// SnippetModel type that wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get a snippet from the database
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest n snippet from the database
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
