package database

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// GetIdOrCreateSites returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Sites table.
func (c *CrmDatabase) GetIdOrCreateSites(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "sites", valRecord)
}

// DelRowSites deletes the row with the specified id from the Sites table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowSites(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "sites", id)
}

// GetIdOrCreateStatuses returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Statuses table.
func (c *CrmDatabase) GetIdOrCreateStatuses(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "statuses", valRecord)
}

// DelRowStatuses deletes the row with the specified id from the Statuses table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowStatuses(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "statuses", id)
}

// GetIdOrCreatContractors returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Contractors table.
func (c *CrmDatabase) GetIdOrCreateContractors(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "contractors", valRecord)
}

// DelRowContractors deletes the row with the specified id from the Contractors table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowContractors(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "contractors", id)
}

// GetIdOrCreateLicensePlates returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the License_plates  table.
func (c *CrmDatabase) GetIdOrCreateLicensePlates(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "license_plates", valRecord)
}

// DelRowLicensePlatesdeletes the row with the specified id from the License_plates table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowLicensePlates(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "license_plates", id)
}

// GetIdOrCreateOperators returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Operators table.
func (c *CrmDatabase) GetIdOrCreateOperators(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "operators", valRecord)
}

// DelRowOperators deletes the row with the specified id from the Operators table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowOperators(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "operators", id)
}

// GetIdOrCreateProviders  returns the ID if the record exists, otherwise it creates the record and returns its ID.
// It's for the Providers table.
func (c *CrmDatabase) GetIdOrCreateProviders(ctx context.Context, valRecord string) (id models.IdEntry, err error) {
	return c.getIdOrCreateAuxilTable(ctx, "providers", valRecord)
}

// DelRowProviders deletes the row with the specified id from the Providers table and returns statusExe==true,
// otherwise statusExe==false.
func (c *CrmDatabase) DelRowProviders(ctx context.Context, id int) (statusExec bool, err error) {
	return c.delRowAuxilTable(ctx, "providers", id)
}

// getIdOrCreateAuxilTableTx  is similar to  getIdOrCreateAuxilTable, but is used pgxpool.Tx.
func (c *CrmDatabase) getIdOrCreateAuxilTableTx(ctx context.Context, txI pgx.Tx, nameTable, valRecord string) (id models.IdEntry, err error) {
	statmentGetId := fmt.Sprintf("SELECT id FROM %s WHERE name = $1;", nameTable)
	statmentCreate := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1) RETURNING id;", nameTable)

	tx, ok := txI.(*pgxpool.Tx)
	if !ok {
		return models.IdEntry{}, fmt.Errorf("%T != *pgxpool.Tx", tx)
	}

	err = tx.QueryRow(ctx, statmentGetId, valRecord).Scan(&(id.ID))

	if errors.Is(err, pgx.ErrNoRows) {
		err = tx.QueryRow(ctx, statmentCreate, valRecord).Scan(&(id.ID))
	}
	return
}

// InserGsmTable inserts a row into the Gsm table.
func (c *CrmDatabase) InserGsmTable(ctx context.Context, gsmEntry models.GsmTableEntry) (id models.IdEntry, err error) {
	tx, err := c.dbpool.Begin(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback(context.Background())

	statmentCreateGsmRow := `INSERT INTO gsm_table (id, 
                                            dt_receiving,
                                            dt_crch,
                                            income_kg,
                                            been_changed,
                                            db_data_creation,
                                            site_id,
                                            operator_id,
                                            provider_id,
                                            contractor_id,
                                            license_plate_id,
                                            status_id,
                                            guid) 
                                            VALUES (DEFAULT, $1, $2, $3, $4,  $5, $6, $7, $8,  $9, $10, $11, $12)
                                            RETURNING id;`

	idSite, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "sites", gsmEntry.Site)
	if err != nil {
		return
	}

	idOperator, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "operators", gsmEntry.Operator)
	if err != nil {
		return
	}

	idProvider, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "providers", gsmEntry.Provider)
	if err != nil {
		return
	}

	idContractor, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "ontractors", gsmEntry.Contractor)
	if err != nil {
		return
	}

	idLicensePlate, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "license_plates", gsmEntry.LicensePlate)
	if err != nil {
		return
	}

	idStatus, err := c.getIdOrCreateAuxilTableTx(ctx, tx, "statuses", gsmEntry.Status)
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx, statmentCreateGsmRow, gsmEntry.DtReceiving, gsmEntry.DtCrch, gsmEntry.IncomeKg, gsmEntry.BeenChanged, time.Now(),
		idSite.ID, idOperator.ID, idProvider.ID, idContractor.ID, idLicensePlate.ID, idStatus.ID, gsmEntry.GUID).Scan(&(id.ID))

	if err = tx.Commit(ctx); err != nil {
		return
	}

	return id, nil
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
