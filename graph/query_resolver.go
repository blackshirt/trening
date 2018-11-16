package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type queryResolver struct {
	service *GraphQLService
}

func (s *GraphQLService) Query() QueryResolver {
	return &queryResolver{
		service: s,
	}
}

func (q *queryResolver) AsnList(ctx context.Context, pagination *models.Pagination) ([]*models.ASN, error) {
	res, err := q.service.asnRepo.ASNList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *queryResolver) OpdList(ctx context.Context, pagination *models.Pagination) ([]*models.OPD, error) {
	if pagination == nil {
		pagination = &models.Pagination{
			Limit:  100,
			Offset: 0,
		}
	}

	res, err := q.service.opdRepo.OPDList(ctx, pagination.Limit, pagination.Offset)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil

}

func (q *queryResolver) OrgList(ctx context.Context, pagination *models.Pagination) ([]*models.Org, error) {
	if pagination == nil {
		pagination = &models.Pagination{
			Limit:  100,
			Offset: 0,
		}
	}
	res, err := q.service.orgRepo.OrgList(ctx, pagination.Limit, pagination.Offset)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
