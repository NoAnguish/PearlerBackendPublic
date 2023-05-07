package database

import (
	"context"
)

func PrepareDefaultWriteSession() (*Session, error) {
	conf, err := InitConfig()
	if err != nil {
		return nil, err
	}

	s := Session{
		Ctx:    context.Background(),
		Config: *conf,
	}

	err = s.Open()
	if err != nil {
		return nil, err
	}
	err = s.OpenTx(TxType{
		AccessMode: ReadWrite,
	})

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func PrepareDefaultReadSession() (*Session, error) {
	conf, err := InitConfig()
	if err != nil {
		return nil, err
	}

	s := Session{
		Ctx:    context.Background(),
		Config: *conf,
	}
	err = s.Open()
	if err != nil {
		return nil, err
	}
	err = s.OpenTx(TxType{})

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func PrepareDefaultScanSession() (*Session, error) {
	conf, err := InitConfig()
	if err != nil {
		return nil, err
	}

	s := Session{
		Ctx:    context.Background(),
		Config: *conf,
	}
	err = s.Open()
	if err != nil {
		return nil, err
	}
	err = s.OpenTx(TxType{
		IsoLevel:   ReadCommited,
		AccessMode: ReadOnly,
	})

	if err != nil {
		return nil, err
	}

	return &s, nil
}
