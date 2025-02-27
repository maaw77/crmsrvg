package database

import (
	"context"
	"log"
	"strconv"
	"testing"

	"github.com/maaw77/crmsrvg/internal/models"
)

var (
	// crmDB *CrmDatabase

	entriesB = map[string]models.IdEntry{}
)

func clearEnries() {
	for k, v := range entriesB {
		s, err := crmDB.DelRowSites(context.Background(), v.ID)

		switch {
		case err != nil:
			log.Printf("DelRowSites-> %s != nil. For %s", err, k)
		case s != true:
			log.Printf("DelRowSites-> %v != true. For %s", s, k)
		}

	}
}
func BenchmarkGetIdOrCreateSites(b *testing.B) {
	b.Cleanup(clearEnries)
	var err error
	for b.Loop() {
		for i := 0; i <= 10; i++ {
			entriesB["Site_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateSites(context.Background(), "Site_"+strconv.Itoa(i))
			if err != nil {
				b.Logf("GetIdOrCreateSites-> %s != nil", err)
			}
		}
	}

}
