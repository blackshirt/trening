package asn

import (
	"context"

	"github.com/blackshirt/trening/models"
)

type ASNRepository interface {
	Get(ctx context.Context, id int) (models.ASN, error)
	Store(ctx context.Context, a models.ASN) (bool, error)
}
