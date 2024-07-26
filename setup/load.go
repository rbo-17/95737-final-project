package setup

import (
	"fmt"
	"github.com/rbo-17/95737-final-project/db"
	"github.com/rbo-17/95737-final-project/utils"
	"strconv"
)

// TODO: Convert "puts" to "putMany"
// LoadStarterDataset loads starter data into dbIns
func LoadStarterDataset(dbIns db.Db, dataType utils.TestDataType) (map[int64]TestRecord, error) {

	// Get starter data
	dataSet, err := GetStarterDataSet(dataType)
	if err != nil {
		return nil, err
	}

	// Load data into database
	//prefix := fmt.Sprintf("[%s][%s]", dbIns.GetName(), string(dataType))
	//fmt.Println(fmt.Sprintf("%s Loading %d items...", prefix, len(dataSet)))

	dataSetList := RecordsMapToList(dataSet)

	utils.Print(fmt.Sprintf("Loading starter dataset of %d items...", len(dataSet)))
	bCnt := 0
	//fmt.Println("len(dataSetList)", len(dataSetList))
	for i := 0; i < len(dataSetList); i += 100 {
		batchList := dataSetList[i : i+100]
		batch := make(map[string][]byte, len(batchList))

		for _, v := range batchList {
			key := dbIns.GetKey(strconv.Itoa(int(v.KeyId)))
			batch[key] = v.Value
			bCnt += len(v.Value)
		}

		err = dbIns.PutMany(batch)
		if err != nil {
			return nil, err
		}
	}

	//utils.Print(fmt.Sprintf("Loading starter dataset of %d items...", len(dataSet)))
	//bCnt := 0
	//for k, v := range dataSet {
	//	key := dbIns.GetKey(strconv.Itoa(int(k)))
	//	value := v.Value
	//
	//	err = dbIns.Put(key, value)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	bCnt += len(value)
	//}

	iCnt := len(dataSet)
	bPerItem := bCnt / iCnt

	//fmt.Println(fmt.Sprintf("%s Finished loading %d bytes of payload data", prefix, bCount))
	utils.Print(fmt.Sprintf("Finished loading %d items (avg of %dB/item) and %d bytes of payload data", iCnt, bPerItem, bCnt))

	return dataSet, nil
}

//	fmt.Println(fmt.Sprintf("%s Loading %d small texts...", prefix, len(dataSet.SmallTexts)))
//	for k, v := range dataSet.SmallTexts {
//		key := fmt.Sprintf("smtxt:%d", k)
//		value := v.Payload
//
//		err := dbIns.Put(key, value)
//		if err != nil {
//			return err
//		}
//	}
//	fmt.Println(fmt.Sprintf("%s Done loading small texts", prefix))
//
//	fmt.Println(fmt.Sprintf("%s Loading %d large texts...", prefix, len(dataSet.LargeTexts)))
//	for k, v := range dataSet.LargeTexts {
//		key := fmt.Sprintf("lgtxt:%d", k)
//		value := v.Payload
//
//		err := dbIns.Put(key, value)
//		if err != nil {
//			return err
//		}
//	}
//	fmt.Println(fmt.Sprintf("%s Done loading large texts", prefix))
//
//	fmt.Println(fmt.Sprintf("%s Loading %d images...", prefix, len(dataSet.Images)))
//	for k, v := range dataSet.Images {
//		key := fmt.Sprintf("img:%d", k)
//		value := v.Payload
//
//		err := dbIns.Put(key, value)
//		if err != nil {
//			return err
//		}
//	}
//	fmt.Println(fmt.Sprintf("%s Done loading images", prefix))
//
//	return nil
//}
