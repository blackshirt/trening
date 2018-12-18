package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

/*

 */
const DEFAULT_FIRST = 40

func (o *resolver) opdList(ctx context.Context, first int, after string) (models.OpdConnection, error) {
	if first == nil {
		first = DEFAULT_FIRST
	}
	if first < 0 {
		log.Fatal()
	}
	res, err := o.service.opdRepo.OpdList(ctx, first, after)
	if err != nil {
		log.Fatalf(err)
	}
}
