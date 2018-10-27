//go:generate gorunpkg github.com/99designs/gqlgen

package resolver

import (
	context "context"

	entity "github.com/blackshirt/trening/entity"
	graph "github.com/blackshirt/trening/graph"
)

type Resolver struct{}

func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Traininglist(ctx context.Context) ([]entity.TrainingItem, error) {
	panic("not implemented")
}
