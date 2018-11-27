package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/org"
	"github.com/blackshirt/trening/core/trx"
	"github.com/blackshirt/trening/models"
)

type RepoServices struct {
	asnRepo asn.AsnRepo
	opdRepo opd.OpdRepo
	orgRepo org.OrgRepo
	trxRepo trx.TrxRepo
}

func NewRepoServices(asn asn.AsnRepo, opd opd.OpdRepo, org org.OrgRepo, trx trx.TrxRepo) *RepoServices {
	return &RepoServices{
		asnRepo: asn,
		opdRepo: opd,
		orgRepo: org,
		trxRepo: trx,
	}
}

type resolver struct {
	service *RepoServices
}

func (s *RepoServices) Asn() AsnResolver {
	return &resolver{
		service: s,
	}
}

func (s *RepoServices) Trx() TrxResolver {
	return &resolver{
		service: s,
	}
}
func (s *RepoServices) TrxDetail() TrxDetailResolver {
	return &resolver{
		service: s,
	}
}

func (a *resolver) CurrentPlaces(ctx context.Context, obj *models.Asn) (*models.Opd, error) {
	row, err := a.service.opdRepo.OpdById(ctx, *obj.CurrentPlaces.ID)
	if err != nil {
		log.Fatal("Error in opd get by id", err, row)
	}
	return row, nil
}

func (a *resolver) Trx(ctx context.Context, obj *models.TrxDetail) (*models.Trx, error) {
	row, err := a.service.trxRepo.Trx(ctx, obj)
	if err != nil {
		log.Fatal("Error in trx get by id", err, row)
	}
	return row, nil
}

func (a *resolver) Organizer(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	row, err := a.service.trxRepo.Organizer(ctx, obj)
	if err != nil {
		log.Fatal("Error in cat get by id", err, row)
	}
	return row, nil
}

func (a *resolver) Location(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	row, err := a.service.trxRepo.Location(ctx, obj)
	if err != nil {
		log.Fatal("Error in cat get by id", err, row)
	}
	return row, nil
}

func (a *resolver) Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error) {
	row, err := a.service.trxRepo.Category(ctx, obj)
	if err != nil {
		log.Fatal("Error in cat get by id", err, row)
	}
	return row, nil
}
func (a *resolver) Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error) {
	row, err := a.service.trxRepo.Type(ctx, obj)
	if err != nil {
		log.Fatal("Error in cat get by id", err, row)
	}
	return row, nil
}
