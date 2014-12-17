package backend

import "database/sql"

const (
	CONFLICT  = iota
	DUPLICATE = iota
	MISSING   = iota
)

type WintsError struct {
	Kind int
	Err  error
}

func (w *WintsError) Error() string {
	return w.Err.Error()
}

func SingleUpdate(db *sql.DB, errNoUpdate error, q string, args ...interface{}) error {
	res, err := db.Exec(q, args...)
	if err != nil {
		return err
	}
	nb, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nb != 1 {
		return errNoUpdate
	}
	return nil
}
