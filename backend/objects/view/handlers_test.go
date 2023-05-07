package view_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/objects/view"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestAccountAndCocktailFullViewHandlers(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{
		Id:          formatters.GenerateId(),
		FirebaseUId: formatters.GenerateId(),
		Name:        "Oleg",
		ImageURL:    "1",
	}
	account2 := account.Account{
		Id:          formatters.GenerateId(),
		FirebaseUId: formatters.GenerateId(),
		Name:        "Serry",
		ImageURL:    "2",
		Description: "somedesc",
	}
	account3 := account.Account{
		Id:          formatters.GenerateId(),
		FirebaseUId: formatters.GenerateId(),
		Name:        "Karas",
		ImageURL:    "3",
	}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, account.Insert(nil, account3))

	subs1 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  account2.Id,
		Target:  account1.Id,
		Deleted: false,
	}
	subs2 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  account2.Id,
		Target:  account3.Id,
		Deleted: false,
	}
	subs3 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  account3.Id,
		Target:  account2.Id,
		Deleted: false,
	}

	require.Nil(t, subscription.Insert(nil, subs1))
	require.Nil(t, subscription.Insert(nil, subs2))
	require.Nil(t, subscription.Insert(nil, subs3))

	cocktail1 := cocktail.Cocktail{
		Id:       formatters.GenerateId(),
		Name:     "Beer",
		ImageURL: "4",
	}
	cocktail2 := cocktail.Cocktail{
		Id:       formatters.GenerateId(),
		Name:     "Wine",
		ImageURL: "5",
	}

	require.Nil(t, cocktail.Insert(nil, cocktail1))
	require.Nil(t, cocktail.Insert(nil, cocktail2))

	pearl1 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		Grade:      1,
		AccountId:  account1.Id,
		CocktailId: cocktail1.Id,
		CreatedAt:  1,
		ImageURL:   "6",
	}
	pearl2 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		Grade:      3,
		AccountId:  account2.Id,
		CocktailId: cocktail2.Id,
		CreatedAt:  2,
		ImageURL:   "7",
	}
	pearl3 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		Grade:      5,
		AccountId:  account2.Id,
		CocktailId: cocktail1.Id,
		CreatedAt:  3,
		ImageURL:   "8",
	}

	pearl1View := view.PearlViewResponse{
		Id:               pearl1.Id,
		Grade:            pearl1.Grade,
		AccountId:        pearl1.AccountId,
		CocktailId:       pearl1.CocktailId,
		CreatedAt:        pearl1.CreatedAt,
		AccountName:      account1.Name,
		CocktailName:     cocktail1.Name,
		ImageURL:         pearl1.ImageURL,
		AccountImageURL:  account1.ImageURL,
		CocktailImageURL: cocktail1.ImageURL,
	}
	pearl2View := view.PearlViewResponse{
		Id:               pearl2.Id,
		Grade:            pearl2.Grade,
		AccountId:        pearl2.AccountId,
		CocktailId:       pearl2.CocktailId,
		CreatedAt:        pearl2.CreatedAt,
		AccountName:      account2.Name,
		CocktailName:     cocktail2.Name,
		ImageURL:         pearl2.ImageURL,
		AccountImageURL:  account2.ImageURL,
		CocktailImageURL: cocktail2.ImageURL,
	}
	pearl3View := view.PearlViewResponse{
		Id:               pearl3.Id,
		Grade:            pearl3.Grade,
		AccountId:        pearl3.AccountId,
		CocktailId:       pearl3.CocktailId,
		CreatedAt:        pearl3.CreatedAt,
		AccountName:      account2.Name,
		CocktailName:     cocktail1.Name,
		ImageURL:         pearl3.ImageURL,
		AccountImageURL:  account2.ImageURL,
		CocktailImageURL: cocktail1.ImageURL,
	}

	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))
	require.Nil(t, pearl.Insert(nil, pearl3))

	// account2
	request := view.ObjectIdRequest{Id: account2.Id}
	responseAccountView, err := view.SelfAccountFullViewHandler(request)
	require.Nil(t, err)
	expectedAccountView := view.SelfAccountView{
		Account:     account.AccountResponse(account2),
		Pearls:      []view.PearlViewResponse{pearl3View, pearl2View},
		PearlsStats: pearl.PearlsStatsResponse{PearlsAmount: 2, AverageRating: 4.0},
		SubsStats:   subscription.SelfStatsResponse{SubscribersAmount: 1, SubscriptionsAmount: 2},
	}
	require.Equal(t, expectedAccountView, *responseAccountView)

	// cocktail1
	request = view.ObjectIdRequest{Id: cocktail1.Id}
	responseCocktailView, err := view.CocktailFullViewHandler(request)
	require.Nil(t, err)

	expectedCocktailView := view.CocktailView{
		Cocktail:    cocktail.CocktailResponse(cocktail1),
		Pearls:      []view.PearlViewResponse{pearl3View, pearl1View},
		PearlsStats: pearl.PearlsStatsResponse{PearlsAmount: 2, AverageRating: 3.0},
	}
	require.Equal(t, expectedCocktailView, *responseCocktailView)

	// generalAccount2
	generalRequest := view.ObjectIdWithSelfRequest{Id: account2.Id, SelfId: account1.Id}
	responseGeneralAccountView, err := view.GeneralAccountFullViewHandler(generalRequest)
	require.Nil(t, err)
	expectedGeneralAccountView := view.GeneralAccountView{
		Account:     account.AccountResponse(account2),
		Pearls:      []view.PearlViewResponse{pearl3View, pearl2View},
		PearlsStats: pearl.PearlsStatsResponse{PearlsAmount: 2, AverageRating: 4.0},
		SubsStats:   subscription.GeneralStatsResponse{SubscribersAmount: 1, SubscriptionsAmount: 2, Subscribed: false},
	}
	require.Equal(t, expectedGeneralAccountView, *responseGeneralAccountView)

}

func TestGetByNamePrefixHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: "321", Name: "NoAnguish", FirebaseUId: "5555"}
	account2 := account.Account{Id: "123", Name: "NoAng", FirebaseUId: "2222"}
	account3 := account.Account{Id: "222", Name: "Kekw", FirebaseUId: "1111"}
	account4 := account.Account{Id: "66", Name: "noangu", FirebaseUId: "24334"}
	account5 := account.Account{Id: "555", Name: "NoA", FirebaseUId: "82828"}

	sub1 := subscription.Subscription{Id: "888", Source: account1.Id, Target: account3.Id, Deleted: false}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, account.Insert(nil, account3))
	require.Nil(t, account.Insert(nil, account4))
	require.Nil(t, account.Insert(nil, account5))
	require.Nil(t, subscription.Insert(nil, sub1))

	request := view.NamePrefixRequest{
		AccountId: account1.Id,
		Prefix:    "NoA",
		Limit:     50,
	}

	response, err := view.GetByNamePrefixHandler(request)
	require.Nil(t, err)
	expected := []account.AccountResponse{
		account.AccountResponse(account5),
		account.AccountResponse(account2),
		account.AccountResponse(account4),
	}
	require.Equal(t, expected, response.Accounts)
}

func TestGetPearlEventsHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: formatters.GenerateId(), Name: "Oleg", FirebaseUId: formatters.GenerateId(), ImageURL: "1"}
	account2 := account.Account{Id: formatters.GenerateId(), Name: "Karas", FirebaseUId: formatters.GenerateId()}

	cocktail1 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Pina Colada", ImageURL: "2"}
	cocktail2 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Gin Tonic"}
	cocktail3 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Beer"}

	sub1 := subscription.Subscription{Id: "888", Source: account1.Id, Target: account2.Id, Deleted: false}
	pearl1 := pearl.Pearl{Id: "222", AccountId: account1.Id, CocktailId: cocktail1.Id, CreatedAt: 2, ImageURL: "3"}
	pearl2 := pearl.Pearl{Id: "44", AccountId: account2.Id, CocktailId: cocktail2.Id, CreatedAt: 1}
	pearl3 := pearl.Pearl{Id: "555", AccountId: account2.Id, CocktailId: cocktail3.Id, CreatedAt: 3}

	pearl1View := view.PearlView{
		Id:               pearl1.Id,
		AccountId:        pearl1.AccountId,
		CocktailId:       pearl1.CocktailId,
		CreatedAt:        pearl1.CreatedAt,
		AccountName:      account1.Name,
		CocktailName:     cocktail1.Name,
		ImageURL:         pearl1.ImageURL,
		AccountImageURL:  account1.ImageURL,
		CocktailImageURL: cocktail1.ImageURL,
	}
	pearl2View := view.PearlView{
		Id:           pearl2.Id,
		AccountId:    pearl2.AccountId,
		CocktailId:   pearl2.CocktailId,
		CreatedAt:    pearl2.CreatedAt,
		AccountName:  account2.Name,
		CocktailName: cocktail2.Name,
	}
	pearl3View := view.PearlView{
		Id:           pearl3.Id,
		AccountId:    pearl3.AccountId,
		CocktailId:   pearl3.CocktailId,
		CreatedAt:    pearl3.CreatedAt,
		AccountName:  account2.Name,
		CocktailName: cocktail3.Name,
	}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, cocktail.Insert(nil, cocktail1))
	require.Nil(t, cocktail.Insert(nil, cocktail2))
	require.Nil(t, cocktail.Insert(nil, cocktail3))
	require.Nil(t, subscription.Insert(nil, sub1))
	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))
	require.Nil(t, pearl.Insert(nil, pearl3))

	request := view.EventsRequest{
		AccountId: account1.Id,
		Limit:     50,
	}

	response, err := view.GetPearlEventsHandler(request)
	require.Nil(t, err)
	expected := []view.PearlViewResponse{
		view.PearlViewResponse(pearl3View),
		view.PearlViewResponse(pearl1View),
		view.PearlViewResponse(pearl2View),
	}
	require.Equal(t, expected, response.Pearls)
}

func TestGetAccountSubscriptionsHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: "321", Name: "NoAnguish", FirebaseUId: "5555"}
	account2 := account.Account{Id: "123", Name: "NoAng", FirebaseUId: "2222"}
	account3 := account.Account{Id: "222", Name: "Kekw", FirebaseUId: "1111"}

	sub1 := subscription.Subscription{Id: "888", Source: account1.Id, Target: account2.Id, Deleted: false}
	sub2 := subscription.Subscription{Id: "2424", Source: account1.Id, Target: account3.Id, Deleted: false}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, account.Insert(nil, account3))
	require.Nil(t, subscription.Insert(nil, sub1))
	require.Nil(t, subscription.Insert(nil, sub2))

	request := view.ObjectIdRequest{
		Id: account1.Id,
	}

	found, err := view.GetAccountSubscriptionsHandler(request)
	require.Nil(t, err)
	expected := []account.AccountResponse{account.AccountResponse(account3), account.AccountResponse(account2)}
	require.Equal(t, expected, found.Accounts)
}
