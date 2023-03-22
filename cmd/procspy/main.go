package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/paulstuart/procspy"
)

var (
	lookup bool
	listen bool
	format bool
)

func init() {
	flag.BoolVar(&lookup, "lookup", false, "look up process name")
	flag.BoolVar(&listen, "listen", false, "show listening ports")
	flag.BoolVar(&format, "json", false, "json format")
}

func main() {
	flag.Parse()

	cs, err := procspy.Connections(lookup, listen)
	if err != nil {
		panic(err)
	}
	var enc *json.Encoder
	if format {
		enc = json.NewEncoder(os.Stdout)
	}

	fmt.Fprintln(os.Stderr, "TCP Connections:")
	for c := cs.Next(); c != nil; c = cs.Next() {
		if format {
			_ = enc.Encode(c)
		} else {
			fmt.Printf(" - %+v\n", c)
		}
	}

}
