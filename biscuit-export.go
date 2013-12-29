package main

import (
	"encoding/csv"
	"fmt"
	goopt "github.com/droundy/goopt"
	"github.com/wilhelm-murdoch/biscuit"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	// Name defines the name of this utility
	Name = "biscuit-export"
	// Version defines this utility's current version
	Version = "1.0.0"
)

var (
	from      = goopt.String([]string{"-f", "--from"}, "", "path to text file you wish to export")
	to        = goopt.String([]string{"-t", "--to"}, "", "path to CSV you wish to export to")
	n         = goopt.Int([]string{"-n", "--length"}, 3, "length of the ngram sequence (default: 3)")
	version   = goopt.Flag([]string{"-v", "--version"}, []string{}, "current version of this utility", "")
	overwrite = goopt.Flag([]string{"-o", "--overwrite"}, []string{}, "overwrite existing ngram table export if it exists?", "")
)

func getNgramTableAsArray(p *biscuit.Profile) [][]string {
	var table [][]string

	for sequence, frequency := range p.Ngrams {
		table = append(table, []string{sequence, strconv.Itoa(frequency)})
	}

	return table
}

func init() {
	goopt.Description = func() string {
		return "A simple utility used to quickly generate precalculated ngram tables and export them as CSV files."
	}
	goopt.Version = Version
	goopt.Summary = "Ngram table to CSV export utility."
	goopt.Parse(nil)
}

func main() {
	if *version {
		fmt.Println(Name, Version)
		os.Exit(0)
	}

	if len(strings.TrimSpace(*from)) == 0 {
		fmt.Println("Please specify a valid file for -f or --from")
		os.Exit(1)
	}

	if len(strings.TrimSpace(*to)) == 0 {
		fmt.Println("Please specify a valid destination path for -t or --to")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(*from)
	if err != nil {
		fmt.Println("Invalid file specified:", err)
		os.Exit(1)
	}

	fileFlags := os.O_RDWR | os.O_CREATE
	if *overwrite {
		fileFlags |= os.O_TRUNC
	} else if _, err := os.Stat(*to); err == nil {
		fmt.Println("Destination file exists; skipping...")
		os.Exit(1)
	}

	profile := biscuit.NewProfileFromText("profile", string(bytes), *n)

	f, err := os.OpenFile(*to, fileFlags, 0666)
	if err != nil {
		fmt.Println("Invalid file specified:", err)
		os.Exit(1)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(getNgramTableAsArray(profile))
	if err != nil {
		fmt.Println("Error in exporting CSV:", err)
		os.Exit(1)
	}
	defer f.Close()

	exit := fmt.Sprintf("%d sequences written to %s ...", len(profile.Ngrams), *to)
	fmt.Println(exit)
	os.Exit(0)
}
