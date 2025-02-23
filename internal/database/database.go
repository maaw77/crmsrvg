package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrmDatabase struct {
	dbpool *pgxpool.Pool
}

// creates a new Pool.
func (d *CrmDatabase) CreatePool(ctx context.Context, config *pgxpool.Config) (err error) {
	d.dbpool, err = pgxpool.NewWithConfig(ctx, config)
	return
}

// NewCrmDatabase allocates and returns a new CrmDatabase.
func NewCrmDatabase() (crmDB *CrmDatabase, err error) {

	return

}

// func Datebase() {

// 	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer dbpool.Close()

// 	var greeting string
// 	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(greeting)
// }
