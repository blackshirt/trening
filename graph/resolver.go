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
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) OPD() OPDResolver {
	return &oPDResolver{r}
}
func (r *Resolver) Organisasi() OrganisasiResolver {
	return &organisasiResolver{r}
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

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTraining(ctx context.Context, input TrainingInput) (models.Training, error) {
	panic("not implemented")
}

type oPDResolver struct{ *Resolver }

func (r *oPDResolver) Address(ctx context.Context, obj *models.OPD) (models.Address, error) {
	panic("not implemented")
}

type organisasiResolver struct{ *Resolver }

func (r *organisasiResolver) Address(ctx context.Context, obj *models.Organisasi) (models.Address, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Traininglist(ctx context.Context) ([]models.Training, error) {
	panic("not implemented")
}

type trainingResolver struct{ *Resolver }

func (r *trainingResolver) HeldAt(ctx context.Context, obj *models.Training) (models.Schedule, error) {
	panic("not implemented")
}
func (r *trainingResolver) Organizer(ctx context.Context, obj *models.Training) (models.Organisasi, error) {
	panic("not implemented")
}
func (r *trainingResolver) Participants(ctx context.Context, obj *models.Training) ([]models.ASN, error) {
	panic("not implemented")
}
