package recommender

import (
	"bitbucket.org/unchain/interfaces/pkg/logger"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/domain"
)

type RecommenderService struct {
	logger logger.Logger
	cfg *backend_api.Config
}

func New(logger logger.Logger, cfg *backend_api.Config) (*RecommenderService) {
	return &RecommenderService{
		logger: logger, cfg: cfg,
	}
}

func (s *RecommenderService) GetRecommendation(request domain.RecommendationRequest) (error) {
	return nil
}

