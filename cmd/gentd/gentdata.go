package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maaw77/crmsrvg/internal/models"
)

func genData(numEvrthng, numRow int, fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\n")

	dateСr := []pgtype.Date{
		{Time: time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC), Valid: true},
		{Time: time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	dateCh := []pgtype.Date{
		{Time: time.Time{}, Valid: true},
		{Time: time.Date(2025, 2, 11, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range numRow {
		gsmE := models.GsmTableEntry{
			DtReceiving:  dateСr[r.Intn(2)],
			DtCrch:       dateCh[r.Intn(2)],
			Site:         "Site_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			IncomeKg:     r.Float64() * 1000,
			Operator:     "Operator_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			Provider:     "Provider_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			Contractor:   "Contractor_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			LicensePlate: "LicensePlate_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			Status:       "Status_" + strconv.Itoa(r.Intn(numEvrthng)+1),
			GUID:         uuid.NewString(),
		}
		if gsmE.DtCrch != dateCh[0] {
			gsmE.BeenChanged = true
		}
		enc.Encode(gsmE)
		// log.Println(gsmE)

	}
	return nil
}

func readData(fileName string) (gsmEntries []models.GsmTableEntry, err error) {
	// fmt.Println("###############################")
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	for {
		gsmE := models.GsmTableEntry{}
		err = dec.Decode(&gsmE)

		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		// log.Println(gsmE)
		gsmEntries = append(gsmEntries, gsmE)
	}

	return gsmEntries, nil
}

func main() {
	// const (
	// 	numEvrthng = 3
	// 	numRow     = 4
	// 	fileName   = "gsm.data"
	// )
	var (
		numEvrthng, numRow int
		fileName           string
	)

	flag.IntVar(&numEvrthng, "nv", 2, "number of variants")
	flag.IntVar(&numRow, "nr", 1, "number of rows")
	flag.StringVar(&fileName, "fname", "gsm.data", "file name")
	// log.Println(os.Args)
	flag.Parse()

	log.Println(numEvrthng, numRow, fileName)

	if err := genData(numEvrthng, numRow, fileName); err != nil {
		log.Fatal(err)
	}

	gsmEntries, err := readData(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(gsmEntries))
	fmt.Println(gsmEntries)

}
