package table

import "strconv"

// ContextRecord describes the fields in the Moodle™ context table.
type ContextRecord struct {
	ID           int64  `parquet:"name=id, type=INT64"`
	ContextLevel int64  `parquet:"name=contextlevel, type=INT64"`
	InstanceID   int64  `parquet:"name=instanceid, type=INT64"`
	Path         string `parquet:"name=path, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Depth        int64  `parquet:"name=depth, type=INT32"`
	Locked       int64  `parquet:"name=locked, type=INT32"`
}

func (r *ContextRecord) parse(line []string) {
	id, _ := strconv.ParseInt(line[0], 10, 64)
	contextlevel, _ := strconv.ParseInt(line[1], 10, 64)
	instanceid, _ := strconv.ParseInt(line[2], 10, 64)
	depth, _ := strconv.ParseInt(line[4], 10, 32)
	locked, _ := strconv.ParseInt(line[5], 10, 32)

	r.ID = id
	r.ContextLevel = contextlevel
	r.InstanceID = instanceid
	r.Depth = depth
	r.Locked = locked
}
