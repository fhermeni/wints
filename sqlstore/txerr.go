package sqlstore

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
		if r.tx == nil {
			//In case there was an issue while opening the transaction
			return r.err
		}
		if e := r.tx.Rollback(); e != nil {
			return mapCstrToError(e)
		}
		return mapCstrToError(r.err)

	}
	err := r.tx.Commit()
	return mapCstrToError(err)
}

//Exec executes the query and store the resulting error variable
func (r *TxErr) Exec(query string, args ...interface{}) {
	if r.err != nil {
		return
	}
	if _, err := r.tx.Exec(query, args...); err != nil {
		r.err = mapCstrToError(err)
	}
}

//QueryRow execute the query that must returns one row if the transaction is not in a failed state
func (r *TxErr) QueryRow(query string, args ...interface{}) *rowErr {
	if r.err == nil {
		return &rowErr{err: r.err, row: r.tx.QueryRow(query, args...)}
	}
	return &rowErr{err: r.err, row: nil}
}

//Update executes the update query, store the resulting error variable and returns
//the number of updated rows
func (r *TxErr) Update(query string, args ...interface{}) int64 {
	if r.err != nil {
		return -1
	}
	res, err := r.tx.Exec(query, args...)
	if err != nil {
		r.err = mapCstrToError(err)
		return -1
	}
	nb, err := res.RowsAffected()
	if err != nil {
		r.err = mapCstrToError(err)
		return -1
	}
	return nb
}
