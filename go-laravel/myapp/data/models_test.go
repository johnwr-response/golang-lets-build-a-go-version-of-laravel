package data

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	db2 "github.com/upper/db/v4"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()
	defer func(fakeDB *sql.DB) {
		_ = fakeDB.Close()
	}(fakeDB)

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)
	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("Wrong type", fmt.Sprintf("%T", m))
	}
	_ = os.Setenv("DATABASE_TYPE", "mysql")
	m = New(fakeDB)
	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("Wrong type", fmt.Sprintf("%T", m))
	}
}

func TestGetInsertID(t *testing.T) {
	var id db2.ID
	id = int64(1)
	returnedID := GetInsertID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("Wrong type", fmt.Sprintf("%T", returnedID))
	}
	id = int32(1)
	returnedID = GetInsertID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("Wrong type", fmt.Sprintf("%T", returnedID))
	}
}
