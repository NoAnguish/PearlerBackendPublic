package database

import (
	postgre "github.com/jackc/pgx/v4"
)

const (
	ReadCommited postgre.TxIsoLevel = postgre.ReadCommitted
	Serializable postgre.TxIsoLevel = postgre.Serializable

	ReadOnly  postgre.TxAccessMode = postgre.ReadOnly
	ReadWrite postgre.TxAccessMode = postgre.ReadWrite
)

type TxType struct {
	IsoLevel   postgre.TxIsoLevel
	AccessMode postgre.TxAccessMode
}

func (t TxType) toOptions() postgre.TxOptions {
	return postgre.TxOptions{
		IsoLevel:   t.isoLevel(),
		AccessMode: t.accessMode(),
	}
}

func (t TxType) accessMode() postgre.TxAccessMode {
	if t.AccessMode != "" {
		return t.AccessMode
	}

	return ReadOnly
}

func (t TxType) isoLevel() postgre.TxIsoLevel {
	if t.IsoLevel != "" {
		return t.IsoLevel
	}

	return Serializable
}
