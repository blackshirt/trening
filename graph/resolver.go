//go:generate gorunpkg github.com/99designs/gqlgen

package graph

import (
	context "context"

	"github.com/segmentio/ksuid"
)

type Resolver struct {
	trlist []Training
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTraining(ctx context.Context, input TrainingInput) (Training, error) {
	tr := Training{
		ID:          ksuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
	}
	r.trlist = append(r.trlist, tr)
	return tr, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Traininglist(ctx context.Context) ([]Training, error) {
	return r.trlist, nil
}
