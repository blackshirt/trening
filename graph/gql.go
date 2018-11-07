package graph

import (
	"context"

	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/orgz"
	"github.com/blackshirt/trening/models"
)

type GraphQLService struct {
	asnRepo asn.ASNRepository
	opdRepo opd.OPDRepository
	orgRepo orgz.OrgRepository
}

func NewGraphQLService() *GraphQLService {
	return &GraphQLService{
		asnRepo: "",
		opdRepo: "",
		orgRepo: "",
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

type queryResolver struct {
	service *GraphQLService
}

func (s *GraphQLService) Query() QueryResolver {
	return &queryResolver{
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

func (a *asnResolver) CurrentPlaces(ctx context.Context, obj *models.ASN) (models.OPD, error) {
	opd, err := a.service.opdRepo.GetByID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return opd, nil
}

func (q *queryResolver) AsnList(ctx context.Context) ([]models.ASN, error) {
	res, err := q.service.asnRepo.ASNList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (q *queryResolver) OpdList(ctx context.Context) ([]models.OPD, error) {
	res, err := q.service.opdRepo.OPDList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (q *queryResolver) OrgList(ctx context.Context) ([]models.Orgz, error) {
	res, err := q.service.orgRepo.OrgList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *trainingResolver) Organizer(ctx context.Context, obj *models.Training) (models.Orgz, error) {
	orgz, err := t.service.orgRepo.GetByID(ctx, obj.Organizer.ID)
	if err != nil {
		return nil, err
	}
	return orgz, nil
}

func (t *trainingResolver) Location(ctx context.Context, obj *models.Training) (models.Orgz, error) {
	orgz, err := t.service.orgRepo.GetByID(ctx, obj.Location.ID)
	if err != nil {
		return nil, err
	}
	return orgz, nil

}

func (t *trainingResolver) Participants(ctx context.Context, obj *models.Training) ([]models.ASN, error) {
	panic("not implemented")
}
