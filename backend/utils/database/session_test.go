package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func buildSession(t *testing.T) Session {
	config, err := InitConfig()
	require.NoError(t, err)
	s := Session{
		Config: *config,
		Ctx:    context.Background(),
	}
	return s
}

func TestOpenAndCloseSession(t *testing.T) {
	s := buildSession(t)
	err := s.Open()
	require.NoError(t, err)

	err = s.Close()
	require.NoError(t, err)
}

func TestCreateTable(t *testing.T) {
	s := buildSession(t)
	_ = s.Open()
	defer s.Close()

	err := s.OpenTx(TxType{AccessMode: ReadWrite})
	require.NoError(t, err)

	err = Modify("CREATE TABLE IF NOT EXISTS test (count integer);", &s)
	require.NoError(t, err)

	err = Modify("DROP TABLE IF EXISTS test;", &s)
	require.NoError(t, err)

	err = s.Close()
	require.NoError(t, err)
}

func TestSimpleQuery(t *testing.T) {
	s := buildSession(t)
	_ = s.Open()
	defer s.Close()

	err := Modify("DROP TABLE IF EXISTS test;", &s)
	require.NoError(t, err)

	err = Modify("CREATE TABLE IF NOT EXISTS test (count integer);", &s)
	require.NoError(t, err)

	err = Modify("INSERT INTO test VALUES (1);", &s)
	require.NoError(t, err)

	err = s.CloseTx()
	require.NoError(t, err)

	val, err := Get[int]("SELECT * FROM test;", &s)
	require.NoError(t, err)
	require.Len(t, val, 1)
	require.Equal(t, int(1), val[0])
	_ = s.CloseTx()

	err = Modify("DROP TABLE IF EXISTS test;", &s)
	require.NoError(t, err)

	err = s.Close()
	require.NoError(t, err)
}
