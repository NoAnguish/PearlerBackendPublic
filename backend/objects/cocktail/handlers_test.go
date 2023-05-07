package cocktail_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/cocktail"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestGetAllCocktailsHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

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

	response, err := cocktail.GetAllHandler()
	require.Nil(t, err)
	expected := []cocktail.CocktailTruncated{oldFashionTruncated, pinaColadaTruncated}
	require.ElementsMatch(t, expected, response.Cocktails)
}

func TestGetByIdHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	pinaColada := cocktail.Cocktail{
		Id:          "2",
		Name:        "pina colada",
		Recipe:      "30 rum\n30 coconut syrop\n90 pineapple juice",
		ImageURL:    "321",
		Description: "classic tropical cocktail",
	}
	err = cocktail.Insert(nil, pinaColada)
	require.Nil(t, err)

	expected := cocktail.CocktailResponse{
		Id:          "2",
		Name:        "pina colada",
		Recipe:      "30 rum\n30 coconut syrop\n90 pineapple juice",
		ImageURL:    "321",
		Description: "classic tropical cocktail",
	}

	request := cocktail.CocktailIdRequest{Id: pinaColada.Id}
	response, err := cocktail.GetByIdHandler(request)
	require.Nil(t, err)
	require.Equal(t, expected, *response)
}

func TestCreateCocktailHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := cocktail.CreateCocktailRequest{
		Name:        "Pina colada",
		Recipe:      "30 rum\n30 coconut syrop\n90 pineapple juice",
		Description: "Nice summer cocktail",
	}

	response, err := cocktail.CreateHandler(request, nil)
	require.Nil(t, err)

	expected := cocktail.Cocktail{
		Id:          response.Id,
		Name:        request.Name,
		Recipe:      request.Recipe,
		Description: request.Description,
	}

	found, err := cocktail.GetById(nil, response.Id)
	require.Nil(t, err)
	require.Equal(t, expected, *found)
}

func TestUpdateCocktailHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	pinaColanda := cocktail.Cocktail{
		Id:          formatters.GenerateId(),
		Name:        "Pina colada",
		Recipe:      "30 rum\n30 coconut syrop\n90 pineapple juice",
		Description: "Nice summer cocktail",
	}

	err = cocktail.Insert(nil, pinaColanda)
	require.Nil(t, err)

	request := cocktail.UpdateCocktailRequest{
		Id:          pinaColanda.Id,
		Recipe:      "New Recipe",
		Description: "Some new desc",
	}

	response, err := cocktail.UpdateHandler(request)
	require.Nil(t, err)

	found, err := cocktail.GetById(nil, response.Id)

	pinaColanda.Description = request.Description
	pinaColanda.Recipe = request.Recipe

	require.Equal(t, pinaColanda, *found)
}
