---
title: 'convert'
---
# Command: convert

The `convert` command transforms a supported input CSV file into a Parquet file.
Unsupported tables will end with an error without touching the output destination.

When specifying a table with the `--table` flag, don't include the `mdl_` prefix.

To use the `--schema` flag, a schema file must be provided that specifies the formats to use.

## Usage

```
l2p convert [flags] <inputfile> <outputfile>
```

## Output

This command writes a new parquet file.

## Examples

Convert a logstore_standard_log table (the default table type):
```
$ l2p convert input.csv output.parquet
```

Convert a different table:
```
$ l2p convert --table logstore_standard_log input.csv output.parquet
```

Convert a table based on an input schema:
```
$ cat input.schema
name=id, type=INT64
name=eventname, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=component, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=action, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=target, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=objecttable, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=objectid, type=INT64
name=crud, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=edulevel, type=INT32
name=contextid, type=INT64
name=contextlevel, type=INT64
name=contextinstanceid, type=INT64
name=userid, type=INT64
name=courseid, type=INT64
name=relateduserid, type=INT64
name=anonymous, type=INT32
name=other, type=BYTE_ARRAY, convertedtype=UTF8
name=timecreated, type=INT64
name=origin, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=ip, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY
name=realuserid, type=INT64

$ l2p convert --schema input.schema input.csv output.parquet
```