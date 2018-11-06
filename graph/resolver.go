//go:generate gorunpkg github.com/99designs/gqlgen

package graph

import (
	context "context"

	models "github.com/blackshirt/trening/models"
)

type Resolver struct{}

func (r *Resolver) ASN() ASNResolver {
	return &aSNResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Training() TrainingResolver {
	return &trainingResolver{r}
}

type aSNResolver struct{ *Resolver }

func (r *aSNResolver) CurrentPlaces(ctx context.Context, obj *models.ASN) (models.OPD, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) AsnList(ctx context.Context) ([]models.ASN, error) {
	panic("not implemented")
}
func (r *queryResolver) OpdList(ctx context.Context) ([]models.OPD, error) {
	panic("not implemented")
}
func (r *queryResolver) OrgList(ctx context.Context) ([]models.Organisasi, error) {
	panic("not implemented")
}

type trainingResolver struct{ *Resolver }

func (r *trainingResolver) Organizer(ctx context.Context, obj *models.Training) (models.Organisasi, error) {
	panic("not implemented")
}
func (r *trainingResolver) Location(ctx context.Context, obj *models.Training) (models.Organisasi, error) {
	panic("not implemented")
}
func (r *trainingResolver) Participants(ctx context.Context, obj *models.Training) ([]models.ASN, error) {
	panic("not implemented")
}
