package setup

import (
	"fmt"
	"github.com/rbo-17/95737-final-project/db"
	"github.com/rbo-17/95737-final-project/utils"
	"strconv"
)

// LoadStarterDataset loads starter data into dbIns
func LoadStarterDataset(dbIns db.Db, dataType utils.TestDataType, opts utils.TestOpts) (map[int64]TestRecord, error) {

	// Get starter data
	dataSet, err := GetStarterDataSet(dataType, opts)
	if err != nil {
		return nil, err
	}

	dataSetList := RecordsMapToList(dataSet)

	utils.Print(fmt.Sprintf("Loading starter dataset of %d items into db...", len(dataSetList)))
	bCnt := 0
	for i := 0; i < len(dataSetList); i += 100 {

		upperLimit := i + 100
		if upperLimit > len(dataSetList) {
			upperLimit = i + len(dataSetList)%100
		}

		batchList := dataSetList[i:upperLimit]
		batch := make(map[string]*[]byte, len(batchList))

		for _, v := range batchList {
			key := dbIns.GetKey(strconv.Itoa(int(v.KeyId)))
			batch[key] = &v.Value
			bCnt += len(v.Value)
		}

		err = dbIns.PutMany(batch)
		if err != nil {
			return nil, err
		}
	}

	iCnt := len(dataSet)
	bPerItem := bCnt / iCnt

	utils.Print(fmt.Sprintf("Finished loading %d items (avg of %dB/item) and %d bytes of payload data into db", iCnt, bPerItem, bCnt))

	return dataSet, nil
}
