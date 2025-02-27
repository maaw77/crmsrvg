package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maaw77/crmsrvg/internal/models"
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
func (c *CrmDatabase) createPool(ctx context.Context, config *pgxpool.Config) (err error) {
	c.dbpool, err = pgxpool.NewWithConfig(ctx, config)
	return
}

// getIdOrCreateAuxilTable returns the ID if the record exists, otherwise creates the record and returns its ID.
// It's for the nameTable
func (c *CrmDatabase) getIdOrCreateAuxilTable(ctx context.Context, nameTable, valRecord string) (id models.IdEntry, err error) {
	statmentGetId := fmt.Sprintf("SELECT id FROM %s WHERE name = $1;", nameTable)
	statmentCreate := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1) RETURNING id;", nameTable)

	err = c.dbpool.QueryRow(ctx, statmentGetId, valRecord).Scan(&(id.ID))

	if errors.Is(err, pgx.ErrNoRows) {
		err = c.dbpool.QueryRow(ctx, statmentCreate, valRecord).Scan(&(id.ID))
	}
	return
}

// delRowAuxilTable deletes the row with the specified id from the nameTable and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) delRowAuxilTable(ctx context.Context, nameTable string, id int) (statusExec bool, err error) {
	statmentDel := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", nameTable)
	comT, err := c.dbpool.Exec(ctx, statmentDel, id)
	statusExec = comT.RowsAffected() == 1
	return
}

// rGetIdOrCreateSites returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Sites table.
func (c *CrmDatabase) GetIdOrCreateSites(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "sites", valRecord)
}

// delRowAuxilTable deletes the row with the specified id from the Sites table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowSites(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "sites", id)
}

// rGetIdOrCreateSites returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Statuses table.
func (c *CrmDatabase) GetIdOrCreateStatuses(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "statuses", valRecord)
}

// delRowAuxilTable deletes the row with the specified id from the Statuses table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowStatuses(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "statuses", id)
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
