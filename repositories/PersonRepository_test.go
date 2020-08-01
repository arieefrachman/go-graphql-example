package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arieefrachman/go-graphql-example/infrastructures"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"testing"
)

func TestPersonRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection\n", err)
	}

	personRepository := new(PersonRepository)
	personRepository.DB = &infrastructures.SQLConnection{
		Connection: &dbr.Connection{
			DB:            db,
			Dialect:       dialect.PostgreSQL,
			EventReceiver: nil,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date"}).
		AddRow(1, "Arif", "Rakhman", "2010-01-01").
		AddRow(2, "Dyta", "Vina", "2020-01-02")

	mock.ExpectQuery("^SELECT (.+) from persons$").WillReturnRows(rows)

}