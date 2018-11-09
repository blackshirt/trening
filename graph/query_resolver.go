package graph

import (
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


func (q *queryResolver) AsnList(ctx context.Context, pagination *Pagination) ([]models.ASN, error) {
	res, err := q.service.asnRepo.ASNList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *queryResolver) OpdList(ctx context.Context, pagination *Pagination) ([]models.OPD, error) {
	res, err := q.service.opdRepo.OPDList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *queryResolver) OrgList(ctx context.Context, pagination *Pagination) ([]models.Org, error) {
	res, err := q.service.orgRepo.OrgList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
