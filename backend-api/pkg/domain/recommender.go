package domain

type RecommendationRequest struct {
	Dimensions Dimensions
	Latitude   float32
	Longitude  float32
	Username   string
}

type RecommenderResponse struct {

}

type Dimensions struct {
	Biodiversity     int
	EnvProtection    int
	FoodSecurity     int
	JobOpportunities int
	IRR              int
}

