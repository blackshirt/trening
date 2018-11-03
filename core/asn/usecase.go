package asn

import (
	"context"

	"github.com/blackshirt/trening/models"
)

type ASNUsecase interface {
	Get(ctx context.Context, id int) (models.ASN, error)
}
