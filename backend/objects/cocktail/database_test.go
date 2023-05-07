package cocktail_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestGetInsertUpdateCocktail(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get part check
	data := cocktail.Cocktail{
		Id:          "1",
		Name:        "old fashion",
		Recipe:      "80 whiskey\n3 dash angostura\nsugar cube",
		ImageURL:    "123",
		Description: "classic whiskey cocktail",
	}
	err = cocktail.Insert(nil, data)
	require.Nil(t, err)

	found, err := cocktail.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)

	// update and get part check
	data.Description = "I like whiskey"
	err = cocktail.Update(nil, data)
	require.Nil(t, err)

	found, err = cocktail.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)
}

func TestGetAll(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get part check
	oldFashion := cocktail.Cocktail{
		Id:          "1",
		Name:        "old fashion",
		Recipe:      "80 whiskey\n3 dash angostura\nsugar cube",
		ImageURL:    "123",
		Description: "classic whiskey cocktail",
	}
	err = cocktail.Insert(nil, oldFashion)
	require.Nil(t, err)

	pinaColada := cocktail.Cocktail{
		Id:          "2",
		Name:        "pina colada",
		Recipe:      "30 rum\n30 coconut syrop\n90 pineapple juice",
		ImageURL:    "321",
		Description: "classic tropical cocktail",
	}
	err = cocktail.Insert(nil, pinaColada)
	require.Nil(t, err)

	oldFashionTruncated := cocktail.CocktailTruncated{
		Id:       oldFashion.Id,
		Name:     oldFashion.Name,
		ImageURL: oldFashion.ImageURL,
	}

	pinaColadaTruncated := cocktail.CocktailTruncated{
		Id:       pinaColada.Id,
		Name:     pinaColada.Name,
		ImageURL: pinaColada.ImageURL,
	}

	found, err := cocktail.GetAll(nil)
	require.Nil(t, err)
	expected := []cocktail.CocktailTruncated{oldFashionTruncated, pinaColadaTruncated}
	require.ElementsMatch(t, expected, *found)
}
