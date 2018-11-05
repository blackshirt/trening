package asn

import (
	"context"
	"time"

	"github.com/blackshirt/trening/models"
)

type ASNUsecase interface {
  Get(ctx context.Context, id int) (models.ASN, error)
}


type asnUcase struct {
  asnRepo ASNRepository
  opdRepo opd.OPDRepository
  ctxTimeout time.Duration
}


func NewASNUcase(a ASNRepository, o opd.OPDRepository, timeout  time.Duration) ASNUsecase {
  return &asnUcase{
    asnRepo: a,
    opdRepo: o,
    ctxTimeout: timeout,
  }
}


func (a asnUcase) GetByID(c context.Context, id int) (models.ASN, error) {
  ctx, cancel := context.WithTimeout(c, a.ctxTimeout)
  defer cancel()
  
  asn, err := a.asnRepo.GetByID(ctx, id)
  if err != nil {
    return nil, err
  }
  
  opd, err := a.opdRepo.GetByID(ctx, asn.CurrentPlaces.ID)
  if err != nil {
    return nil, err
  }
  
  asn.CurrentPlaces = *opd
  
  return asn, nil
}



