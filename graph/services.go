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


// RepoServices implement ResolverRoot interface for generated graphql runtime
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

func (s *RepoServices) Query() QueryResolver {
	return &resolver{
		service: s,
	}
}

func (s *RepoServices) Mutation() MutationResolver {
	return &resolver{service: s}
}



// Service implementation call underlying repository backed by sql database
// Asn resolver
func (a *resolver) CurrentPlaces(ctx context.Context, obj *models.Asn) (*models.Opd, error) {
	row, err := a.service.opdRepo.OpdById(ctx, *obj.CurrentPlaces.ID)
	if err != nil {
		log.Fatal("Error in opd get by id", err, row)
	}
	return row, nil
}

// Trx detail resolver
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


// Trx resolver
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


// Query resolver
func (q *resolver) AsnList(ctx context.Context) ([]*models.Asn, error) {
	res, err := q.service.asnRepo.AsnList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *resolver) OpdList(ctx context.Context) ([]*models.Opd, error) {
	res, err := q.service.opdRepo.OpdList(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (q *resolver) OrgList(ctx context.Context) ([]*models.Org, error) {
	res, err := q.service.orgRepo.OrgList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (q *resolver) TrxList(ctx context.Context) ([]*models.TrxDetail, error) {
	rows, err := q.service.trxRepo.TrxList(ctx)
	if rows == nil {
		log.Fatal("Error rows", rows)
	}
	if err != nil {
		log.Fatal("ERRRORR BRO", rows, err)
		return nil, err
	}
	return rows, nil
}


// Mutation resolver
func (m *resolver) CreateOpd(ctx context.Context, input models.OpdInput) (*models.Opd, error) {
	res, err := m.service.opdRepo.OpdCreate(ctx, input)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (m *resolver) CreateOrg(ctx context.Context, input models.OrgInput) (*models.Org, error) {
	res, err := m.service.orgRepo.OrgCreate(ctx, input)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}




