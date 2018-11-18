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

func (q *queryResolver) OpdList(ctx context.Context, pagination *models.Pagination) ([]*models.Opd, error) {
	if &pagination == nil {
		pagination = &models.Pagination{
			Limit:  100,
			Offset: 0,
		}
	}
	res, err := q.service.opdRepo.OpdList(ctx, pagination.Limit, pagination.Offset)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil

}

func (q *queryResolver) OrgList(ctx context.Context, pagination *models.Pagination) ([]*models.Org, error) {
	if &pagination == nil {
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

func (q *queryResolver) TrxCatList(ctx context.Context) ([]*models.TrxCat, error) {
	rows, err := q.service.trxCatRepo.CatList(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (q *queryResolver) TrxTypeList(ctx context.Context) ([]*models.TrxType, error) {
	rows, err := q.service.trxTypeRepo.TypeList(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

/*
type ResolverRoot interface {
	Asn() AsnResolver
	Mutation() MutationResolver
	Query() QueryResolver
	Trx() TrxResolver
	TrxHistory() TrxHistoryResolver
}
type AsnResolver interface {
	CurrentPlaces(ctx context.Context, obj *models.Asn) (*models.Opd, error)
}
type MutationResolver interface {
	CreateOpd(ctx context.Context, input models.OpdInput) (*models.Opd, error)
	CreateOrg(ctx context.Context, input models.OrgInput) (*models.Org, error)
}
type QueryResolver interface {
	AsnList(ctx context.Context) ([]*models.Asn, error)
	OpdList(ctx context.Context, pagination *models.Pagination) ([]*models.Opd, error)
	OrgList(ctx context.Context, pagination *models.Pagination) ([]*models.Org, error)
	TrxCatList(ctx context.Context) ([]*models.TrxCat, error)
	TrxTypeList(ctx context.Context) ([]*models.TrxType, error)
}
type TrxResolver interface {
	Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error)
	Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error)
}
type TrxHistoryResolver interface {
	Organizer(ctx context.Context, obj *models.TrxHistory) (*models.Org, error)
	Location(ctx context.Context, obj *models.TrxHistory) (*models.Org, error)
	Participants(ctx context.Context, obj *models.TrxHistory) ([]*models.Asn, error)
}
*/
