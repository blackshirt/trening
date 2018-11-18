package graph

import (
	"context"

	"github.com/blackshirt/trening/models"
)

type trxResolver struct {
	service *RepoServices
}

type trxHistory struct {
	service *RepoServices
}

func (gs *RepoServices) Trx() TrxResolver {
	return &trxResolver{
		service: gs,
	}
}

func (gs *RepoServices) TrxHistory() TrxHistoryResolver {
	return &trxHistory{service: gs}
}

func (tr *trxResolver) Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error) {
	trxCat, err := tr.service.trxCatRepo.CatById(ctx, *obj.Category.ID)
	if err != nil {
		return nil, err
	}
	return trxCat, nil
}

func (tr *trxResolver) Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error) {
	trxType, err := tr.service.trxTypeRepo.TypeById(ctx, *obj.Type.ID)
	if err != nil {
		return nil, err
	}
	return trxType, nil
}

func (tr *trxResolver) TrxCatList(ctx context.Context) ([]*models.TrxCat, error) {
	rows, err := tr.service.trxCatRepo.CatList(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (tr *trxResolver) TrxTypeList(ctx context.Context) ([]*models.TrxType, error) {
	rows, err := tr.service.trxTypeRepo.TypeList(ctx)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (th *trxHistory) Organizer(ctx context.Context, obj *models.TrxHistory) (*models.Org, error) {
	row, err := th.service.trxHistoryRepo.OrgById(obj.Organizer.ID)
	if err != nil {
		return nil, err
	}
	return row, nil
}
func (th *trxHistory) Location(ctx context.Context, obj *models.TrxHistory) (*models.Org, error) {
	row, err := th.service.trxHistoryRepo.OrgById(obj.Location.ID)
	if err != nil {
		return nil, err
	}
	return row, nil
}
func (th *trxHistory) Participants(ctx context.Context, obj *models.TrxHistory) ([]*models.Asn, error) {
	panic("not implemented")
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
	AsnList(ctx context.Context, pagination *models.Pagination) ([]*models.Asn, error)
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
