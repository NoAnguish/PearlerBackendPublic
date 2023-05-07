package pearl_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/pearl"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestGetInsertPearl(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get part check
	data := pearl.Pearl{
		Id:         formatters.GenerateId(),
		AccountId:  formatters.GenerateId(),
		CocktailId: formatters.GenerateId(),
		ImageURL:   formatters.GenerateId(),
		Grade:      5,
		Review:     "Sehr gut!",
		CreatedAt:  formatters.GetTimestap(),
	}

	err = pearl.Insert(nil, data)
	require.Nil(t, err)

	found, err := pearl.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)
}

func TestGetByAccountAndCocktailIds(t *testing.T) {
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
	pearl3 := pearl.Pearl{
		Id:         formatters.GenerateId(),
		AccountId:  formatters.GenerateId(),
		CocktailId: cocktailId,
		ImageURL:   formatters.GenerateId(),
		Grade:      5,
		Review:     "Na",
		CreatedAt:  2,
	}

	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))
	require.Nil(t, pearl.Insert(nil, pearl3))

	found, err := pearl.GetByAccountId(nil, accountId)
	require.Nil(t, err)
	expected := []pearl.Pearl{pearl2, pearl1}
	require.Equal(t, expected, *found)

	found, err = pearl.GetByCocktailId(nil, cocktailId)
	require.Nil(t, err)
	expected = []pearl.Pearl{pearl2, pearl3}
	require.Equal(t, expected, *found)
}

func TestGetPearlStats(t *testing.T) {
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
		Grade:      4,
		Review:     "Nicht so gut:(",
		CreatedAt:  3,
	}

	require.Nil(t, pearl.Insert(nil, pearl1))
	require.Nil(t, pearl.Insert(nil, pearl2))

	found, err := pearl.GetStatsByCocktailId(nil, cocktailId)
	require.Nil(t, err)
	expected := pearl.PearlStats{PearlsAmount: 1, PearlsGradesSum: 4}
	require.Equal(t, expected, *found)

	found, err = pearl.GetStatsByAccountId(nil, accountId)
	require.Nil(t, err)
	expected = pearl.PearlStats{PearlsAmount: 2, PearlsGradesSum: 9}
	require.Equal(t, expected, *found)
}

func TestGetPearlStatsEmpty(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	cocktailId := formatters.GenerateId()
	accountId := formatters.GenerateId()

	found, err := pearl.GetStatsByCocktailId(nil, cocktailId)
	require.Nil(t, err)
	expected := pearl.PearlStats{PearlsAmount: 0, PearlsGradesSum: 0}
	require.Equal(t, expected, *found)

	found, err = pearl.GetStatsByAccountId(nil, accountId)
	require.Nil(t, err)
	expected = pearl.PearlStats{PearlsAmount: 0, PearlsGradesSum: 0}
	require.Equal(t, expected, *found)
}
