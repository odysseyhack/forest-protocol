package domain

type RecommendationRequest struct {
	Dimensions
}

type RecommendationResponse struct {
	RecipeIndex string     `json:"recipeId"`
	Dimensions  Dimensions `json:"dimensions"`
	Flora       []Flora    `json:"flora"`
}

type Dimensions struct {
	JobOpportunities int `json:"jobOpportunities"`
	Biodiversity     int `json:"bioDiversity"`
	Irr              int `json:"irr"`
	EnvProtection    int `json:"envProtection"`
	FoodSecurity     int `json:"foodSecurity"`
}

type Flora struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
