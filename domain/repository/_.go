package member

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/google/uuid"

    "github.com/go-sql-driver/mysql"
)

// Member represents a member of the system.
type Member struct {
    ID        uuid.UUID `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// MemberRepository is an interface for interacting with the member database.
type MemberRepository interface {
    // Create a new member.
    Create(member *Member) error

    // Get a member by ID.
    GetByID(id uuid.UUID) (*Member, error)

    // Update a member.
    Update(member *Member) error

    // Delete a member.
    Delete(id uuid.UUID) error
}

// MemberRepositoryImpl is a concrete implementation of the MemberRepository interface.
type MemberRepositoryImpl struct {
    db *sql.DB
}

// NewMemberRepository creates a new MemberRepositoryImpl instance.
func NewMemberRepository(db *sql.DB) *MemberRepositoryImpl {
    return &MemberRepositoryImpl{
        db: db,
    }
}

// Create a new member.
func (r *MemberRepositoryImpl) Create(member *Member) error {
    // Create a prepared statement.
    stmt, err := r.db.Prepare("INSERT INTO members (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }

    // Execute the prepared statement.
    _, err = stmt.Exec(member.ID, member.Name, member.Email, member.CreatedAt, member.UpdatedAt)
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

// Get a member by ID.
func (r *MemberRepositoryImpl) GetByID(id uuid.UUID) (*Member, error) {