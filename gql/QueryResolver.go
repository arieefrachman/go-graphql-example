package gql

import (
	"github.com/arieefrachman/go-graphql-example/repositories"
	"github.com/graphql-go/graphql"
)

type QueryResolver struct {
	PersonRepository repositories.IPersonRepository
}

type IQueryResolver interface {
	ListPerson(params graphql.ResolveParams)(interface{}, error)
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{
		PersonRepository: repositories.NewPersonRepository(),
	}
}

func (r *QueryResolver) ListPerson(params graphql.ResolveParams)(interface{}, error) {
	data, err := r.PersonRepository.GetAll()
	return data, err
}
