package repositories

import (
	"github.com/arieefrachman/go-graphql-example/infrastructures"
	"github.com/arieefrachman/go-graphql-example/models"
	"github.com/spf13/viper"
)

type IPersonRepository interface {
	GetAll() []models.Person
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

func (r *PersonRepository) GetAll() []models.Person{
	return nil
}