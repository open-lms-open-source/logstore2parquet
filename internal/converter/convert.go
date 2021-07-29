package converter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

// LogstoreRecord describes the fields in the Moodle™ logstore_standard_log table.
type LogstoreRecord struct {
	ID                int64  `parquet:"name=id, type=INT64"`
	Eventname         string `parquet:"name=eventname, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Component         string `parquet:"name=component, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Action            string `parquet:"name=action, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Target            string `parquet:"name=target, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ObjectTable       string `parquet:"name=objecttable, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	ObjectID          int64  `parquet:"name=objectid, type=INT64"`
	CRUD              string `parquet:"name=crud, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	EduLevel          int64  `parquet:"name=edulevel, type=INT32"`
	ContextID         int64  `parquet:"name=contextid, type=INT64"`
	ContextLevel      int64  `parquet:"name=contextlevel, type=INT64"`
	ContextInstanceID int64  `parquet:"name=contextinstanceid, type=INT64"`
	UserID            int64  `parquet:"name=userid, type=INT64"`
	CourseID          int64  `parquet:"name=courseid, type=INT64"`
	RelatedUserID     int64  `parquet:"name=relateduserid, type=INT64"`
	Anonymous         int64  `parquet:"name=anonymous, type=INT32"`
	Other             string `parquet:"name=other, type=BYTE_ARRAY, convertedtype=UTF8"`
	TimeCreated       int64  `parquet:"name=timecreated, type=INT64"`
	Origin            string `parquet:"name=origin, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	IP                string `parquet:"name=ip, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	RealUserID        int64  `parquet:"name=realuserid, type=INT64"`
}

func doConvert(inputFile string, outputFile string) {
	var err error

	csvFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Cannot find input file: %s", inputFile)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	fw, err := local.NewLocalFileWriter(outputFile)
	if err != nil {
		log.Fatalf("Can't create local file (%s)", err.Error())
	}

	pw, err := writer.NewParquetWriter(fw, new(LogstoreRecord), 2)
	if err != nil {
		log.Fatalf("Can't create parquet writer (%s)", err.Error())
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		id, _ := strconv.ParseInt(line[0], 10, 64)
		objectID, _ := strconv.ParseInt(line[6], 10, 64)
		eduLevel, _ := strconv.ParseInt(line[8], 10, 64)
		contextID, _ := strconv.ParseInt(line[9], 10, 64)
		contextLevel, _ := strconv.ParseInt(line[10], 10, 64)
		contextInstanceID, _ := strconv.ParseInt(line[11], 10, 64)
		userID, _ := strconv.ParseInt(line[12], 10, 64)
		courseID, _ := strconv.ParseInt(line[13], 10, 64)
		relatedUserID, _ := strconv.ParseInt(line[14], 10, 64)
		anonymous, _ := strconv.ParseInt(line[15], 10, 64)
		timeCreated, _ := strconv.ParseInt(line[17], 10, 64)
		realUserID, _ := strconv.ParseInt(line[20], 10, 64)

		logRecord := LogstoreRecord{
			ID:                id,
			Eventname:         line[1],
			Component:         line[2],
			Action:            line[3],
			Target:            line[4],
			ObjectTable:       line[5],
			ObjectID:          objectID,
			CRUD:              line[7],
			EduLevel:          eduLevel,
			ContextID:         contextID,
			ContextLevel:      contextLevel,
			ContextInstanceID: contextInstanceID,
			UserID:            userID,
			CourseID:          courseID,
			RelatedUserID:     relatedUserID,
			Anonymous:         anonymous,
			Other:             line[16],
			TimeCreated:       timeCreated,
			Origin:            line[18],
			IP:                line[19],
			RealUserID:        realUserID,
		}
		if err = pw.Write(logRecord); err != nil {
			log.Println("Write error", err)
		}
	}

	if err = pw.WriteStop(); err != nil {
		log.Fatalf("WriteStop error (%s)", err.Error())
		return
	}

	log.Println("Write Finished")
	fw.Close()
}
