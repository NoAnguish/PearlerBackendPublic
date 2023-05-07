package account_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestGetInsertUpdateExistsAccount(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get
	data := account.Account{Id: "321", Name: "NoAnguish", Email: "kukech@gmail.com"}
	err = account.Insert(nil, data)
	require.Nil(t, err)

	found, err := account.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)

	// update and get
	data.Name = "Serrriy"
	err = account.Update(nil, data)
	require.Nil(t, err)

	found, err = account.GetById(nil, data.Id)
	require.Nil(t, err)
	require.Equal(t, data, *found)
}

func TestGetByFirebaseUId(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	// insert and get
	data := account.Account{Id: "321", Name: "NoAnguish", Email: "kukech@gmail.com", FirebaseUId: "5555"}
	err = account.Insert(nil, data)
	require.Nil(t, err)

	found, err := account.GetByFirebaseUId(nil, data.FirebaseUId)
	require.Nil(t, err)
	require.Equal(t, data, *found)

	// exists
	exists, err := account.ExistsByFirebaseUId(nil, data.FirebaseUId)
	require.Nil(t, err)
	require.True(t, exists)
}
