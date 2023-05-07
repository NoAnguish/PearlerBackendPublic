package migrations

import (
	"testing"

	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/stretchr/testify/require"
)

func TestMakeCoreMigration(t *testing.T) {
	err := dropDatabase()
	require.NoError(t, err)
	s, err := database.PrepareDefaultWriteSession()
	require.NoError(t, err)

	err = createVersionTable(s)
	require.NoError(t, err)

	setDbVersion(s, "v_00000")
	s.Close()
	MakeCoreMigration(tables.GetTableData())

}
