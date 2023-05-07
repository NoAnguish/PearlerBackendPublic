package account_test

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/objects/account"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := account.CreateAccountRequest{
		Email:       "lolkek@gmail.com",
		FirebaseUId: "123",
	}
	response, err := account.CreateAccountHandler(request)
	require.Nil(t, err)

	expected := account.Account{
		Id:          response.Id,
		Email:       request.Email,
		FirebaseUId: request.FirebaseUId,
	}

	actual, err := account.GetById(nil, response.Id)
	require.Nil(t, err)

	require.Equal(t, expected, *actual)
}

func TestCreateAccountAlreadyExists(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "1243232",
		Email:       "lolkek@gmail.com",
		FirebaseUId: "123",
		Name:        "oleg",
	}

	request := account.CreateAccountRequest{
		Email:       "lolkek1@gmail.com",
		FirebaseUId: userAccount.FirebaseUId,
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	_, err = account.CreateAccountHandler(request)
	require.Error(t, err)
}

func TestGetIdByFirebaseUIdHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "123",
		Email:       "oleg",
		FirebaseUId: "55",
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	request := account.FirebaseUIdRequest{FirebaseUId: userAccount.FirebaseUId}
	response, err := account.GetIdByFirebaseUIdHandler(request)
	require.Nil(t, err)
	require.Equal(t, userAccount.Id, response.Id)
}

func TestGetIdByFirebaseUIdHandlerUserNotExists(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	request := account.FirebaseUIdRequest{FirebaseUId: "123"}
	response, err := account.GetIdByFirebaseUIdHandler(request)
	require.Nil(t, err)
	require.Empty(t, response.Id)
}

func TestGetByIdHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "123",
		Email:       "oleg",
		FirebaseUId: "55",
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	request := account.AccountIdRequest{Id: userAccount.Id}
	response, err := account.GetByIdHandler(request)
	require.Nil(t, err)
	require.Equal(t, userAccount, account.Account(*response))
}

func TestUpdateHandler(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "123",
		Email:       "oleg",
		FirebaseUId: "55",
		Name:        "noang",
		Description: "one desc",
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	request := account.UpdateAccountRequest{
		Id:          userAccount.Id,
		Name:        "noanguish",
		Description: "some description",
	}
	response, err := account.UpdateAccountHandler(request, nil)
	require.Nil(t, err)

	userAccount.Name = request.Name
	userAccount.Description = request.Description

	found, err := account.GetById(nil, userAccount.Id)
	require.Nil(t, err)

	require.Equal(t, userAccount.Id, response.Id)
	require.Equal(t, userAccount, *found)
}

func TestUpdateHandlerPartly(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "123",
		Email:       "oleg",
		FirebaseUId: "55",
		Name:        "noang",
		Description: "one desc",
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	// if some fields not given they should not change
	request := account.UpdateAccountRequest{
		Id:          userAccount.Id,
		Description: "some description",
	}
	response, err := account.UpdateAccountHandler(request, nil)
	require.Nil(t, err)

	userAccount.Description = request.Description

	found, err := account.GetById(nil, userAccount.Id)
	require.Nil(t, err)

	require.Equal(t, userAccount.Id, response.Id)
	require.Equal(t, userAccount, *found)
}

func TestUpdateHandlerEmpty(t *testing.T) {
	err := migrations.MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	userAccount := account.Account{
		Id:          "123",
		Email:       "oleg",
		FirebaseUId: "55",
		Name:        "noang",
		Description: "one desc",
	}

	err = account.Insert(nil, userAccount)
	require.Nil(t, err)

	request := account.UpdateAccountRequest{
		Id: userAccount.Id,
	}
	response, err := account.UpdateAccountHandler(request, nil)
	require.Nil(t, err)

	// name should not change, description should be empty
	userAccount.Description = ""

	found, err := account.GetById(nil, userAccount.Id)
	require.Nil(t, err)

	require.Equal(t, userAccount.Id, response.Id)
	require.Equal(t, userAccount, *found)
}
