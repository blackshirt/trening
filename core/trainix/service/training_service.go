package training

import (
	"context"

	"github.com/blackshirt/trening/models"
)

type trainingService struct {
	repository Repository
}

func NewTrainingService(r Repository) Service {
	return trainingService{r}
}

func (s trainingService) PostTraining(ctx context.Context, i models.Training) error {

}
