package migrations

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/stretchr/testify/require"
)

func TestSetAndGetDbVersion(t *testing.T) {
	err := MakeTestMigration(tables.GetTableData())
	require.NoError(t, err)

	err = createVersionTable(nil)
	require.NoError(t, err)

	rand.Seed(time.Now().Unix())
	seedVers := "v_" + fmt.Sprint(rand.Int())

	setDbVersion(nil, seedVers)

	expected := getDbVersion(nil)
	require.NoError(t, err)
	require.Equal(t, expected, seedVers)
}
