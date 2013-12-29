### biscuit-export
***
Command line utility for creating CSV representations of a Biscuit ngram table. This is to be used in conjunction with the [Biscuit](http://github.com/wilhelm-murdoch/biscuit) package and the importing of precalculated ngram tables.

### Installation
You'll have to compile this on your own, so make sure you have the Go compiler installed on your machine. This utility was written with version `1.1.2` and it might work with earlier versions, though I have not tested it yet.

1. Clone into your `$GOPATH/src` directory.
2. Fetch all external dependancies with `go get -v`
3. Navigate into the `$GOPATH/src/github.com/wilhelm-murdoch/biscuit-export` directory and run `go install`

If all went well, the executable should now reside within `$GOPATH/bin`. If you want it available throughout your system, just add `$GOPATH/bin` to your systems' `$PATH`.

### Usage

You can find usage documentation with the following command:

```
$: biscuit-export --help
Usage of biscuit-export:
  Ngram table to CSV export utility.
Options:
  -f    --from=      path to text file you wish to export
  -t    --to=        path to CSV you wish to export to
  -n 3  --length=3   length of the ngram sequence (default: 3)
  -v    --version    current version of this utility
  -o    --overwrite  overwrite existing ngram table export if it exists?
        --help       show usage message
```

The following will take the specified source text `en.txt`, compile it into a Biscuit ngram table and then export it as a CSV file named `en.csv`. If `en.csv` already exist, it will be truncated and overwritten.

```
$: biscuit-export -f en.txt -t en.csv -o
2573 sequences written to en.csv ...
```
or, for verbosity ...
```
$: biscuit-export --from en.txt --to en.csv --overwrite
2573 sequences written to en.csv ...
```

Omitting the `-o, --overwrite` flag will cause the utility to exit with the following message if the file already exists.

```
$: biscuit-export -f en.txt -t en.csv
Destination file exists; skipping...
exit status 1
```
