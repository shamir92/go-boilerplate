package mysqlwriter

import (
    "database/sql"
    "fmt"
    "log"
    "time"
)

// Interface for writing data to a MySQL database.
type Writer interface {
    // Write a single row of data to the database.
    WriteRow(row []interface{}) error

    // Write a batch of rows of data to the database.
    WriteBatch(rows [][]interface{}) error

    // Close the connection to the database.
    Close()
}

// Struct that implements the Writer interface.
type WriterImpl struct {
    db *sql.DB
}

// NewWriter creates a new WriterImpl instance.
func NewWriter(db *sql.DB) *WriterImpl {
    return &WriterImpl{
        db: db,
    }
}

// WriteRow writes a single row of data to the database.
func (w *WriterImpl) WriteRow(row []interface{}) error {
    // Create a prepared statement.
    stmt, err := w.db.Prepare("INSERT INTO table_name (column1, column2, ...) VALUES (?, ?, ...)")
    if err != nil {
        return err
    }

    // Execute the prepared statement.
    _, err = stmt.Exec(row...)
    if err != nil {
        return err
    }

    // Close the prepared statement.
    err = stmt.Close()
    if err != nil {
        return err
    }

    return nil
}

// WriteBatch writes a batch of rows of data to the database.
func (w *WriterImpl) WriteBatch(rows [][]interface{}) error {
    // Create a transaction.
    tx, err := w.db.Begin()
    if err != nil {
        return err
    }

    // Create a prepared statement.
    stmt, err := tx.Prepare("INSERT INTO table_name (column1, column2, ...) VALUES (?, ?, ...)")
    if err != nil {
        return err
    }

    // Execute the prepared statement for each row.
    for _, row := range rows {
        _, err = stmt.Exec(row...)
        if err != nil {
            return err
        }
    }

    // Close the prepared statement.
    err = stmt.Close()
    