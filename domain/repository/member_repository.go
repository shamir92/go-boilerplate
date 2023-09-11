package repository

import (
	"database/sql"
	"fmt"
	"log"
	entities "simple-invitation/domain/entities"
	"strings"

	"github.com/google/uuid"
)

type memberRepository struct {
	dbWriter *sql.DB
	dbReader *sql.DB
	member   entities.Member
}

type IMemberRepository interface {
	// Create a new member.
	FetchDBReader() *sql.DB
	FetchDBWriter() *sql.DB
	FetchMemberByListId(ids []uuid.UUID) ([]*entities.Member, error)
	FetchMemberById(idMember uuid.UUID) (*entities.Member, error)
	FetchMemberByEmail(email string) (*entities.Member, error)
	StoreNewMeber(email string, firstname string, lastname string) error
	FetchAll() ([]*entities.Member, error)
	// Get a member by ID.
	// GetByID(id uuid.UUID) (*Member, error)
	// Update a member.
	// Update(member *Member) error
	// Delete a member.
	// Delete(id uuid.UUID) error
}

func NewMemberRepository(dbReader *sql.DB, dbWriter *sql.DB) *memberRepository {
	return &memberRepository{
		dbWriter: dbWriter,
		dbReader: dbReader,
	}
}

func (mr *memberRepository) FetchDBWriter() *sql.DB {
	return mr.dbWriter
}

func (mr *memberRepository) FetchDBReader() *sql.DB {
	return mr.dbReader
}

func (mr *memberRepository) FetchMemberByEmail(email string) (*entities.Member, error) {
	var member entities.Member
	var id, firstname, lastname, memberemail string
	var created_at, updated_at, deleted_at sql.NullTime
	err := mr.dbReader.QueryRow("SELECT id, firstname, lastname, email,created_at, updated_at, deleted_at FROM member WHERE email=?", email).Scan(
		&id, &firstname, &lastname, &memberemail, &created_at, &updated_at, &deleted_at)
	log.Println(memberemail)

	member.Id = id
	member.Firstname = firstname
	member.Lastname = lastname
	member.Email = memberemail
	member.CreatedAt = created_at
	member.UpdatedAt = updated_at
	member.DeletedAt = deleted_at

	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (mr *memberRepository) FetchMemberById(idMember uuid.UUID) (*entities.Member, error) {
	var member entities.Member
	var id, firstname, lastname, memberemail string
	var created_at, updated_at, deleted_at sql.NullTime
	err := mr.dbReader.QueryRow("SELECT id, firstname, lastname, email,created_at, updated_at, deleted_at FROM member WHERE id=?", idMember).Scan(
		&id, &firstname, &lastname, &memberemail, &created_at, &updated_at, &deleted_at)

	member.Id = id
	member.Firstname = firstname
	member.Lastname = lastname
	member.Email = memberemail
	member.CreatedAt = created_at
	member.UpdatedAt = updated_at
	member.DeletedAt = deleted_at

	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (mr *memberRepository) FetchAll() ([]*entities.Member, error) {
	var members []*entities.Member
	result, err := mr.dbReader.Query("SELECT id, firstname, lastname, email,created_at, updated_at, deleted_at FROM member")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var member entities.Member
		var id, firstname, lastname, memberemail string
		var created_at, updated_at, deleted_at sql.NullTime
		result.Scan(&id, &firstname, &lastname, &memberemail, &created_at, &updated_at, &deleted_at)
		member.Id = id
		member.Firstname = firstname
		member.Lastname = lastname
		member.Email = memberemail
		member.CreatedAt = created_at
		member.UpdatedAt = updated_at
		member.DeletedAt = deleted_at
		members = append(members, &member)
	}

	if err != nil {
		return nil, err
	}
	return members, nil
}

func (mr *memberRepository) StoreNewMeber(email string, firstname string, lastname string) error {
	stmt, err := mr.dbWriter.Prepare("INSERT INTO member (id, firstname, lastname, email) VALUES (?,?, ?, ?)")
	if err != nil {
		return err
	}

	id := uuid.New()

	// Execute the prepared statement.
	_, err = stmt.Exec(id.String(), firstname, lastname, email)
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

func (mr *memberRepository) FetchMemberByListId(ids []uuid.UUID) ([]*entities.Member, error) {

	var uuidString []string
	for _, value := range ids {
		uuidString = append(uuidString, value.String())
	}

	sqlInquiry := strings.Join(uuidString, "\",\"")
	sqlBuilder := fmt.Sprintf("SELECT id, firstname, lastname, email,created_at, updated_at, deleted_at FROM member where id in (\"%s\")", sqlInquiry)
	var members []*entities.Member
	result, err := mr.dbReader.Query(sqlBuilder)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var id, firstname, lastname, memberemail string
		var created_at, updated_at, deleted_at sql.NullTime
		result.Scan(&id, &firstname, &lastname, &memberemail, &created_at, &updated_at, &deleted_at)
		members = append(members, &entities.Member{
			Id:        id,
			Firstname: firstname,
			Lastname:  lastname,
			Email:     memberemail,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			DeletedAt: deleted_at,
		})
	}

	if err != nil {
		return nil, err
	}
	return members, nil
}
