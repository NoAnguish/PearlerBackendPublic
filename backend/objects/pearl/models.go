package pearl

type Pearl struct {
	Id         string
	AccountId  string `db:"account_id"`
	CocktailId string `db:"cocktail_id"`
	ImageURL   string `db:"image_url"`
	Grade      int32
	Review     string
	CreatedAt  int64 `db:"created_at"`
}

type PearlStats struct {
	PearlsAmount    int `db:"pearls_amount"`
	PearlsGradesSum int `db:"pearls_grades_sum"`
}

type CreatePearlRequest struct {
	AccountId  string `json:"account_id,omitempty"`
	CocktailId string `json:"cocktail_id,omitempty"`
	Grade      int32  `json:"grade,omitempty"`
	Review     string `json:"review,omitempty"`
}

type ObjectIdRequest struct {
	Id string `json:"id,omitempty"`
}

type PearlResponse struct {
	Id         string `json:"id"`
	AccountId  string `json:"account_id"`
	CocktailId string `json:"cocktail_id"`
	ImageURL   string `json:"image_url"`
	Grade      int32  `json:"grade"`
	Review     string `json:"review"`
	CreatedAt  int64  `json:"created_at"`
}

type PearlIdResponse struct {
	Id string `json:"id"`
}

type PearlIdImageURLResponse struct {
	Id       string `json:"id"`
	ImageURL string `json:"image_url"`
}

type PearlsListResponse struct {
	Pearls []PearlResponse `json:"pearls"`
}

type PearlsStatsResponse struct {
	PearlsAmount  int     `json:"pearls_amount"`
	AverageRating float64 `json:"average_rating"`
}
