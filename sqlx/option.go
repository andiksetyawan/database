package sqlx

import (
	"fmt"

	"github.com/andiksetyawan/database"
)

type OptFunc func(*sqlxdb) error

func WithPostgres(config database.Config) OptFunc {
	return func(s *sqlxdb) (err error) {
		s.driver = "postgres"
		s.dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)
		return
	}
}

func WithMysql(config database.Config) OptFunc {
	return func(s *sqlxdb) (err error) {
		s.driver = "mysql"
		s.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?sslmode=%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
			config.SSLMode,
		)
		return
	}
}
