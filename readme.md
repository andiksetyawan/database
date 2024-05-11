### Install

```go get -u github.com/andiksetyawan/database```

#### Example:
```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andiksetyawan/database"
	"github.com/andiksetyawan/database/sqlx"
)

type user struct {
	ID   int
	Name string
}

func main() {
	ctx := context.TODO()

	db, err := sqlx.New(sqlx.WithPostgres(database.Config{
		Database: "foo",
		User:     "postgres",
		Host:     "localhost",
		Port:     "5432",
		Password: "secret",
	}))
	if err != nil {
		log.Fatal(fmt.Errorf("fail to connecting psql database : %w", err))
	}

	//with sqlx.Tx impl
	tx, _ := db.StartTransaction(ctx)
	_, _ = getUserRepo(ctx, tx)
	tx.Commit()

	//with sqlx.Db impl
	_, _ = getUserRepo(ctx, db.GetMaster())
}

func getUserRepo(ctx context.Context, q sqlx.Q) (user user, err error) {
	err = q.GetContext(ctx, &user, "SELECT * FROM user")
	return
}

```
