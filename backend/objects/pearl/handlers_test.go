package pearl_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestCreatePearlHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := pearl.CreatePearlRequest{
		AccountId:  formatters.GenerateId(),
		CocktailId: formatters.GenerateId(),
		Grade:      5,
		Review:     "Sehr gut!",
	}

	response, err := pearl.CreatePearlHandler(request, nil)
	require.Nil(t, err)

	expected := pearl.Pearl{
		Id:         response.Id,
		AccountId:  request.AccountId,
		CocktailId: request.CocktailId,
		Review:     request.Review,
		Grade:      request.Grade,
	}

	found, err := pearl.GetById(nil, response.Id)
	require.Nil(t, err)
	expected.CreatedAt = found.CreatedAt
	require.Equal(t, expected, *found)
}

func TestPearlsStatsHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	cocktailId := formatters.GenerateId()
	accountId := formatters.GenerateId()

	pearl1 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		AccountId:  accountId,
		CocktailId: formatters.GenerateId(),
		ImageURL:   formatters.GenerateId(),
		Grade:      5,
		Review:     "Sehr gut!",
		CreatedAt:  1,
	}
	pearl2 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		AccountId:  accountId,
		CocktailId: cocktailId,
		ImageURL:   formatters.GenerateId(),
		Grade:      5,
		Review:     "Nicht so gut:(",
		CreatedAt:  3,
	}

	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))

	request := pearl.ObjectIdRequest{Id: accountId}
	response, err := pearl.GetStatsByAccountIdHandler(request)
	require.Nil(t, err)
	expected := pearl.PearlsStatsResponse{PearlsAmount: 2, AverageRating: 5.0}
	require.Equal(t, expected, *response)

	request = pearl.ObjectIdRequest{Id: cocktailId}
	response, err = pearl.GetStatsByCocktailIdHandler(request)
	require.Nil(t, err)
	expected = pearl.PearlsStatsResponse{PearlsAmount: 1, AverageRating: 5.0}
	require.Equal(t, expected, *response)
}
