---
title: 'version'
---
# Command: version

The `version` command displays build information about the installed binary,
including the release version and the date and time it was built.

## Usage

```
l2p version
```

## Output

This command prints the version number and the exact date and time it was
built.  The version may include additional information if the build was not
performed from a tagged commit.  This may include the number of additional
commits since the version tag, the commit hash, and/or the string `dirty`
indicating that local, uncommitted change were detected at build time.

## Examples

```
$ l2p version
v0.2-dirty, built 2021-08-04T11:48:52+0930
```
