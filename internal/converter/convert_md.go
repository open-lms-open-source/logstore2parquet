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
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func doConvertMD(schemaFile string, inputFile string, outputFile string) {
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

	md, err := readLines(schemaFile)
	if err != nil {
		log.Fatalf("Schema problem (%s)", err.Error())
	}

	pw, err := writer.NewCSVWriter(md, fw, 4)
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

		record := make([]*string, len(line))
		for j := 0; j < len(line); j++ {
			record[j] = &line[j]
		}

		if err = pw.WriteString(record); err != nil {
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
