package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/clfs/simple/movegen"
)

var (
	fenFlag   = flag.String("fen", fen.Starting, "position")
	depthFlag = flag.Int("depth", 1, "search depth")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *depthFlag < 1 {
		log.Fatal("error: -depth must be at least 1")
	}

	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(w io.Writer) error {
	p, err := fen.Decode(*fenFlag)
	if err != nil {
		return fmt.Errorf("invalid FEN: %v", err)
	}

	var (
		sum  int
		rows []string
	)

	for m, n := range movegen.Divide(p, *depthFlag) {
		sum += n
		rows = append(rows, fmt.Sprintf("%s\t%d", pcn.Encode(m), n))
	}

	slices.Sort(rows)

	fmt.Fprintln(w, sum)
	for _, row := range rows {
		fmt.Fprintln(w, row)
	}

	return nil
}
