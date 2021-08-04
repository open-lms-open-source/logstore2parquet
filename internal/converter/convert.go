package converter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"

	"logstore2parquet/internal/table"
)

func doConvert(tableName string, inputFile string, outputFile string) {
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

	pw, err := writer.NewParquetWriter(fw, table.NewRecord(tableName), 2)
	if err != nil {
		log.Fatalf("Can't create parquet writer (%s)", err.Error())
	}

	log.Printf("Parsing records for table: %s", tableName)

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		record := table.NewRecordParsed(tableName, line)

		if err = pw.Write(record); err != nil {
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
