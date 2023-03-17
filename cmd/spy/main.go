package main

import (
	"flag"
	"fmt"

	"github.com/paulstuart/procspy"
)

var (
	lookup bool
	listen bool
)

func init() {
	flag.BoolVar(&lookup, "lookup", false, "look up process name")
	flag.BoolVar(&listen, "listen", false, "show listening ports")
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		fmt.Println("arg:", arg)
	}
	cs, err := procspy.Connections(lookup, listen)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TCP Connections:\n")
	for c := cs.Next(); c != nil; c = cs.Next() {
		fmt.Printf(" - %+v\n", c)
	}

}
