package subscription_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestCreateSubscriptionHandlerNew(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := subscription.SourceTargetRequest{
		Source: "231",
		Target: "123",
	}
	response, err := subscription.CreateSubscriptionHandler(request)
	require.Nil(t, err)

	expected := subscription.Subscription{
		Id:      response.Id,
		Source:  request.Source,
		Target:  request.Target,
		Deleted: false,
	}

	actual, err := subscription.GetById(nil, response.Id)
	require.Nil(t, err)

	require.Equal(t, expected, *actual)
}

func TestCreateSubscriptionHandlerDeleted(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	subs := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  formatters.GenerateId(),
		Deleted: true,
	}

	request := subscription.SourceTargetRequest{
		Source: subs.Source,
		Target: subs.Target,
	}

	err = subscription.Insert(nil, subs)
	require.Nil(t, err)

	response, err := subscription.CreateSubscriptionHandler(request)
	require.Nil(t, err)

	require.Equal(t, response.Id, subs.Id)

	expected := subs
	expected.Deleted = false

	found, err := subscription.GetById(nil, response.Id)
	require.Nil(t, err)

	require.Equal(t, expected, *found)
}

func TestDeleteSubscriptionHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	subs := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  formatters.GenerateId(),
		Deleted: false,
	}

	request := subscription.SourceTargetRequest{
		Source: subs.Source,
		Target: subs.Target,
	}

	err = subscription.Insert(nil, subs)
	require.Nil(t, err)

	response, err := subscription.DeleteSubscriptionHandler(request)
	require.Nil(t, err)

	require.Equal(t, response.Id, subs.Id)

	expected := subs
	expected.Deleted = true

	found, err := subscription.GetById(nil, response.Id)
	require.Nil(t, err)

	require.Equal(t, expected, *found)
}

func TestDeleteSubscriptionHandlerDeleted(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	subs := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  formatters.GenerateId(),
		Deleted: true,
	}

	request := subscription.SourceTargetRequest{
		Source: subs.Source,
		Target: subs.Target,
	}

	err = subscription.Insert(nil, subs)
	require.Nil(t, err)

	_, err = subscription.DeleteSubscriptionHandler(request)
	require.Error(t, err)
}

func TestDeleteSubscriptionHandlerDoesNotExist(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := subscription.SourceTargetRequest{
		Source: formatters.GenerateId(),
		Target: formatters.GenerateId(),
	}

	_, err = subscription.DeleteSubscriptionHandler(request)
	require.Error(t, err)
}

func TestCreateSubscriptionHandlerDuplicate(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	subs := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  formatters.GenerateId(),
		Deleted: false,
	}

	request := subscription.SourceTargetRequest{
		Source: subs.Source,
		Target: subs.Target,
	}
	err = subscription.Insert(nil, subs)
	require.Nil(t, err)

	_, err = subscription.CreateSubscriptionHandler(request)
	require.Error(t, err)
}

func TestGetStatsByUserIdHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	accountId := formatters.GenerateId()

	subs1 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  accountId,
		Deleted: false,
	}
	subs2 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  formatters.GenerateId(),
		Target:  accountId,
		Deleted: false,
	}
	subs3 := subscription.Subscription{
		Id:      formatters.GenerateId(),
		Source:  accountId,
		Target:  formatters.GenerateId(),
		Deleted: false,
	}

	require.Nil(t, subscription.Insert(nil, subs1))
	require.Nil(t, subscription.Insert(nil, subs2))
	require.Nil(t, subscription.Insert(nil, subs3))

	request := subscription.UserIdRequest{
		Id: accountId,
	}

	response, err := subscription.GetSelfStatsByUserIdHandler(request)
	require.Nil(t, err)

	expected := subscription.SelfStatsResponse{
		SubscribersAmount:   int32(2),
		SubscriptionsAmount: int32(1),
	}
	require.Equal(t, expected, *response)

	generalRequest := subscription.UserIdWithSelfRequest{
		SelfId: accountId,
		Id:     subs3.Target,
	}

	generalResponse, err := subscription.GetGeneralStatsByUserIdHandler(generalRequest)
	require.Nil(t, err)

	generalExpected := subscription.GeneralStatsResponse{
		SubscribersAmount:   int32(1),
		SubscriptionsAmount: int32(0),
		Subscribed:          true,
	}
	require.Equal(t, generalExpected, *generalResponse)
}
