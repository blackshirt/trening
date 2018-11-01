//go:generate gorunpkg github.com/99designs/gqlgen

package graph

import (
	context "context"

	models "github.com/blackshirt/trening/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) OPD() OPDResolver {
	return &oPDResolver{r}
}
func (r *Resolver) Organizer() OrganizerResolver {
	return &organizerResolver{r}
}
func (r *Resolver) Participant() ParticipantResolver {
	return &participantResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Training() TrainingResolver {
	return &trainingResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTraining(ctx context.Context, input TrainingInput) (models.Training, error) {
	panic("not implemented")
}

type oPDResolver struct{ *Resolver }

func (r *oPDResolver) Address(ctx context.Context, obj *models.OPD) (models.Address, error) {
	panic("not implemented")
}

type organizerResolver struct{ *Resolver }

func (r *organizerResolver) Address(ctx context.Context, obj *models.Organizer) (models.Address, error) {
	panic("not implemented")
}

type participantResolver struct{ *Resolver }

func (r *participantResolver) CurrentPlaces(ctx context.Context, obj *models.Participant) (models.OPD, error) {
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
func (r *trainingResolver) Organizer(ctx context.Context, obj *models.Training) (models.Organizer, error) {
	panic("not implemented")
}
func (r *trainingResolver) Participants(ctx context.Context, obj *models.Training) ([]models.Participant, error) {
	panic("not implemented")
}
