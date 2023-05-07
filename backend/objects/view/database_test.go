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

func TestGetByPrefixName(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: "321", Name: "NoAnguish", FirebaseUId: "5555"}
	account2 := account.Account{Id: "123", Name: "NoAng", FirebaseUId: "2222"}
	account3 := account.Account{Id: "222", Name: "Kekw", FirebaseUId: "1111"}
	account4 := account.Account{Id: "66", Name: "noangu", FirebaseUId: "24334"}
	account5 := account.Account{Id: "555", Name: "NoA", FirebaseUId: "82828"}

	sub1 := subscription.Subscription{Id: "888", Source: account1.Id, Target: account2.Id, Deleted: false}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, account.Insert(nil, account3))
	require.Nil(t, account.Insert(nil, account4))
	require.Nil(t, account.Insert(nil, account5))
	require.Nil(t, subscription.Insert(nil, sub1))

	found, err := view.GetByNamePrefix(nil, account1.Id, "NoA", 50)
	require.Nil(t, err)
	expected := []account.Account{account4, account5}
	require.ElementsMatch(t, expected, *found)

	found, err = view.GetByNamePrefix(nil, account1.Id, "NoA", 1)
	require.Nil(t, err)
	require.Equal(t, 1, len(*found))
	require.Equal(t, account5, (*found)[0])
}

func TestGetPearlEvents(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: formatters.GenerateId(), Name: "Oleg", FirebaseUId: formatters.GenerateId(), ImageURL: "1"}
	account2 := account.Account{Id: formatters.GenerateId(), Name: "Karas", FirebaseUId: formatters.GenerateId(), ImageURL: "2"}

	cocktail1 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Pina Colada", ImageURL: "3"}
	cocktail2 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Gin Tonic", ImageURL: "4"}
	cocktail3 := cocktail.Cocktail{Id: formatters.GenerateId(), Name: "Beer", ImageURL: "5"}

	sub1 := subscription.Subscription{Id: "888", Source: account1.Id, Target: account2.Id, Deleted: false}
	pearl1 := pearl.Pearl{Id: "222", AccountId: account1.Id, CocktailId: cocktail1.Id, CreatedAt: 2, ImageURL: "23232"}
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
		CocktailImageURL: cocktail1.ImageURL,
		AccountImageURL:  account1.ImageURL,
	}
	pearl2View := view.PearlView{
		Id:               pearl2.Id,
		AccountId:        pearl2.AccountId,
		CocktailId:       pearl2.CocktailId,
		CreatedAt:        pearl2.CreatedAt,
		AccountName:      account2.Name,
		CocktailName:     cocktail2.Name,
		CocktailImageURL: cocktail2.ImageURL,
		AccountImageURL:  account2.ImageURL,
	}
	pearl3View := view.PearlView{
		Id:               pearl3.Id,
		AccountId:        pearl3.AccountId,
		CocktailId:       pearl3.CocktailId,
		CreatedAt:        pearl3.CreatedAt,
		AccountName:      account2.Name,
		CocktailName:     cocktail3.Name,
		CocktailImageURL: cocktail3.ImageURL,
		AccountImageURL:  account2.ImageURL,
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

	found, err := view.GetPearlEvents(nil, account1.Id, 50)
	require.Nil(t, err)
	expected := []view.PearlView{pearl3View, pearl1View, pearl2View}
	require.Equal(t, expected, *found)

	found, err = view.GetPearlEvents(nil, account1.Id, 1)
	require.Nil(t, err)
	expected = []view.PearlView{pearl3View}
	require.Equal(t, expected, *found)
}

func TestGetAccountSubscriptions(t *testing.T) {
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

	found, err := view.GetAccountSubscriptions(nil, account1.Id)
	require.Nil(t, err)
	expected := []account.Account{account3, account2}
	require.Equal(t, expected, *found)
}

func TestGetFilledPearls(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	account1 := account.Account{Id: "321", Name: "NoAnguish", FirebaseUId: "5555", ImageURL: "1"}
	account2 := account.Account{Id: "22", Name: "LolKek", FirebaseUId: "5555"}
	cocktail1 := cocktail.Cocktail{Id: "585", Name: "Beer", ImageURL: "2"}
	cocktail2 := cocktail.Cocktail{Id: "222332", Name: "Wine"}

	pearl1 := pearl.Pearl{
		Id:         "222",
		AccountId:  account1.Id,
		CocktailId: cocktail1.Id,
		CreatedAt:  2,
		Grade:      5,
		Review:     "Gut!",
	}
	pearl2 := pearl.Pearl{Id: "44", AccountId: account2.Id, CocktailId: cocktail1.Id, CreatedAt: 1}
	pearl3 := pearl.Pearl{Id: "22232234", AccountId: account1.Id, CocktailId: cocktail2.Id, CreatedAt: 3}

	require.Nil(t, account.Insert(nil, account1))
	require.Nil(t, account.Insert(nil, account2))
	require.Nil(t, cocktail.Insert(nil, cocktail1))
	require.Nil(t, cocktail.Insert(nil, cocktail2))
	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))
	require.Nil(t, pearl.Insert(nil, pearl3))

	pearl1View := view.PearlView{
		Id:               pearl1.Id,
		AccountId:        pearl1.AccountId,
		CocktailId:       pearl1.CocktailId,
		CreatedAt:        pearl1.CreatedAt,
		Grade:            pearl1.Grade,
		Review:           pearl1.Review,
		AccountName:      account1.Name,
		CocktailName:     cocktail1.Name,
		AccountImageURL:  account1.ImageURL,
		CocktailImageURL: cocktail1.ImageURL,
	}
	pearl2View := view.PearlView{
		Id:               pearl2.Id,
		AccountId:        pearl2.AccountId,
		CocktailId:       pearl2.CocktailId,
		CreatedAt:        pearl2.CreatedAt,
		AccountName:      account2.Name,
		CocktailName:     cocktail1.Name,
		CocktailImageURL: cocktail1.ImageURL,
	}
	pearl3View := view.PearlView{
		Id:              pearl3.Id,
		AccountId:       pearl3.AccountId,
		CocktailId:      pearl3.CocktailId,
		CreatedAt:       pearl3.CreatedAt,
		AccountName:     account1.Name,
		CocktailName:    cocktail2.Name,
		AccountImageURL: account1.ImageURL,
	}

	found, err := view.GetFilledPearlsByCocktailId(nil, cocktail1.Id)
	require.Nil(t, err)
	expected := []view.PearlView{pearl1View, pearl2View}
	require.Equal(t, expected, *found)

	found, err = view.GetFilledPearlsByAccountId(nil, account1.Id)
	require.Nil(t, err)
	expected = []view.PearlView{pearl3View, pearl1View}
	require.Equal(t, expected, *found)
}
