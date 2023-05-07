package cocktail

type CocktailTruncated struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url" db:"image_url"`
}

type Cocktail struct {
	Id          string
	Name        string
	Recipe      string
	ImageURL    string `db:"image_url"`
	Description string
}

type CocktailsTruncatedResponse struct {
	Cocktails []CocktailTruncated `json:"cocktails"`
}

type CocktailIdRequest struct {
	Id string `json:"id,omitempty"`
}

type CocktailResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Recipe      string `json:"recipe"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type CreateCocktailRequest struct {
	Name        string `json:"name,omitempty"`
	Recipe      string `json:"recipe,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateCocktailRequest struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Recipe      string `json:"recipe,omitempty"`
	Description string `json:"description,omitempty"`
}

type CocktailIdResponse struct {
	Id string `json:"id"`
}
