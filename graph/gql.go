package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/org"
	"github.com/blackshirt/trening/models"
)

type GraphQLService struct {
	asnRepo asn.ASNRepository
	opdRepo opd.OPDRepository
	orgRepo org.OrgRepository
}

func NewGraphQLService(asn asn.ASNRepository, opd opd.OPDRepository, org org.OrgRepository) *GraphQLService {
	return &GraphQLService{
		asnRepo: asn,
		opdRepo: opd,
		orgRepo: org,
	}
}

type asnResolver struct {
	service *GraphQLService
}

func (s *GraphQLService) ASN() ASNResolver {
	return &asnResolver{
		service: s,
	}
}

type trainingResolver struct {
	service *GraphQLService
}

func (s *GraphQLService) Training() TrainingResolver {
	return &trainingResolver{
		service: s,
	}
}

func (a *asnResolver) CurrentPlaces(ctx context.Context, obj *models.ASN) (*models.OPD, error) {
	opd, err := a.service.opdRepo.GetByID(ctx, obj.CurrentPlaces.ID)
	if err != nil {
		log.Fatal(err)
	}
	return opd, nil
}

func (t *trainingResolver) Organizer(ctx context.Context, obj *models.Training) (*models.Org, error) {
	org, err := t.service.orgRepo.GetByID(ctx, obj.Organizer.ID)
	if err != nil {
		log.Fatal(err)
	}
	return org, nil
}

func (t *trainingResolver) Location(ctx context.Context, obj *models.Training) (*models.Org, error) {
	org, err := t.service.orgRepo.GetByID(ctx, obj.Location.ID)
	if err != nil {
		log.Fatal(err)
	}
	return org, nil

}

func (t *trainingResolver) Participants(ctx context.Context, obj *models.Training) ([]*models.ASN, error) {
	panic("not implemented")
}
