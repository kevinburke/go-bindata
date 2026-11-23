## v4.0.3 (2026-11-22)

- Document the credits and licenses for the benchmark photos and their
  embedded copies. Thanks to Helmut Grohne for reporting the missing credits.
- Test with Go 1.26 and 1.27 in CI. Validate generated fixtures on both
  versions and compare them byte for byte on Go 1.26.
- Modernize the Go build constraints, package documentation, and prefix
  handling.

## v4

This is meant to satisfy go module versioning behavior; see
https://stackoverflow.com/a/52058348/329700 for more.

## 3.24.0

Remove uses of io/ioutil; you must use Go 1.18 or higher with this version of
go-bindata and its generated asset files.

Update generated doc comments for compatibility with Go's updated doc comment
guidelines.

## 3.21.0

Replace "Debug" with "AssetDebug" to reduce the likelihood of conflicts.

## 3.20.0

Add the "Debug" constant if assets have been generated using the `--debug` flag
at the command line.
