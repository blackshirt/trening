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

func (m *mutationResolver) CreateOPD(ctx context.Context, input models.OPDInput) (bool, error) {
	_, err := m.service.opdRepo.Insert(ctx, input)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func (m *mutationResolver) CreateOrg(ctx context.Context, input models.OrgInput) (bool, error) {
	_, err := m.service.orgRepo.Insert(ctx, input)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}
