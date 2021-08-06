package table

import "strconv"

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

func (t *LogstoreRecord) parse(line []string) {
	id, _ := strconv.ParseInt(line[0], 10, 64)
	objectID, _ := strconv.ParseInt(line[6], 10, 64)
	eduLevel, _ := strconv.ParseInt(line[8], 10, 32)
	contextID, _ := strconv.ParseInt(line[9], 10, 64)
	contextLevel, _ := strconv.ParseInt(line[10], 10, 64)
	contextInstanceID, _ := strconv.ParseInt(line[11], 10, 64)
	userID, _ := strconv.ParseInt(line[12], 10, 64)
	courseID, _ := strconv.ParseInt(line[13], 10, 64)
	relatedUserID, _ := strconv.ParseInt(line[14], 10, 64)
	anonymous, _ := strconv.ParseInt(line[15], 10, 32)
	timeCreated, _ := strconv.ParseInt(line[17], 10, 64)
	realUserID, _ := strconv.ParseInt(line[20], 10, 64)

	t.ID = id
	t.Eventname = line[1]
	t.Component = line[2]
	t.Action = line[3]
	t.Target = line[4]
	t.ObjectTable = line[5]
	t.ObjectID = objectID
	t.CRUD = line[7]
	t.EduLevel = eduLevel
	t.ContextID = contextID
	t.ContextLevel = contextLevel
	t.ContextInstanceID = contextInstanceID
	t.UserID = userID
	t.CourseID = courseID
	t.RelatedUserID = relatedUserID
	t.Anonymous = anonymous
	t.Other = line[16]
	t.TimeCreated = timeCreated
	t.Origin = line[18]
	t.IP = line[19]
	t.RealUserID = realUserID
}
