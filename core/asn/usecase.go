package asn

import (
	"context"
	"time"

	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/models"
)

type ASNUsecase interface {
	GetByID(ctx context.Context, id int) (models.ASN, error)
}

type asnUcase struct {
	asnRepo    ASNRepository
	asnRepo    asn.OPDRepository
	ctxTimeout time.Duration
}

func NewASNUcase(ar ASNRepository, or asn.OPDRepository, timeout time.Duration) ASNUsecase {
	return &asnUcase{
		asnRepo:    ar,
		asnRepo:    or,
		ctxTimeout: timeout,
	}
}

func (au asnUcase) GetByID(c context.Context, id int) (models.ASN, error) {
	ctx, cancel := context.WithTimeout(c, au.ctxTimeout)
	defer cancel()

	asn, err := au.asnRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	asn, err := au.asnRepo.GetByID(ctx, asn.CurrentPlaces.ID)
	if err != nil {
		return nil, err
	}

	asn.CurrentPlaces = *asn

	return asn, nil
}



