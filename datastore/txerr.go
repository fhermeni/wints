package datastore

import "database/sql"

type TxErr struct {
	tx  *sql.Tx
	err error
}

func newTxErr(db *sql.DB) TxErr {
	tx, err := db.Begin()
	return TxErr{tx: tx, err: err}
}

func (r *TxErr) Done() error {
	if r.err != nil {
		return r.tx.Rollback()
	}
	return r.tx.Commit()
}

func (r *TxErr) Exec(query string, args ...interface{}) {
	if r.err != nil {
		return
	}
	_, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = err
	}
}

func (r *TxErr) Update(query string, args ...interface{}) int64 {
	if r.err != nil {
		return -1
	}
	res, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = err
		return -1
	}
	nb, err := res.RowsAffected()
	if err != nil {
		r.err = err
		return -1
	}
	return nb
}
