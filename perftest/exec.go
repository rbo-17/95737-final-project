package perftest

import (
	"encoding/csv"
	"errors"
	"fmt"
	dbi "github.com/rbo-17/95737-final-project/db"
	"github.com/rbo-17/95737-final-project/setup"
	"github.com/rbo-17/95737-final-project/utils"
	"os"
	"strconv"
	"sync"
	"time"
)

func RunTest(db dbi.Db, testType utils.TestType, dataType utils.TestDataType) {
	utils.UpdatePrefix(db.GetName(), testType, dataType)

	// Set up db connection and load test data
	err := db.Init()
	if err != nil {
		panic(err)
	}

	sds, err := setup.LoadStarterDataset(db, dataType)
	if err != nil {
		panic(err)
	}

	// Get new (unloaded) records to perform test with
	writeFactor, err := utils.TestTypeToWriteFactor(testType)
	if err != nil {
		panic(err)
	}

	// Add extra to write factor to account for random variation
	nds, err := setup.GetTestDataSet(writeFactor+0.05, dataType)
	if err != nil {
		panic(err)
	}

	utils.Print("Loading new dataset for testing...")

	// Prepare & run test
	ops, err := Prepare(sds, nds, testType, dataType)
	if err != nil {
		panic(err)
	}

	utils.Print("Loading complete. Starting test now.")
	start := time.Now()
	err = Run(db, testType, dataType, ops)
	if err != nil {
		panic(err)
	}

	dur := int(time.Since(start).Seconds())
	utils.Print(fmt.Sprintf("Testing completed in %d seconds.", dur))

	// Clean up db
	err = db.DeleteAll()
	if err != nil {
		panic(err)
	}

	// Close connections
	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func Run(dbIns dbi.Db, testType utils.TestType, dataType utils.TestDataType, ops []TestOp) error {

	inCh := make(chan TestOp, len(ops))
	outCh := make(chan TestOpResult, len(ops))
	reqWg := new(sync.WaitGroup)
	responses := make([]TestOpResult, 0)

	utils.Print(fmt.Sprintf("Spinning up %d workers & loading into channel", utils.WorkerCount))

	// Create workers
	for i := 0; i < utils.WorkerCount; i++ {
		reqWg.Add(1)
		go PerformOpWorker(dbIns, inCh, outCh, reqWg)
	}

	// Start timer
	start := time.Now()

	// Load data into channel
	for _, op := range ops {
		inCh <- op
	}

	utils.Print(fmt.Sprintf("All data loaded into channels, proceeding to wait..."))

	close(inCh)
	reqWg.Wait()

	// Stop timer
	dur := time.Since(start)

	close(outCh)

	// Read results
	for res := range outCh {
		responses = append(responses, res)
	}

	utils.Print("All data has been received, writing to file")

	// Write results to file
	err := WriteResultsToFile(dbIns.GetName(), dur, testType, dataType, &responses)
	if err != nil {
		return err
	}

	return nil
}

func PerformOpWorker(db dbi.Db, inCh chan TestOp, outCh chan TestOpResult, wg *sync.WaitGroup) {

	defer wg.Done()

	for op := range inCh {
		res := PerformOp(db, op)
		outCh <- res
	}
}

func PerformOp(db dbi.Db, op TestOp) TestOpResult {

	start := time.Now()

	if op.OpType == utils.OpTypeGet {
		key := db.GetKey(strconv.Itoa(int(op.KeyId)))
		res, err := db.Get(key)
		if err != nil {
			return TestOpResult{
				Time:    start,
				OpType:  op.OpType,
				KeyId:   op.KeyId,
				Latency: 0,
				Ok:      false,
				Err:     err,
			}
		}

		// Validate returned bytes matches expected value
		if len(*res) != op.ValueSize {
			errMsg := fmt.Sprintf("bytes returned (%d) do not match expected count (%d)", len(*res), op.ValueSize)
			return TestOpResult{
				Time:    start,
				OpType:  op.OpType,
				KeyId:   op.KeyId,
				Latency: 0,
				Ok:      false,
				Err:     errors.New(errMsg),
			}
		}

	} else if op.OpType == utils.OpTypePut {
		key := db.GetKey(strconv.Itoa(int(op.KeyId)))
		err := db.Put(key, op.Value)
		if err != nil {
			return TestOpResult{
				Time:    start,
				OpType:  op.OpType,
				KeyId:   op.KeyId,
				Latency: 0,
				Ok:      false,
				Err:     err,
			}
		}

	} else {
		return TestOpResult{
			Time:    start,
			OpType:  op.OpType,
			KeyId:   0,
			Latency: 0,
			Ok:      false,
			Err:     errors.New("invalid OpType detected"),
		}
	}

	dur := time.Since(start)

	return TestOpResult{
		Time:    start,
		OpType:  op.OpType,
		KeyId:   op.KeyId,
		Latency: dur.Microseconds(),
		Ok:      true,
		Err:     nil,
	}
}

func WriteResultsToFile(dbName string, totalTimeTaken time.Duration, testType utils.TestType, dataType utils.TestDataType, results *[]TestOpResult) error {

	// Convert results to list of lists of strings
	output := make([][]string, len(*results))
	for i, res := range *results {

		errMsg := "nil"
		if res.Err != nil {
			errMsg = res.Err.Error()
		}

		output[i] = []string{
			res.Time.Format("2006-01-02T15:04:05.999999-07:00"),
			string(res.OpType),
			strconv.Itoa(int(res.KeyId)),
			strconv.Itoa(int(res.Latency)),
			strconv.FormatBool(res.Ok),
			errMsg,
		}
	}

	// Create results dir if it doesn't exist
	resultsDir := "results/raw"
	err := os.MkdirAll(resultsDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create new file
	ts := time.Now().Unix()
	fname := fmt.Sprintf("%s/%s-%s-%s-%d-%d.csv", resultsDir, dbName, testType, dataType, totalTimeTaken.Milliseconds(), ts)
	csvFile, err := os.Create(fname)
	if err != nil {
		return err
	}

	utils.Print(fmt.Sprintf("Writing results to file: %s", fname))

	// Write the CSV data
	writer := csv.NewWriter(csvFile)

	headers := []string{"Time", "Op", "Key id", "Latency", "Ok", "Error"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, row := range output {
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	// Cleanup
	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}

	err = csvFile.Close()
	if err != nil {
		return err
	}

	return nil
}
