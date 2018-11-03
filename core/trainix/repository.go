package training

import (
	"context"

	"github.com/blackshirt/trening/models"
)

type Repository interface {
	PutTraining(ctx context.Context, t models.Training) error
	GetTraining(ctx context.Context, id string) (models.Training, error)
	GetTrainings(ctx context.Context, offset, limit int) ([]models.Training, error)
}
