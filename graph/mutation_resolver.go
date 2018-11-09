package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type mutationResolver struct {
	service *GraphQLService
}

func (s *GraphQLService) Mutation() MutationResolver {
	return &mutationResolver{service: s}
}

func (m *mutationResolver) CreateOPD(ctx context.Context, input models.OPDInput) (models.OPD, error) {
	res, err := m.service.opdRepo.Insert(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}