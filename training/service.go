package training

type Service interface {
	PostTraining(ctx context.Context, t graph.Training) error
	GetTraining(ctx context.Context, id string) (graph.Training, error)
	GetTrainings(ctx context.Context, offset, limit int) ([]graph.Training, error)
}

type trainingService struct {
	repository Repository
}

func NewTrainingService(r Repository) Service {
	return trainingService{r}
}

func (s trainingService) PostTraining(ctx context.Context, i graph.Training) error {

}
