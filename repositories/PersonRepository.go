package repositories

import (
	"github.com/arieefrachman/go-graphql-example/infrastructures"
	"github.com/arieefrachman/go-graphql-example/models"
	"github.com/spf13/viper"
)

type IPersonRepository interface {
	GetAll() ([]models.Person, error)
	FindByID(ID int) (models.Person, error)
}

type PersonRepository struct {
	DB infrastructures.ISQLConnection
}

func NewPersonRepository() *PersonRepository{
	return &PersonRepository{
		DB: &infrastructures.SQLConnection{
			Connection: infrastructures.CreateConnection(
				viper.GetString("database.dialect"),
				viper.GetString("database.descriptors"),
				viper.GetInt("database.max_conn"),
				viper.GetInt("database.max_idle")),
		},
	}
}

func (r *PersonRepository) GetAll() (persons []models.Person, err error){
	db := r.DB.Connect()
	_, err = db.Select("id", "first_name", "last_name", "birth_date").From("persons").Load(&persons)
	return
}

func (r *PersonRepository) FindByID(ID int) (person models.Person, err error) {
	db := r.DB.Connect()
	err = db.Select("id", "first_name", "last_name", "birth_date").From("persons").Where("id = ?", ID).LoadOne(&person)
	return person, err
}