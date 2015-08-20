package datastore

import "database/sql"

//TxErr is a structure to manage transactions more confortablely.
//It maintains the ongoing error to state if the transaction must be committed or rollbacked
type TxErr struct {
	tx  *sql.Tx
	err error
}

//newTxErr begins a new transaction
func newTxErr(db *sql.DB) TxErr {
	tx, err := db.Begin()
	return TxErr{tx: tx, err: err}
}

//Done terminates the transaction.
//It is committed if err == nil. Rollbacked otherwise.
func (r *TxErr) Done() error {
	if r.err != nil {
		return r.tx.Rollback()
	}
	return r.tx.Commit()
}

//Exec executes the query and store the resulting error variable
func (r *TxErr) Exec(query string, args ...interface{}) {
	if r.err != nil {
		return
	}
	_, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = err
	}
}

//Exec executes the update query, store the resulting error variable and returns
//the number of updated rows
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
