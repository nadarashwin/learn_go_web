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

	// The place holer ? works in MySQL and for PostgreSQL it is $n
	// $1, $2, $3 ... ans so on
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Exec method to execute the statement. Ti will return a sql.Result object that contain
	// some basic information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// LastInsertID method will get the ID of our newly inserted record in snippets table.
	// this is not supported by PostgreSQL -> https://github.com/lib/pq/issues/24
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, hence convert to an int type.
	return int(id), nil
}

// Get a snippet from the database
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct
	s := &models.Snippet{}

	// row.Scan() will copy the vlaues from each field in sql.Row to the corresponding files in
	// Snippet struct.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Latest n snippet from the database
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires from snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialize a empty slice to hold the modesl.Snippets objects.
	snippets := []*models.Snippet{}

	// rows.Next() is used to iterated through the rows in the resulset.
	// it prepares the first(and then each subsequent) row to be acted on by the
	// rows.Scan() method. resultset closes itself after iterating through all the rows.
	for rows.Next() {

		// A pointed to a new zeroed Snippet struct.
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)

		// When the rows.Next() loop has finished we call rows.Err() to retrieve any error
		// that was encountered during the iteration.
		if err = rows.Err(); err != nil {
			return nil, err
		}

	}
	return snippets, nil
}
