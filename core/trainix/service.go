package training

type Service interface {
	PostTraining(ctx context.Context, t graph.Training) error
	GetTraining(ctx context.Context, id string) (graph.Training, error)
	GetTrainings(ctx context.Context, offset, limit int) ([]graph.Training, error)
}
