package subscription_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/subscription"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestGetInsertUpdateSubscription(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get part check
	data := subscription.Subscription{Id: "321", Source: "55", Target: "22", Deleted: false}
	err = subscription.Insert(nil, data)
	require.Nil(t, err)

	found, err := subscription.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)

	// update and get part check
	data.Deleted = true
	err = subscription.Update(nil, data)
	require.Nil(t, err)

	found, err = subscription.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)
}

func TestGetStats(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	sourceId := formatters.GenerateId()
	targetId := formatters.GenerateId()

	sub1 := subscription.Subscription{Id: "321", Source: sourceId, Target: targetId, Deleted: false}
	sub2 := subscription.Subscription{Id: "255", Source: formatters.GenerateId(), Target: targetId, Deleted: false}
	sub3 := subscription.Subscription{Id: "23232", Source: sourceId, Target: targetId, Deleted: true}
	sub4 := subscription.Subscription{Id: "12333", Source: formatters.GenerateId(), Target: targetId, Deleted: true}

	require.Nil(t, subscription.Insert(nil, sub1))
	require.Nil(t, subscription.Insert(nil, sub2))
	require.Nil(t, subscription.Insert(nil, sub3))
	require.Nil(t, subscription.Insert(nil, sub4))

	targetsAmount, err := subscription.GetSubsAmountBySourceId(nil, sourceId)
	require.Nil(t, err)
	require.Equal(t, int32(1), *targetsAmount)

	sourcesAmount, err := subscription.GetSubsAmountByTargetId(nil, targetId)
	require.Nil(t, err)
	require.Equal(t, int32(2), *sourcesAmount)
}
