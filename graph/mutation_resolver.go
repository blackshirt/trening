package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type mutationResolver struct {
	service *RepoServices
}

func (s *RepoServices) Mutation() MutationResolver {
	return &mutationResolver{service: s}
}

func (m *mutationResolver) CreateOpd(ctx context.Context, input models.OpdInput) (*models.Opd, error) {
	res, err := m.service.opdRepo.OpdCreate(ctx, input)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (m *mutationResolver) CreateOrg(ctx context.Context, input models.OrgInput) (*models.Org, error) {
	res, err := m.service.orgRepo.OrgCreate(ctx, input)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}
