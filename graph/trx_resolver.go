package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type trxResolver struct {
	service *RepoServices
}

type trxDetail struct {
	service *RepoServices
}

func (gs *RepoServices) Trx() TrxResolver {
	return &trxResolver{
		service: gs,
	}
}

func (gs *RepoServices) TrxDetail() TrxDetailResolver {
	return &trxDetail{service: gs}
}

//type TrxResolver interface {
//	Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error)
//	Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error)
//}

func (tr *trxResolver) Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error) {
	trxCat, err := tr.service.trxRepo.Category(ctx, obj)
	if err != nil {
		return nil, err
	}
	return trxCat, nil
}

func (tr *trxResolver) Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error) {
	trxType, err := tr.service.trxRepo.Type(ctx, obj)
	if err != nil {
		return nil, err
	}
	return trxType, nil
}

// trxdetail
// type TrxDetailResolver interface {
//	Trx(ctx context.Context, obj *models.TrxDetail) (*models.Trx, error)
//	Organizer(ctx context.Context, obj *models.TrxDetail) (*models.Org, error)
//	Location(ctx context.Context, obj *models.TrxDetail) (*models.Org, error)
//	Participants(ctx context.Context, obj *models.TrxDetail) ([]*models.Asn, error)
//}

func (th *trxDetail) Trx(ctx context.Context, obj *models.TrxDetail) (*models.Trx, error) {
	row, err := th.service.trxRepo.Trx(ctx, obj)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (th *trxDetail) Organizer(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	row, err := th.service.trxRepo.Organizer(ctx, obj)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (th *trxDetail) Location(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	row, err := th.service.trxRepo.Location(ctx, obj)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (th *trxDetail) Participants(ctx context.Context, obj *models.TrxDetail) ([]*models.Asn, error) {
	rows, err := th.service.trxRepo.Participants(ctx, obj)
	if err != nil {
		log.Fatal("ERRRORR BRO tddfd")
		return nil, err
	}
	return rows, nil
}
