package view

import (
	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
)

func CocktailFullViewHandler(request ObjectIdRequest) (*CocktailView, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	cocktail, err := cocktail.GetByIdHandler(cocktail.CocktailIdRequest(request))
	if err != nil {
		return nil, err
	}

	stats, err := pearl.GetStatsByCocktailIdHandler(pearl.ObjectIdRequest(request))
	if err != nil {
		return nil, err
	}

	cocktailPearls, err := GetFilledPearlsByCocktailId(s, request.Id)
	if err != nil {
		return nil, err
	}

	pearlsResponse := make([]PearlViewResponse, len(*cocktailPearls))

	for i, currentPearl := range *cocktailPearls {
		pearlsResponse[i] = PearlViewResponse(currentPearl)
	}

	response := CocktailView{
		Cocktail:    *cocktail,
		PearlsStats: *stats,
		Pearls:      pearlsResponse,
	}
	return &response, nil
}

func SelfAccountFullViewHandler(request ObjectIdRequest) (*SelfAccountView, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	userAccount, err := account.GetByIdHandler(account.AccountIdRequest(request))
	if err != nil {
		return nil, err
	}

	pearlsStats, err := pearl.GetStatsByAccountIdHandler(pearl.ObjectIdRequest(request))
	if err != nil {
		return nil, err
	}

	subsStats, err := subscription.GetSelfStatsByUserIdHandler(subscription.UserIdRequest(request))
	if err != nil {
		return nil, err
	}

	accountPearls, err := GetFilledPearlsByAccountId(s, request.Id)
	if err != nil {
		return nil, err
	}

	pearlsResponse := make([]PearlViewResponse, len(*accountPearls))

	for i, currentPearl := range *accountPearls {
		pearlsResponse[i] = PearlViewResponse(currentPearl)
	}

	response := SelfAccountView{
		Account:     *userAccount,
		PearlsStats: *pearlsStats,
		SubsStats:   *subsStats,
		Pearls:      pearlsResponse,
	}
	return &response, nil
}

func GeneralAccountFullViewHandler(request ObjectIdWithSelfRequest) (*GeneralAccountView, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	userAccount, err := account.GetByIdHandler(account.AccountIdRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}

	pearlsStats, err := pearl.GetStatsByAccountIdHandler(pearl.ObjectIdRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}

	subsStats, err := subscription.GetGeneralStatsByUserIdHandler(subscription.UserIdWithSelfRequest(request))
	if err != nil {
		return nil, err
	}

	accountPearls, err := GetFilledPearlsByAccountId(s, request.Id)
	if err != nil {
		return nil, err
	}

	pearlsResponse := make([]PearlViewResponse, len(*accountPearls))

	for i, currentPearl := range *accountPearls {
		pearlsResponse[i] = PearlViewResponse(currentPearl)
	}

	response := GeneralAccountView{
		Account:     *userAccount,
		PearlsStats: *pearlsStats,
		SubsStats:   *subsStats,
		Pearls:      pearlsResponse,
	}
	return &response, nil
}

func GetByNamePrefixHandler(request NamePrefixRequest) (*account.AccountsListResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	accounts, err := GetByNamePrefix(s, request.AccountId, request.Prefix, request.Limit)
	if err != nil {
		return nil, err
	}

	response := account.AccountsListResponse{Accounts: make([]account.AccountResponse, len(*accounts))}
	for i, userAccount := range *accounts {
		response.Accounts[i] = account.AccountResponse(userAccount)
	}

	return &response, nil
}

func GetPearlEventsHandler(request EventsRequest) (*PearlViewsListResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	pearls, err := GetPearlEvents(s, request.AccountId, request.Limit)
	if err != nil {
		return nil, err
	}

	response := PearlViewsListResponse{Pearls: make([]PearlViewResponse, len(*pearls))}
	for i, currentPearl := range *pearls {
		response.Pearls[i] = PearlViewResponse(currentPearl)
	}

	return &response, nil
}

func GetAccountSubscriptionsHandler(request ObjectIdRequest) (*account.AccountsListResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	accounts, err := GetAccountSubscriptions(s, request.Id)
	if err != nil {
		return nil, err
	}

	response := account.AccountsListResponse{Accounts: make([]account.AccountResponse, len(*accounts))}
	for i, userAccount := range *accounts {
		response.Accounts[i] = account.AccountResponse(userAccount)
	}

	return &response, nil
}
