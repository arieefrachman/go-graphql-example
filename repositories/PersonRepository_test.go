package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arieefrachman/go-graphql-example/infrastructures"
	"github.com/arieefrachman/go-graphql-example/models"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockDB() (*PersonRepository, sqlmock.Sqlmock){
	db, mock, _ := sqlmock.New()
	personRepository := new(PersonRepository)
	personRepository.DB = &infrastructures.SQLConnection{
		Connection: &dbr.Connection{
			DB:            db,
			Dialect:       dialect.PostgreSQL,
			EventReceiver: &dbr.NullEventReceiver{},
		},
	}
	return personRepository, mock
}

func TestPersonRepository(t *testing.T) {
	t.Run("Get all data", func(t *testing.T) {
		repo, mock := mockDB()

		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date"}).
			AddRow(1, "Arif", "Rakhman", "2010-01-01").
			AddRow(2, "Dyta", "Vina", "2020-01-02")

		mock.ExpectQuery("SELECT id, first_name, last_name, birth_date FROM persons").WillReturnRows(rows)

		persons, err := repo.GetAll()

		expectedResults := []models.Person{
			{
				ID:        1,
				FirstName: "Arif",
				LastName:  "Rakhman",
				BirthDate: "2010-01-01",
			},
			{
				ID:        2,
				FirstName: "Dyta",
				LastName:  "Vina",
				BirthDate: "2020-01-02",
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedResults, persons)
	})

	t.Run("Find by ID", func(t *testing.T) {
		repo, mock := mockDB()

		ID := 1
		row := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date"}).
			AddRow(1, "Arif", "Rakhman", "2010-01-01")
		mock.ExpectQuery("SELECT id, first_name, last_name, birth_date FROM persons WHERE*").WillReturnRows(row)

		person, err := repo.FindByID(ID)

		expectedResult := models.Person{
			ID:        1,
			FirstName: "Arif",
			LastName:  "Rakhman",
			BirthDate: "2010-01-01",
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, person)
	})
}