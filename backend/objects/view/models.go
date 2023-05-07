package view

import (
	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
)

type PearlView struct {
	Id               string
	AccountId        string `db:"account_id"`
	CocktailId       string `db:"cocktail_id"`
	ImageURL         string `db:"image_url"`
	Grade            int32
	Review           string
	CreatedAt        int64  `db:"created_at"`
	AccountName      string `db:"account_name"`
	CocktailName     string `db:"cocktail_name"`
	AccountImageURL  string `db:"account_image_url"`
	CocktailImageURL string `db:"cocktail_image_url"`
}

type PearlViewResponse struct {
	Id               string `json:"id"`
	AccountId        string `json:"account_id"`
	CocktailId       string `json:"cocktail_id"`
	ImageURL         string `json:"image_url"`
	Grade            int32  `json:"grade"`
	Review           string `json:"review"`
	CreatedAt        int64  `json:"created_at"`
	AccountName      string `json:"account_name"`
	CocktailName     string `json:"cocktail_name"`
	AccountImageURL  string `json:"account_image_url"`
	CocktailImageURL string `json:"cocktail_image_url"`
}

type CocktailView struct {
	Cocktail    cocktail.CocktailResponse `json:"cocktail"`
	Pearls      []PearlViewResponse       `json:"pearls"`
	PearlsStats pearl.PearlsStatsResponse `json:"pearls_statistics"`
}

type SelfAccountView struct {
	Account     account.AccountResponse        `json:"account"`
	Pearls      []PearlViewResponse            `json:"pearls"`
	PearlsStats pearl.PearlsStatsResponse      `json:"pearls_statistics"`
	SubsStats   subscription.SelfStatsResponse `json:"subscriptions_statistics"`
}

type GeneralAccountView struct {
	Account     account.AccountResponse           `json:"account"`
	Pearls      []PearlViewResponse               `json:"pearls"`
	PearlsStats pearl.PearlsStatsResponse         `json:"pearls_statistics"`
	SubsStats   subscription.GeneralStatsResponse `json:"subscriptions_statistics"`
}

type ObjectIdRequest struct {
	Id string `json:"id,omitempty"`
}

type ObjectIdWithSelfRequest struct {
	Id     string `json:"id,omitempty"`
	SelfId string `json:"self_id,omitempty"`
}

type NamePrefixRequest struct {
	AccountId string `json:"account_id"`
	Prefix    string `json:"prefix,omitempty"`
	Limit     uint   `json:"limit,string"`
}

type EventsRequest struct {
	AccountId string `json:"account_id"`
	Limit     uint   `json:"limit,string"`
}

type PearlViewsListResponse struct {
	Pearls []PearlViewResponse `json:"pearls"`
}
