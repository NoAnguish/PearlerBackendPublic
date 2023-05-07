package database

import (
	"context"
	"fmt"

	postgre "github.com/jackc/pgx/v4"
)

type Session struct {
	connector *postgre.Conn
	Ctx       context.Context
	Config    PostgreConfig
	tx        postgre.Tx
}

func (s *Session) Open() error {
	var err error

	s.connector, err = postgre.ConnectConfig(s.Ctx, &s.Config.connConfig)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) OpenTx(txType TxType) error {
	var err error
	if s.tx != nil {
		return fmt.Errorf("tx is already open")
	}

	s.tx, err = s.connector.BeginTx(s.Ctx, txType.toOptions())
	return err
}

func (s *Session) CancelTx() error {
	var err error
	if s.tx == nil {
		return fmt.Errorf("no tx to close")
	}

	err = s.tx.Rollback(s.Ctx)
	s.tx = nil
	return err
}

func (s *Session) CloseTx() error {
	var err error
	if s.tx == nil {
		return fmt.Errorf("no tx to close")
	}

	err = s.tx.Commit(s.Ctx)
	s.tx = nil
	return err
}

func (s *Session) simpleModify(query string) error {
	_, err := s.tx.Exec(s.Ctx, query)
	return err
}

func (s *Session) modify(query string) error {
	if s.tx == nil {
		err := s.OpenTx(TxType{AccessMode: ReadWrite})

		if err != nil {
			return err
		}
	}

	err := s.simpleModify(query)

	if err != nil {
		_ = s.CancelTx()
		return err
	}
	return nil
}

func (s *Session) simpleGet(query string) (postgre.Rows, error) {
	rows, err := s.tx.Query(s.Ctx, query)
	return rows, err
}

func (s *Session) get(query string) (postgre.Rows, error) { // TODO(noangusih) think about better impl
	if s.tx == nil {
		err := s.OpenTx(TxType{AccessMode: ReadOnly})

		if err != nil {
			return nil, err
		}
	}

	rows, err := s.simpleGet(query)

	if err != nil {
		_ = s.CancelTx()
		return nil, err
	}

	return rows, nil
}

func (s *Session) Close() error {
	var err error

	if s.tx != nil {
		err = s.CloseTx()
		if err != nil {
			_ = s.CancelTx()
			return err
		}
	}

	err = s.connector.Close(s.Ctx)
	return err
}
