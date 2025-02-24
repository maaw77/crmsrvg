package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	// ErrConStrEmty occurs when argument function NewCrmDatabase is empty.
	ErrConStrEmty = errors.New("connection string is empty")
)

// CrmDatabase is the storage for the CRM server.
type CrmDatabase struct {
	dbpool *pgxpool.Pool
}

// createPool creates a new Pool.
func (d *CrmDatabase) createPool(ctx context.Context, config *pgxpool.Config) (err error) {
	d.dbpool, err = pgxpool.NewWithConfig(ctx, config)
	return
}

// NewCrmDatabase allocates and returns a new CrmDatabase.
func NewCrmDatabase(ctx context.Context, connStr string) (crmDB *CrmDatabase, err error) {
	if connStr == "" {
		return crmDB, ErrConStrEmty
	}
	config, err := pgxpool.ParseConfig(connStr)

	if err != nil {
		return crmDB, err
	}

	crmDB = &CrmDatabase{}

	err = crmDB.createPool(ctx, config)

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
