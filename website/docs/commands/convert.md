---
title: 'convert'
---
# Command: convert

The `convert` command transforms a supported input CSV file into a Parquet file.
Unsupported tables will end with an error without touching the output destination.

When specifying a table with the `--table` flag, don't include the `mdl_` prefix.

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