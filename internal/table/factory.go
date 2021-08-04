package table

import "log"

var tables = map[string]Table{
	"logstore_standard_log": &LogstoreRecord{},
}

// Table is the interface implemented by all supported tables.
type Table interface {
	parse(line []string)
}

// NewRecord provides an empty record of the appropriate type.
func NewRecord(tableName string) interface{} {
	record, exists := tables[tableName]
	if !exists {
		log.Fatalf("Unsupported table (%s)", tableName)
	}

	return record
}

// NewRecordParsed provides a record with data parsed into the appropriate field types.
func NewRecordParsed(tableName string, line []string) interface{} {
	record, exists := tables[tableName]
	if !exists {
		log.Fatalf("Unsupported table (%s)", tableName)
	}

	record.parse(line)

	return record
}
