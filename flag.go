package main

import "flag"

var (
	inputVar  string
	outputVar string
)

func init() {
	flag.StringVar(&inputVar, "i", "", "input file name")
	flag.StringVar(&outputVar, "o", "", "output file name")
}
