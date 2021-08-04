package utils

import (
	"axobase/app/datamodels"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ParseAntibodiesCSV(filename string) ([]datamodels.AntibodyRecord, error) {
	var err error
	var records = make([]datamodels.AntibodyRecord, 0)

	dataPath, err := ReadDataPath()
	if err != nil {
		return records, err
	}

	filepath := fmt.Sprintf("%s/%s", dataPath, filename)
	fileHandle, err := os.Open(filepath)
	if err != nil {
		return records, err
	}

	// Get all line items for the statement.
	defer fileHandle.Close()
	reader := csv.NewReader(fileHandle)
	lines, err := reader.ReadAll()
	if err != nil {
		return records, err
	}

	for _, elem := range lines[1:] {
		fmt.Println(elem)
		r := datamodels.AntibodyRecord{
			AntigenTarget:       elem[0],
			CellTypes:           elem[1],
			ExampleImages:       elem[2],
			Vendor:              elem[3],
			SecondaryValidation: elem[4],
			Dilution:            elem[5],
			Fixation:            elem[6],
			Application:         elem[7],
			Isotype:             elem[8],
			CloneName:           elem[9],
			SpeciesReactivity:   elem[10],
			Description:         elem[11],
			CatalogNum:          elem[12],
			VendorPageLink:      elem[13],
			PublicationLinks:    elem[14],
			Notes:               elem[15],
			SubmittedBy:         elem[16],
		}
		rating, err := strconv.ParseFloat(elem[17], 64)
		if err != nil {
			// TODO: Log the error
			r.Rating = -9999
		} else {
			r.Rating = int64(rating)
		}

		records = append(records, r)
	}
	return records, nil
}
