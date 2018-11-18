package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

type asnResolver struct {
	service *RepoServices
}

func (gs *RepoServices) Asn() AsnResolver {
	return &asnResolver{
		service: gs,
	}
}

func (ar *asnResolver) CurrentPlaces(ctx context.Context, obj *models.Asn) (*models.Opd, error) {
	opd, err := ar.service.opdRepo.OpdById(ctx, obj.CurrentPlaces.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return opd, nil
}
