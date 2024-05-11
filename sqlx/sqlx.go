package sqlx

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlxdb struct {
	driver string
	dsn    string
	master *sqlx.DB
	//slave *sqlx.DB
}

func New(opt ...OptFunc) (DB, error) {
	s := new(sqlxdb)
	for _, fn := range opt {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	db, err := sqlx.Connect(s.driver, s.dsn)
	if err != nil {
		return nil, err
	}

	s.master = db
	return s, nil
}

func (s sqlxdb) GetMaster() *sqlx.DB {
	return s.master
}

func (s sqlxdb) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return s.master.BeginTxx(ctx, nil)
}
