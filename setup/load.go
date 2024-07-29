package setup

import (
	"fmt"
	"github.com/rbo-17/95737-final-project/db"
	"github.com/rbo-17/95737-final-project/utils"
	"strconv"
)

// LoadStarterDataset loads starter data into dbIns
func LoadStarterDataset(dbIns db.Db, dataType utils.TestDataType) (map[int64]TestRecord, error) {

	// Get starter data
	dataSet, err := GetStarterDataSet(dataType)
	if err != nil {
		return nil, err
	}

	dataSetList := RecordsMapToList(dataSet)

	utils.Print(fmt.Sprintf("Loading starter dataset of %d items...", len(dataSet)))
	bCnt := 0
	for i := 0; i < len(dataSetList); i += 100 {
		batchList := dataSetList[i : i+100]
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

	utils.Print(fmt.Sprintf("Finished loading %d items (avg of %dB/item) and %d bytes of payload data", iCnt, bPerItem, bCnt))

	return dataSet, nil
}
