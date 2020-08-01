package repositories

import (
	"github.com/arieefrachman/go-graphql-example/infrastructures"
	"github.com/arieefrachman/go-graphql-example/models"
)

type IPhoneNumberRepository interface {
	GetByPersonID(ID int) ([]models.PhoneNumber, error)
}

type PhoneNumberRepository struct {
	DB infrastructures.ISQLConnection
}

func (r *PhoneNumberRepository) GetByPersonID(ID int) (numbers []models.PhoneNumber, err error){
	db := r.DB.Connect()

	_, err = db.SelectBySql("select * from phone_number where person_id = ?", ID).Load(&numbers)
	return 
}


