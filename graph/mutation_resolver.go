package graph

import (
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
	opd, err := m.service.opdRepo.Insert(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return opd, nil
}
