package sqlstore

import "database/sql"

type stmtErr struct {
	st  *sql.Stmt
	err error
}

type rowErr struct {
	err error
	row *sql.Row
}

func (r *rowErr) Scan(dest ...interface{}) error {
	if r.err != nil {
		return r.err
	}
	return r.row.Scan(dest...)
}

func (st *stmtErr) Exec(args ...interface{}) (sql.Result, error) {
	if st.err != nil {
		return nil, st.err
	}
	return st.st.Exec(args...)
}

func (st *stmtErr) Query(args ...interface{}) (*sql.Rows, error) {
	if st.err != nil {
		return nil, st.err
	}
	return st.st.Query(args...)
}

func (st *stmtErr) QueryRow(args ...interface{}) *rowErr {
	if st.err == nil {
		return &rowErr{err: st.err, row: st.st.QueryRow(args...)}
	}
	return &rowErr{err: st.err, row: nil}
}

func (st *stmtErr) Close() {
	st.st.Close()
}
