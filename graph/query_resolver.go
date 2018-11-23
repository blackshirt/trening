package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type queryResolver struct {
	service *RepoServices
}

func (s *RepoServices) Query() QueryResolver {
	return &queryResolver{
		service: s,
	}
}

func (q *queryResolver) AsnList(ctx context.Context) ([]*models.Asn, error) {
	res, err := q.service.asnRepo.AsnList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *queryResolver) OpdList(ctx context.Context) ([]*models.Opd, error) {

	res, err := q.service.opdRepo.OpdList(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil

}

func (q *queryResolver) OrgList(ctx context.Context) ([]*models.Org, error) {

	res, err := q.service.orgRepo.OrgList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *queryResolver) TrxList(ctx context.Context) ([]*models.TrxDetail, error) {
	rows, err := q.service.trxRepo.TrxList(ctx)
	if err != nil {
		log.Fatal("ERRRORR BRO")
		return nil, err
	}
	return rows, nil
}
